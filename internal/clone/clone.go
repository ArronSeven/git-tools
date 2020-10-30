package clone

import (
	"fmt"
	"git-tools/internal"
	"git-tools/internal/branch"
	"os/exec"
	"strings"
	"sync"
)

// clone handler
type Handler struct {
	// 组的树结构
	// the trees struct of group
	groupTree []*Tree
	// 组的列表结构
	// the list struct fo group
	groups []*Tree
	client Client
	// 标识是否克隆过项目
	// if clone success is true
	success bool
}

// A
//  B
//   C
//  D
//   E
// groupName = C 只克隆C组下面的,并创建A/B/C文件夹目录
// groupName = A 克隆A组下面的
// clone
func (e *Handler) clone() {
	if len(e.groups) < 1 {
		return
	}
	var wg sync.WaitGroup

	for _, group := range e.groups {
		// 排除不需要克隆的组
		// exclude group name
		if !group.contain(e.client.config.CloneConfig.GroupName) {
			continue
		}
		projects, err := e.client.GetProjectsByGroupId(group.ID)
		if err != nil {
			fmt.Printf("Get projects by group id:(%s) occurs err:%v\n", group.Name, err)
			continue
		}
		go e.cloneProject(projects, group, &wg)
	}
	wg.Wait()
	if !e.success {
		fmt.Println("No project was cloned this time, please try check [group name] or whether existing project in the group")
	}

}

// 克隆项目
// execute clone projects
func (e *Handler) cloneProject(projects []*Project, group *Tree, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	// 只有存在项目才创建目录
	if len(projects) > 0 {
		err := group.makeDir(e.client.config.CloneConfig.Output)
		if err != nil {
			fmt.Printf("Make directory occurs err:%v, path:%s\n", err, group.dirPath)
			return
		}
	}

	for _, pro := range projects {
		// 执行git命令
		cmd := exec.Command("git", "clone", pro.CloneAddr)
		cmd.Dir = group.absolutePath
		err := internal.Exec(cmd)
		if err != nil {
			fmt.Printf("execute [git clone] command occure err:%v,clone address:%s\n", err, pro.CloneAddr)
			fmt.Println()
			return
		}
		e.success = true
	}
	// checkout branch
	if strings.TrimSpace(internal.Cfg.BranchConfig.BranchName) != "" {
		internal.Cfg.BranchConfig.DirPath = group.absolutePath
		branchHandler := branch.Handler{Config: internal.Cfg.BranchConfig}
		branchHandler.SwitchBranch()
	}
}
