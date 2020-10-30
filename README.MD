# Git-tools

This is Git tools that helping us develop more conveniently.

Translations:

* [中文文档](README_zh.md)

## Overview

- Clone project by group name
- Switch branch of a directory

## Getting Started

1.  Generate `SSH KEY` and put it into the `SSH KEYS` on the Gitlab.
    ```cmd
    git config --global user.name "yourname"
    git config --global user.email "youremail"
    ssh-keygen -t rsa -C “youremail”
    ```
2. In Gitlab `Setting->Access Tokens` and generate an only read token.
3. Usage:
    
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
#### [Download](https://github.com/ArronSeven/git-tools/releases)