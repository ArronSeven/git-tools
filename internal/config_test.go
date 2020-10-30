package internal

import (
	"flag"
	"os"
	"strings"
	"testing"
)

func TestCloneParamParse(t *testing.T) {
	var args = `git_tools -t cg -addr 127.0.0.1:8080 -token RmMtvCTdmnsmZZcjV9Sb -group IOT2 -out D:\tmp -branch dev-1.0.0`
	os.Args = strings.Split(args, " ")
	flag.Parse()
	if Cfg.Type != "cg" {
		t.Errorf("the '-t' param parse fail, Expect:'cg', Actual:'%s'\n", Cfg.Type)
	}

	if Cfg.CloneConfig.Addr != "127.0.0.1:8080" {
		t.Errorf("the '-addr' param parse fail, Expect:'127.0.0.1:8080', Actual:'%s'\n", Cfg.CloneConfig.Addr)
	}
	if Cfg.CloneConfig.AccessToken != "RmMtvCTdmnsmZZcjV9Sb" {
		t.Errorf("the '-token' param parse fail, Expect:'RmMtvCTdmnsmZZcjV9Sb', Actual:'%s'\n", Cfg.CloneConfig.AccessToken)
	}
	if Cfg.CloneConfig.GroupName != "IOT2" {
		t.Errorf("the '-group' param parse fail, Expect:'IOT2', Actual:'%s'\n", Cfg.CloneConfig.GroupName)
	}
	if Cfg.CloneConfig.Output != "D:\\tmp" {
		t.Errorf("the '-out' param parse fail, Expect:'D:\\tmp', Actual:'%s'\n", Cfg.CloneConfig.Output)
	}
	if Cfg.BranchConfig.BranchName != "dev-1.0.0" {
		t.Errorf("the '-branch' param parse fail, Expect:'dev-1.0.0', Actual:'%s'\n", Cfg.BranchConfig.BranchName)
	}
}

func TestBranchParamParse(t *testing.T) {
	var args = `git_tools -t b -dir D:\tmp -branch dev1.0.0'`
	os.Args = strings.Split(args, " ")
	flag.Parse()
	if Cfg.Type != "b" {
		t.Errorf("the '-t' param parse fail, Expect:'b', Actual:'%s'\n", Cfg.Type)
	}

	if Cfg.BranchConfig.BranchName != "dev-1.0.0" {
		t.Errorf("the '-branch' param parse fail, Expect:'dev-1.0.0', Actual:'%s'\n", Cfg.BranchConfig.BranchName)
	}
	if Cfg.BranchConfig.DirPath != "D:\\tmp" {
		t.Errorf("the '-dir' param parse fail, Expect:'D:\\tmp', Actual:'%s'\n", Cfg.BranchConfig.DirPath)
	}
}
