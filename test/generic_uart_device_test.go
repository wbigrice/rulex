package test

import (
	"time"

	"github.com/i4de/rulex/core"
	"github.com/i4de/rulex/engine"
	"github.com/i4de/rulex/glogger"
	httpserver "github.com/i4de/rulex/plugin/http_server"

	"testing"

	"github.com/i4de/rulex/typex"
)

/*
*
* Test_UART_Device
*
 */
func Test_UART_Device(t *testing.T) {
	mainConfig := core.InitGlobalConfig("conf/rulex.ini")
	core.GlobalConfig.AppDebugMode = true
	glogger.StartGLogger(true, core.GlobalConfig.LogPath)
	glogger.StartLuaLogger(core.GlobalConfig.LuaLogPath)
	core.StartStore(core.GlobalConfig.MaxQueueSize)
	core.SetLogLevel()
	core.SetPerformance()
	engine := engine.NewRuleEngine(mainConfig)
	engine.Start()

	// HttpApiServer loaded default
	if err := engine.LoadPlugin("plugin.http_server", httpserver.NewHttpApiServer()); err != nil {
		t.Fatal("HttpServer load failed:", err)
	}

	// YK8 Inend
	GUART := typex.NewDevice(typex.GENERIC_UART,
		"UART", "UART", "UART", map[string]interface{}{
			"baudRate":  9600,
			"dataBits":  8,
			"frequency": 5,
			"parity":    "N",
			"stopBits":  1,
			"tag":       "tag1",
			"timeout":   5,
			"uart":      "COM2",
		})
	GUART.UUID = "GUART1"
	if err := engine.LoadDevice(GUART); err != nil {
		t.Fatal("GUART load failed:", err)
	}

	rule := typex.NewRule(engine,
		"uuid",
		"test",
		"test",
		[]string{},
		[]string{GUART.UUID},
		`function Success() print("[LUA Success Callback]=> OK") end`,
		`
		Actions = {
		function(data)
			print(data)
			return true, data
		end
	}
`,
		`function Failed(error) print("[LUA Failed Callback]", error) end`)
	if err := engine.LoadRule(rule); err != nil {
		t.Fatal(err)
	}
	time.Sleep(20 * time.Second)
	engine.Stop()
}
