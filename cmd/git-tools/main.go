package main

import (
	"flag"
	"fmt"
	"git-tools/internal"
	"git-tools/internal/branch"
	"git-tools/internal/clone"
	"os"
)

var usage = `Usage:
	%s -t <the type that 'cg' or 'b'(branch)> -addr <ip:port> -group <group name> -token <gitlab token> -out [the output directory] -dir [the branch name]
	%s -t <the type that 'cg' or 'b'(branch)> -branch <the branch name> -dir <Need to switch the directory path of the branch>`

func main() {
	flag.Usage = func() {
		fmt.Printf(usage, os.Args[0], os.Args[0])
		fmt.Println()
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}
	switch internal.Cfg.Type {
	case "cg":
		// clone
		err := internal.Cfg.CloneConfig.CheckParam()
		if err != nil {
			fmt.Println(err)
			flag.Usage()
			os.Exit(1)
		}
		client := clone.InitClient()
		err = clone.Exec(client)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("clone finished")
		}
	case "b":
		err := internal.Cfg.BranchConfig.CheckParam()
		if err != nil {
			fmt.Println(err)
			flag.Usage()
			os.Exit(1)
		}
		branchHandler := branch.Handler{Config: internal.Cfg.BranchConfig}
		branchHandler.SwitchBranch()
	default:
		flag.Usage()

	}
}
