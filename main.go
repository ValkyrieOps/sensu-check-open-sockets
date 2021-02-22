package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"io/ioutil"
	"strings"
	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Warn int
	Crit int
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-check-open-sockets",
			Short:    "The Sensu Go Open Sockets plugin",
			Keyspace: "sensu.io/plugins/sensu-check-open-sockets/config",
		},
	}

	options = []*sensu.PluginConfigOption{
                &sensu.PluginConfigOption{
                        Path:      "warn",
                        Env:       "CHECK_WARN",
                        Argument:  "warn",
                        Shorthand: "w",
                        Default:   10,
                        Usage:     "Warning threshold - count of open sockets required for warning state",
                        Value:     &plugin.Warn,
                },
                &sensu.PluginConfigOption{
                        Path:      "crit",
                        Env:       "CHECK_CRITICAL",
                        Argument:  "crit",
                        Shorthand: "c",
                        Default:   20,
                        Usage:     "Critical threshold - count of open sockets required for critical state",
                        Value:     &plugin.Crit,
                },
        }
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
        if plugin.Warn == 0 {
                return sensu.CheckStateWarning, fmt.Errorf("--warn or CHECK_WARN must be greater than zero")
        }
        return sensu.CheckStateOK, nil

        if plugin.Warn > plugin.Crit {
                return sensu.CheckStateWarning, fmt.Errorf("--crit or CHECK_CRITICAL must be greater than warning")
        }
        return sensu.CheckStateOK, nil
}

func handleError(err error) {
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}

func executeCheck(event *types.Event) (int, error) {
	get_sockets := exec.Command("ss", "-s")
	get_head := exec.Command("head", "-1")
	awk_sockets := exec.Command("awk", "{print $2}")


	get_head.Stdin, _ = get_sockets.StdoutPipe()
	awk_sockets.Stdin, _ = get_head.StdoutPipe()
	stdout, err := awk_sockets.StdoutPipe()

	awk_sockets.Start()
	get_head.Start()
	err = get_sockets.Start()
	handleError(err)

	defer awk_sockets.Wait()
	defer get_head.Wait()

	socket_count, _ := ioutil.ReadAll(stdout)
	sockets, _ := strconv.Atoi(strings.TrimSuffix(string(socket_count), "\n"))


	if sockets >= plugin.Warn && sockets < plugin.Crit {
		fmt.Println("WARNING\nOpen Sockets:", sockets)
		return sensu.CheckStateWarning, nil
	} else if sockets >= plugin.Crit {
		fmt.Println("CRITICAL\nOpen Sockets:", sockets)
		return sensu.CheckStateCritical, nil
	} else {
		fmt.Println("OK\nOpen Sockets:", sockets)
		return sensu.CheckStateOK, nil
	}


}
