package iotutil

import (
	"container/list"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Job struct {
	F         func(...interface{}) interface{}
	Args      []interface{}
	Result    interface{}
	Err       error
	added     chan bool
	Worker_id uint
	Job_id    uint64
}

type stats struct {
	Submitted int
	Running   int
	Completed int
}

type Pool struct {
	workers_started      bool
	supervisor_started   bool
	num_workers          int
	job_wanted_pipe      chan chan *Job
	done_pipe            chan *Job
	add_pipe             chan *Job
	result_wanted_pipe   chan chan *Job
	jobs_ready_to_run    *list.List
	num_jobs_submitted   int
	num_jobs_running     int
	num_jobs_completed   int
	jobs_completed       *list.List
	interval             time.Duration
	working_wanted_pipe  chan chan bool
	stats_wanted_pipe    chan chan stats
	worker_kill_pipe     chan bool
	supervisor_kill_pipe chan bool
	worker_wg            sync.WaitGroup
	supervisor_wg        sync.WaitGroup
	next_job_id          uint64
}

func (pool *Pool) subworker(job *Job) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic while running job:", err)
			job.Result = nil
			job.Err = fmt.Errorf(err.(string))
		}
	}()
	job.Result = job.F(job.Args...)
}

func (pool *Pool) worker(worker_id uint) {
	job_pipe := make(chan *Job)
WORKER_LOOP:
	for {
		pool.job_wanted_pipe <- job_pipe
		job := <-job_pipe
		if job == nil {
			time.Sleep(pool.interval * time.Millisecond)
		} else {
			job.Worker_id = worker_id
			pool.subworker(job)
			pool.done_pipe <- job
		}
		select {
		case <-pool.worker_kill_pipe:
			break WORKER_LOOP
		default:
		}
	}
	pool.worker_wg.Done()
}

func NewPool(workers int) (pool *Pool) {
	pool = new(Pool)
	pool.num_workers = workers
	pool.job_wanted_pipe = make(chan chan *Job)
	pool.done_pipe = make(chan *Job)
	pool.add_pipe = make(chan *Job)
	pool.result_wanted_pipe = make(chan chan *Job)
	pool.jobs_ready_to_run = list.New()
	pool.jobs_completed = list.New()
	pool.working_wanted_pipe = make(chan chan bool)
	pool.stats_wanted_pipe = make(chan chan stats)
	pool.worker_kill_pipe = make(chan bool)
	pool.supervisor_kill_pipe = make(chan bool)
	pool.interval = 1000
	pool.next_job_id = 0
	pool.startSupervisor()
	return
}

func (pool *Pool) supervisor() {
SUPERVISOR_LOOP:
	for {
		select {
		case job := <-pool.add_pipe:
			pool.jobs_ready_to_run.PushBack(job)
			pool.num_jobs_submitted++
			job.added <- true
		case job_pipe := <-pool.job_wanted_pipe:
			element := pool.jobs_ready_to_run.Front()
			var job *Job = nil
			if element != nil {
				job = element.Value.(*Job)
				pool.num_jobs_running++
				pool.jobs_ready_to_run.Remove(element)
			}
			job_pipe <- job
		case job := <-pool.done_pipe:
			pool.num_jobs_running--
			pool.jobs_completed.PushBack(job)
			pool.num_jobs_completed++
		case result_pipe := <-pool.result_wanted_pipe:
			close_pipe := false
			job := (*Job)(nil)
			element := pool.jobs_completed.Front()
			if element != nil {
				job = element.Value.(*Job)
				pool.jobs_completed.Remove(element)
			} else {
				if pool.num_jobs_running == 0 && pool.num_jobs_completed == pool.num_jobs_submitted {
					close_pipe = true
				}
			}
			if close_pipe {
				close(result_pipe)
			} else {
				result_pipe <- job
			}
		case working_pipe := <-pool.working_wanted_pipe:
			working := true
			if pool.jobs_ready_to_run.Len() == 0 && pool.num_jobs_running == 0 {
				working = false
			}
			working_pipe <- working
		case stats_pipe := <-pool.stats_wanted_pipe:
			pool_stats := stats{pool.num_jobs_submitted, pool.num_jobs_running, pool.num_jobs_completed}
			stats_pipe <- pool_stats
		case <-pool.supervisor_kill_pipe:
			break SUPERVISOR_LOOP
		}
	}
	pool.supervisor_wg.Done()
}

func (pool *Pool) Run() {
	if pool.workers_started {
		panic("trying to start a pool that's already running")
	}
	for i := uint(0); i < uint(pool.num_workers); i++ {
		pool.worker_wg.Add(1)
		go pool.worker(i)
	}
	pool.workers_started = true
	if !pool.supervisor_started {
		pool.startSupervisor()
	}
}

func (pool *Pool) Stop() {
	if !pool.workers_started {
		panic("trying to stop a pool that's already stopped")
	}
	for i := 0; i < pool.num_workers; i++ {
		pool.worker_kill_pipe <- true
	}
	pool.worker_wg.Wait()
	pool.workers_started = false
	if pool.supervisor_started {
		pool.stopSupervisor()
	}
}

func (pool *Pool) startSupervisor() {
	pool.supervisor_wg.Add(1)
	go pool.supervisor()
	pool.supervisor_started = true
}

func (pool *Pool) stopSupervisor() {
	pool.supervisor_kill_pipe <- true
	pool.supervisor_wg.Wait()
	pool.supervisor_started = false
}

func (pool *Pool) Add(f func(...interface{}) interface{}, args ...interface{}) {
	job := &Job{f, args, nil, nil, make(chan bool), 0, pool.getNextJobId()}
	pool.add_pipe <- job
	<-job.added
}

func (pool *Pool) getNextJobId() uint64 {
	return atomic.AddUint64(&pool.next_job_id, 1)
}

func (pool *Pool) Wait() {
	working_pipe := make(chan bool)
	for {
		pool.working_wanted_pipe <- working_pipe
		if !<-working_pipe {
			break
		}
		time.Sleep(pool.interval * time.Millisecond)
	}
}

func (pool *Pool) Results() (res []*Job) {
	res = make([]*Job, pool.jobs_completed.Len())
	i := 0
	for e := pool.jobs_completed.Front(); e != nil; e = e.Next() {
		res[i] = e.Value.(*Job)
		i++
	}
	pool.jobs_completed = list.New()
	return
}

func (pool *Pool) WaitForJob() *Job {
	result_pipe := make(chan *Job)
	var job *Job
	var ok bool
	for {
		pool.result_wanted_pipe <- result_pipe
		job, ok = <-result_pipe
		if !ok {
			return nil
		}
		if job == (*Job)(nil) {
			time.Sleep(pool.interval * time.Millisecond)
		} else {
			break
		}
	}
	return job
}

func (pool *Pool) Status() stats {
	stats_pipe := make(chan stats)
	if pool.supervisor_started {
		pool.stats_wanted_pipe <- stats_pipe
		return <-stats_pipe
	}
	return stats{}
}
