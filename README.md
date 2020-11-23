# Catilina

> Quo usque tandem abutere, Catilina, patientia nostra?

## 简介

此项目用于演示如何使用 Go 来搭建一个简单的 Web 项目。

- 自上而下

    依赖顺序：`cmd -> app -> controller -> service -> model`, 其中前面的包可以引用后面包的方法，而后面则不能引用前面的包（防止循环调用）