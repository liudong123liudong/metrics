package hub

import (
	"time"
)

type Hub struct {
	ScriptTask *Script
}

var MyHub = &Hub{
	ScriptTask: NewScript(),
}

func (h *Hub) Run() {
	go sendHostInfo()

	go sendHostMetrics()

	go findScriptTask()

	for {
		select {
		case ret := <-h.ScriptTask.Ch:
			sendScriptResult(ret)
		default:
			time.Sleep(5*time.Second)
		}
	}
}


