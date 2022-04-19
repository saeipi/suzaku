### `/build`

打包和持续集成。

将云（AMI），容器（Docker），操作系统（deb，rpm，pkg）软件包配置和脚本放在`/build/package`目录中。

将CI（travis、circle、drone）配置文件和就脚本放在`build/ci`目录中。请注意，有一些CI工具（如，travis CI）对于配置文件的位置有严格的要求。尝试将配置文件放在`/build/ci`目录，然后链接到CI工具想要的位置。
