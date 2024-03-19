package service

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func StartAuto() bool {
	h := md5.New()
	execName := os.Args[0]
	dir := filepath.Dir(execName)
	dirAbs, err := filepath.Abs(dir)
	if err == nil && dirAbs != dir {
		execName = filepath.Join(dirAbs, execName)
	}
	h.Write([]byte(execName))
	filename := hex.EncodeToString(h.Sum(nil))
	p, e := Home()
	if e != nil {
		p = "./"
	}
	fp := filepath.Join(p, "."+filepath.Base(execName)+"_"+filename+".sock")
	return Start(fp)
}

func Start(socketFile string) bool {
	c, _ := net.Dial("unix", socketFile)
	if c != nil {
		return true
	}
	go func() {
		os.Remove(socketFile)
		l, err := net.Listen("unix", socketFile)
		if err != nil {
			panic(err)
		}
		defer l.Close()
		for {
			conn, err := l.Accept()
			if err != nil {
			} else {
				conn.Close()
			}
		}
	}()
	return false
}

func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}
	if "windows" == runtime.GOOS {
		return homeWindows()
	}
	return homeUnix()
}

func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}
	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
