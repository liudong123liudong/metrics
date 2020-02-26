package agent_id

import (
	"fmt"
	"futong_server_agent_go/utils"
	"futong_server_agent_go/utils/global"
	"futong_server_agent_go/utils/logger"
	"io/ioutil"
	"os"
)

var (
	AgentId string
	agentIdFilePath = global.Global.Path + "/agent.uuid"
)

func generateAgentId() {
	_, err := os.Stat(agentIdFilePath)
	if err == nil {
		return
	}

	f, err := os.Create(agentIdFilePath)
	defer f.Close()
	if err != nil {
		logger.Sugar.Error("create agent.uuid err: ", err)
	}

	_, err = f.WriteString(utils.GenerateUUID())
	if err != nil {
		logger.Sugar.Error("write to agent.uuid err: ", err)
	}
}

func init() {
	fmt.Println("init agent.uuid......")
	generateAgentId()

	content, err :=  ioutil.ReadFile(agentIdFilePath)
	if err != nil || content == nil {
		logger.Sugar.Error("open and read agent.uuid err: ", err)
	}

	AgentId = string(content)
}

//
//func GetAgentId() string {
//	generateAgentId()
//
//	content, err :=  ioutil.ReadFile(agentIdFilePath)
//	if err != nil || content == nil {
//		logger.Sugar.Error("open and read agent.uuid err: ", err)
//	}
//
//	return string(content)
//}