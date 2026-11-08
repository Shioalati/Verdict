# Verdict

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![License](https://img.shields.io/github/license/Shioalati/Verdict)

[English](README.md) | **简体中文**

「轻量、易用」的哈希 CLI

一个用 Golang 编写的命令行哈希工具，专注于简洁的交互与开箱即用的体验。

## 特性

- 简洁的交互体验

## 安装
```
go install github.com/Shioalati/Verdict/cmd/vec@latest
```
或从 Releases 页面下载预编译二进制文件。

## 基础用法

```
vec ./example.txt
```
校验指定文件的哈希值。

```
vec ./some-dir
```
校验目录：递归遍历目录下的所有文件（跳过符号链接），按相对路径排序后聚合，输出单个哈希值。

```
vec -l ./some-dir
```
以明细模式输出目录中每个文件的哈希与相对路径，每行格式为 `<hash>  <相对路径>`，不输出聚合值。

## 配置脚本

```
vec init
```
将在运行目录下创建 `Verdict.json`。
