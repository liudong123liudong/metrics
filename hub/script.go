package hub

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

var (
	Timeout    = 3 * time.Second
	ErrTimeout = errors.New("command timed out")
)

type Script struct {
	Mutex sync.Mutex
	Cache map[string]string
	Ch    chan string
}

func NewScript() *Script {
	return &Script{
		Cache: make(map[string]string),
		Ch:    make(chan string, 256),
	}
}

func (s *Script) Exec(checkNum, name string) {
	filePos := fmt.Sprintf("upload/%s", name)
	fileAbs, _ := filepath.Abs(filePos)

	cmd := exec.Command("/bin/bash", fileAbs)
	ret, err := cmd.Output()

	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if err != nil {
		s.Cache[checkNum] = err.Error()
		ret, _ := json.Marshal(map[string]string{checkNum: err.Error()})
		s.Ch <- string(ret)
	} else {
		s.Cache[checkNum] = string(ret)
		ret, _ := json.Marshal(map[string]string{checkNum: string(ret)})
		s.Ch <- string(ret)
	}

	// 删除脚本文件
	_ = os.Remove(fileAbs)
}
