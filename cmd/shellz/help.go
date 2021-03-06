package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/evilsocket/shellz/core"
	"github.com/evilsocket/shellz/log"
)

var (
	examples = []struct {
		cmd  string
		help string
	}{
		{"shellz -list", "list all available identities and shells"},
		{"shellz -list -on web", "list all available identities and shells of the group web"},

		{"shellz -enable \"machineA, machineB\"", "enable the shells named machineA and machineB"},
		{"shellz -disable machineA", "disable the shell named machineA (commands won't be executed on it)"},

		{"shellz -test", "test all shells and disable the not responding ones"},
		{"shellz -test -on \"machineA, machineB\" -connection-timeout 1s", "test two shells and disable them if they don't respond within 1 second"},

		{"shellz -run id", "run the command 'id' on each shell"},
		{"shellz -run id -stats", "run the command 'id' on each shell and print some statistics once finished"},
		{"shellz -run id -on machineA", "run the command 'id' on a single shell named 'machineA'"},
		{"shellz -run id -on 'machineA, machineB'", "run the command 'id' on machineA and machineB"},
		{"shellz -run uptime -to all.txt", "run the command 'uptime' on every shell and append all outputs to the 'all.txt' file"},
		{"shellz -run uptime -to \"{{.Identity.Username}}_{{.Name}}.txt\"", "run the command 'uptime' on every shell and save each outputs to a different file using per-shell data."},
	}
)

func init() {
	flag.BoolVar(&doList, "list", doList, "List available shells and exit.")

	flag.StringVar(&doEnable, "enable", "", "Enable the specified shells.")
	flag.StringVar(&doDisable, "disable", "", "Disable the specified shells.")

	flag.BoolVar(&doTest, "test", doTest, "Attempt to run a test command on the selected shells and disable the ones who failed.")
	flag.BoolVar(&doForce, "force", doForce, "Include disabled shells in the selection.")
	flag.BoolVar(&doStats, "stats", doStats, "Print some statistics after the -run and -test commands.")

	flag.StringVar(&command, "run", command, "Command to run on the selected shells.")
	flag.StringVar(&onFilter, "on", onFilter, "Comma separated list of shell names to select or * for all.")
	flag.StringVar(&toOutput, "to", toOutput, "If filled, commands output will be saved to this file instead of being printed on the standard output.")

	flag.DurationVar(&timeouts.Connect, "connection-timeout", timeouts.Connect, "Connection timeout.")
	flag.DurationVar(&timeouts.Read, "read-timeout", timeouts.Read, "Read timeout.")
	flag.DurationVar(&timeouts.Write, "write-timeout", timeouts.Write, "Write timeout.")

	flag.BoolVar(&log.DebugMessages, "debug", log.DebugMessages, "Enable debug messages.")
	flag.StringVar(&log.File, "log-file", log.File, "Log messages on this file instead of the standard output.")
	flag.BoolVar(&log.NoColors, "no-colors", log.NoColors, "Disable colors for log messages.")
	flag.BoolVar(&noBanner, "no-banner", noBanner, "Don't print the initial banner.")
	flag.Parse()
}

func showHelp() {
	log.Info("none of the -run, -test or -list parameters have been specified")

	fmt.Println()
	fmt.Printf("Usage:\n\n")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Printf("Examples:\n\n")

	for _, e := range examples {
		fmt.Printf("  %s\n", core.Dim("# "+e.help))
		fmt.Printf("  %s\n", core.Bold(e.cmd))
		fmt.Println()
	}

	os.Exit(1)
}
