package main

// 一种常用的GO设计模式，
// 用于确保HandlerBasedOnMap肯定实现了这个接口
var _ Handler = &HandlerBasedOnMap{}
