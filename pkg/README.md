### `/pkg`

外部应用程序可以使用的库代码（如，`/pkg/mypubliclib`）。其他项目将会导入这些库来保证项目可以正常运行，所以在将代码放在这里前，一定要三思而行。请注意，`internal`目录是一个更好的选择来确保项目私有代码不会被其他人导入，因为这是Go强制执行的。使用`/pkg`目录来明确表示代码可以被其他人安全的导入仍然是一个好方式。Travis Jeffery撰写的关于 [I’ll take pkg over internal](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) 文章很好地概述了`pkg`和`inernal`目录以及何时使用它们。

当您的根目录包含大量非Go组件和目录时，这也是一种将Go代码分组到一个位置的方法，从而使运行各种Go工具更加容易（在如下的文章中都有提到：2018年GopherCon [Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)，[Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) ，Golab 2018 [Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)）。

点击查看`/pkg`就能看到那些使用这个布局模式的流行Go代码仓库。这是一种常见的布局模式，但未被普遍接受，并且Go社区中的某些人不推荐这样做。

如果项目确实很小并且嵌套的层次并不会带来多少价值（除非你就是想用它），那么就不要使用它。请仔细思考这种情况，当项目变得很大，并且根目录中包含的内容相当繁杂（尤其是有很多非Go的组件）。
