package conf

import (
	"fmt"
	"futong_server_agent_go/utils/global"
	"github.com/Unknwon/goconfig"
	"log"
)

var Config *goconfig.ConfigFile

func init() {
	fmt.Println("init conf.ini......")
	var err error
	Config, err = goconfig.LoadConfigFile(global.Global.Path + "/agent-config.ini")
	if err != nil {
		log.Fatal("init conf.ini error:", err)
	}
}
