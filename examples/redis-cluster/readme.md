# Git 恢复误删的远程分支

### 在当前git目录下，输入以下命令查找删除分支的commitId

```bash
git reflog --date=iso
```

reflog是reference log的意思，也就是引用log，记录HEAD在各个分支上的移动轨迹。选项 --date=iso，表示以标准时间格式展示。这里为什么不用git log？因为git log是用来记录当前分支的commit log，分支都删除了，找不到commit log了。  
找到目标分支最后一次的commitid，

### 切出本地分支

```bash
git checkout -b recovery_branch_name commitid
```

### 推到远程即可

```bash
git push  origin recovery_branch_name 
```