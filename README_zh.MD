# Git-tools
Git小工具

## 概要
- 克隆一个组的项目
- 切换目录下的所有项目分支

## 快速开始
    
 1. 首先生成`SSH KEY`在`gitlab`上配置好.
    ```cmd
     git config --global user.name "yourname"
     git config --global user.email "youremail"
     ssh-keygen -t rsa -C “youremail”
     ```
 2. 在Gitlab的setting->Access Tokens, 生成一个只读的token,我目前是使用Gitlab-v4的rest api.

 3. [下载使用](https://github.com/ArronSeven/git-tools/releases)
 
    clone projects
    ```cmd
    git_tools -t cg -addr <ip:port> -token RmMtvCTdmnsmZZcjV9Sb -group IOT2 -out D:\tmp
    ```
    clone projects and switch branch
    ```cmd
    git_tools -t cg -addr <ip:port> -token RmMtvCTdmnsmZZcjV9Sb -group IOT2 -out [D:\tmp] -branch [dev-1.0.0] 
    ```
    switch branch
    ```cmd
    git_tools -t b -dir D:\tmp -branch dev-1.0.0 
    ```