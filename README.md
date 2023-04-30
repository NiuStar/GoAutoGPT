# GoAutoGPT

[English Version README](./README_EN.md)

[中文版 README](./README.md)

## 使用方法

采用ChatGPT实现人工智能编程，输入一个名字，自动完成架构设计、功能分解、代码编程

使用方法：

1、`git clone git@github.com:NiuStar/GoAutoGPT.git`

2、`cd GoAutoGPT`

3、`go build`

4、修改config目录下的config文件中的UserId和src，userId是你申请的用户唯一标识，src是代码生成目录

目前config中默认配置了个userId为1的测试账户，如果需要另外账户，请邮件与我沟通，我的邮箱：yjkj2@qq.com



5、`./GoCodeGPT create -p 洗衣机预约管理系统 -d 这是一个测试的系统简介描述`

一杯咖啡的时间，如果create顺利就会在$src下看到你要生成的代码了

如果出错了，请使用

`./GoCodeGPT generate -p '你所要生成的projectId'`

![Image description](https://github.com/NiuStar/GoAutoGPT/raw/main/example/generate.png)

projectId查看方式为：

`./GoCodeGPT list`

![Image description](https://github.com/NiuStar/GoAutoGPT/raw/main/example/list.jpg)

找到你要生成的project的uuid

其中的nameEn就是$src下代码目录名称



腾讯云低价主机网速太慢，请见谅。

## 未来计划：

1）更友好的交互方式

2）更多新的特性

3）支持web、Android、iOS的人工智能编程计划
