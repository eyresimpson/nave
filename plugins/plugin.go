package plugins

import (
	"fmt"
	shared "nave/plugins/db/shard"
	"nave/tools/log"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func main() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugin/db/dbo"),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Err("ERROR", err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Err("Error", err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(shared.Greeter)
	fmt.Println(greeter.Greet())
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// pluginMap is the map of plugins we can dispense.
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
