# Git-tools
> Git小工具,目前提供克隆一个组的项目到本地,切换目录下所有项目的分支.

## 使用教程
 1. 首先生成SSH KEY在gitlab上配置好,可以使用git协议克隆下来.
 2. 在Gitlab的setting->Access Tokens, 生成一个只读的token,我目前是使用Gitlab-v4的rest api.
 3. 克隆项目并切换分支: git_tools -t cg -addr <ip:port> -token RmMtvCTdmnsmZZcjV9Sb -group IOT2 -out [D:\tmp] -branch [dev-1.0.0]
 4. 只克隆项目: git_tools -t cg -addr <ip:port> -token RmMtvCTdmnsmZZcjV9Sb -group IOT2 -out [D:\tmp]
 5. 只切换分支: git_tools -t b -dir D:\tmp -branch dev-1.0.0
 