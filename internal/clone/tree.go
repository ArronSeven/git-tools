package clone

import (
	"os"
	"strings"
)

type (
	// Gitlab中的组, 只留下了id
	//represents gitlab group object, only keep the group ID
	Group struct {
		ID       int        `json:"id"`
		Projects []*Project `json:"projects"`
	}
	// 每个组中的项目,只用到了克隆地址
	// Represents a Project in the group, only keep the address
	Project struct {
		CloneAddr string `json:"ssh_url_to_repo"`
	}
	// 用于保存组与组之间的层次关系,创建目录结构层次
	// The tree use project with group
	Tree struct {
		ID   int    `json:"id"`
		PId  int    `json:"parent_id"`
		Name string `json:"name"`
		Path string `json:"path"`
		// 创建目录的路径
		// the trees path
		dirPath      string
		absolutePath string
		Nodes        []*Tree
	}
)

func (t *Tree) contain(name string) bool {
	return strings.Contains(t.dirPath, name)
}

// make directory
func (t *Tree) makeDir(path string) error {
	// 创建目录
	path = strings.ReplaceAll(path, "\\", "/")
	if !strings.HasSuffix(path, "/") {
		t.absolutePath = path + "/" + t.dirPath
	}
	if _, err := os.Stat(t.absolutePath); os.IsNotExist(err) {
		err := os.MkdirAll(t.absolutePath, 0665)
		if err != nil {
			return err
		}
	}
	return nil
}

func Exec(client *Client) error {
	groups, err := client.GetGroups()
	if err != nil {
		return err
	}
	trees := ToTrees(groups)
	handle := &Handler{
		groupTree: trees,
		groups:    groups,
		client:    *client,
	}
	handle.clone()
	return nil
}

// 把数组结构转换成树结构
// list group to tree group
func ToTrees(data []*Tree) []*Tree {
	var root []*Tree
	var node []*Tree
	//  root node
	for _, m := range data {
		if m.PId == 0 {

			m.dirPath = m.Name
			root = append(root, m)
		} else {
			node = append(node, m)
		}
	}
	for _, r := range root {
		buildTree(node, r)
	}
	return root
}

// 构建组的树结构
// build tree
func buildTree(node []*Tree, r *Tree) {
	for _, n := range node {
		if r.ID == n.PId {
			n.dirPath = r.dirPath + "/" + n.Name
			r.Nodes = append(r.Nodes, n)
			buildTree(node, n)
		}
	}
}
