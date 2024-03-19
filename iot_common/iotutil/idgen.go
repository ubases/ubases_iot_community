package iotutil

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var IDGen IDWorker

func GetNextSeq() uint64 {
	for {
		if u64, err := IDGen.NextID(); err == nil {
			return uint64(u64)
		}
		time.Sleep(50 * time.Nanosecond)
	}
}

func GetNextSeqInt64() int64 {
	return int64(GetNextSeq())
}

func GetNextSeqInt32() int32 {
	return int32(GetNextSeq())
}

func GetNextSeqUint64() uint64 {
	return uint64(GetNextSeq())
}

func GetNextSeqStr() string {
	for {
		if u64, err := IDGen.NextID(); err == nil {
			return strconv.FormatUint(uint64(u64), 10)
		}
		time.Sleep(50 * time.Nanosecond)
	}
}

type IDWorker struct {
	startTime             int64
	workerIDBits          uint
	datacenterIdBits      uint
	maxWorkerId           int64
	maxDatacenterId       int64
	sequenceBits          uint
	workerIdLeftShift     uint
	datacenterIdLeftShift uint
	timestampLeftShift    uint
	sequenceMask          int64
	workerId              int64
	datacenterId          int64
	sequence              int64
	lastTimestamp         int64
	signMask              int64
	idLock                *sync.Mutex
}

func (iw *IDWorker) InitIDWorker(workerId, datacenterId int64) error {
	var baseValue int64 = -1
	iw.startTime = 1463834116272
	iw.workerIDBits = 5
	iw.datacenterIdBits = 5
	iw.maxWorkerId = baseValue ^ (baseValue << iw.workerIDBits)
	iw.maxDatacenterId = baseValue ^ (baseValue << iw.datacenterIdBits)
	iw.sequenceBits = 12
	iw.workerIdLeftShift = iw.sequenceBits
	iw.datacenterIdLeftShift = iw.workerIDBits + iw.workerIdLeftShift
	iw.timestampLeftShift = iw.datacenterIdBits + iw.datacenterIdLeftShift
	iw.sequenceMask = baseValue ^ (baseValue << iw.sequenceBits)
	iw.sequence = 0
	iw.lastTimestamp = -1
	iw.signMask = ^baseValue + 1
	iw.idLock = &sync.Mutex{}
	if iw.workerId < 0 || iw.workerId > iw.maxWorkerId {
		return fmt.Errorf("workerId[%v] is less than 0 or greater than maxWorkerId[%v].", workerId, datacenterId)
	}
	if iw.datacenterId < 0 || iw.datacenterId > iw.maxDatacenterId {
		return fmt.Errorf("datacenterId[%d] is less than 0 or greater than maxDatacenterId[%d].", workerId, datacenterId)
	}
	iw.workerId = workerId
	iw.datacenterId = datacenterId
	return nil
}

func (iw *IDWorker) NextID() (int64, error) {
	iw.idLock.Lock()
	defer iw.idLock.Unlock()
	timestamp := time.Now().UnixNano()
	if timestamp < iw.lastTimestamp {
		return -1, fmt.Errorf("Clock moved backwards.  Refusing to generate id for %d milliseconds", iw.lastTimestamp-timestamp)
	}

	if timestamp == iw.lastTimestamp {
		iw.sequence = (iw.sequence + 1) & iw.sequenceMask
		if iw.sequence == 0 {
			timestamp = iw.tilNextMillis()
			iw.sequence = 0
		}
	} else {
		iw.sequence = 0
	}

	iw.lastTimestamp = timestamp

	id := ((timestamp - iw.startTime) << iw.timestampLeftShift) |
		(iw.datacenterId << iw.datacenterIdLeftShift) |
		(iw.workerId << iw.workerIdLeftShift) |
		iw.sequence

	if id < 0 {
		id = -id
	}

	return id, nil
}

func (iw *IDWorker) tilNextMillis() int64 {
	timestamp := time.Now().UnixNano()
	if timestamp <= iw.lastTimestamp {
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return timestamp
}

func init() {
	//todo 多数据中心注意将datacenterId改为可配置
	IDGen.InitIDWorker(1000, 2)
}
