package plugins

import (
	shared "nave/plugins/noah/shared"
	"nave/tools/log"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// Load 加载插件
func Load(mod string) {
	// 日志工具
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// RPC通信主机
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("plugins/noah.nmod"),
		Logger:          logger,
	})

	defer client.Kill()

	// 通过RPC连接到插件
	rpcClient, err := client.Client()
	if err != nil {
		log.Err("ERROR", err)
	}

	// 请求插件
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Err("Error Load Plugin", err)
	}

	// 具体功能实现
	greeter := raw.(shared.Greeter)
	log.Success("Plug-in Noah install Success" + greeter.Greet())
}

// 用于握手的口令，必须对应才能链接成功
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "noahjones",
}

// pluginMap 是我们可以分发的插件地图。
var pluginMap = map[string]plugin.Plugin{
	"greeter": &shared.GreeterPlugin{},
}

// Exec 执行操作
func Exec(path string, params []string) {
	// ----------------------------------------------
	// TODO : 此处临时测试使用，尽快改成真正的动态调用方法
	// ----------------------------------------------
	switch path {
	case "Print":
		println(params[0])
		break
	case "Info":
		log.Info(params[0])
	case "Warn":
		log.Warn(params[0])
	case "Err":
		log.Err(params[0], nil)
	case "SQL":
		log.Info("MySQL")
	default:
		log.Warn("Cannot find plugins " + path)
	}
}
