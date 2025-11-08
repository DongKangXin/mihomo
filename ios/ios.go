package mihomo

import (
	C "github.com/metacubex/mihomo/constant"
	"github.com/metacubex/mihomo/hub"
	"github.com/metacubex/mihomo/hub/executor"
	"github.com/metacubex/mihomo/log"
)

func GetConfig() string {
	return C.Path.Config()
}

func StartClash(homeDir string, fileName string) {
	C.SetHomeDir(homeDir)
	C.SetConfig(fileName)
	log.SetLogFile(C.Path.Resolve("logs/clash.log"))
	cfg, err := executor.Parse()
	if err != nil {
		log.Errorln("clash 内核启动失败 + error: %s", err.Error())
	}
	hub.ApplyConfig(cfg)
}

func StopClash() {
	executor.Shutdown()
}

func main() {
	//StartClash("/Users/dongkangxin/Documents/clash", "config.yaml")
}
