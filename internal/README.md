### `/internal`

私有的应用程序代码库。这些是不希望被其他人导入的代码。请注意：这种模式是Go编译器强制执行的。详细内容情况Go 1.4的[release notes](https://golang.org/doc/go1.4#internalpackages)。再次注意，在项目的目录树中的任意位置都可以有`internal`目录，而不仅仅是在顶级目录中。

可以在内部代码包中添加一些额外的结构，来分隔共享和非共享的内部代码。这不是必选项（尤其是在小项目中），但是有一个直观的包用途是很棒的。应用程序实际的代码可以放在`/internal/app`目录（如，`internal/app/myapp`），而应用程序的共享代码放在`/internal/pkg`目录（如，`internal/pkg/myprivlib`）中。
