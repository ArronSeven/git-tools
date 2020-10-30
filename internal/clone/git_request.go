package clone

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"git-tools/internal"
	"io/ioutil"
	"net/http"
)

// 获取所有组
// Get all groups from Gitlab api
// help/api/api_resources.md
const groupUriFormat = "https://%s/api/v4/groups?private_token=%s&per_page=999&simple=true"

// 组下的项目
// Get projects by the group Api
// help/api/api_resources.md
const groupProjectFormat = "https://%s/api/v4/groups/%d?private_token=%s&per_page=999&simple=true"

type Client struct {
	client *http.Client
	config internal.Config
}

//获取Gitlab中所有的组
// Get all groups
func (c *Client) GetGroups() ([]*Tree, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf(groupUriFormat, c.config.CloneConfig.Addr, c.config.CloneConfig.AccessToken), nil)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		var data []*Tree
		err = json.Unmarshal(all, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, errors.New(string(all))
}

// 获取组下的项目
// Get projects by the group id.
func (c *Client) GetProjectsByGroupId(id int) ([]*Project, error) {
	url := fmt.Sprintf(groupProjectFormat, c.config.CloneConfig.Addr, id, c.config.CloneConfig.AccessToken)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(response.Body)
	if response.StatusCode == http.StatusOK {
		var group = &Group{}
		err = json.Unmarshal(all, group)
		if err != nil {
			return nil, err
		}
		return group.Projects, nil
	}
	return nil, errors.New(string(all))
}

// 初始化https请求client
// Initialize https client
func InitClient() *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Client{client: &http.Client{Transport: tr}, config: internal.Cfg}
}
