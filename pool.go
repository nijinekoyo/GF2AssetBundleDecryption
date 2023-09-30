/*
 * @Author: nijineko
 * @Date: 2023-09-30 23:07:50
 * @LastEditTime: 2023-09-30 23:07:53
 * @LastEditors: nijineko
 * @Description: 协程池实现
 * @FilePath: \GF2AssetBundleDecryption\pool.go
 */
package main

import "sync"

type Pool struct {
	workChan chan int
	wg       sync.WaitGroup
}

/**
 * @description: 生成一个协程池
 * @param {int} CoreNum 最大并发
 * @return {*Pool} 协程池
 */
func NewPool(CoreNum int) *Pool {
	ch := make(chan int, CoreNum)
	return &Pool{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}

/**
 * @description: 添加协程
 * @param {int} Num 协程数量
 * @return {*}
 */
func (P *Pool) Add(Num int) {
	for i := 0; i < Num; i++ {
		P.workChan <- i
		P.wg.Add(1)
	}
}

/**
 * @description: 结束一个协程
 * @param {*}
 * @return {*}
 */
func (P *Pool) Done() {
LOOP:
	for {
		<-P.workChan
		break LOOP
	}
	P.wg.Done()
}

/**
 * @description: 等待所有协程完成
 * @param {*}
 * @return {*}
 */
func (P *Pool) Wait() {
	P.wg.Wait()
}
