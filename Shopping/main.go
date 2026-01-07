package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

type Product struct {
    name string
	price int
} 

var Menu = map[int]Product{
	1:{
		"薯片",
		5,
	},
	2:{
		"可乐",
		3,
	},
	3:{
		"果冻",
		1,
	},
	4:{
		"奶龙手办",
		9999,
	},
}

var Customer = make(chan int,5)
var Fst bool = false
var Sec bool = false
var Wait sync.WaitGroup

func go1(){
	for{
		Fst = true
		if(len(Customer)==0){
			break
		}		   
		id := <-Customer
		fmt.Println("1号台接收到顾客 商品id : ",id," 商品名称 :",Menu[id].name,"商品价格 : ",Menu[id].price,"$")
		time.Sleep(time.Duration(rand.Intn(4)*1000))
		fmt.Println("1号台结算结束")
        Fst = false
	}
	Wait.Done()
}

func go2(){
	for{
		if(!Fst){
			continue
		}		   
		Sec = true
		if(len(Customer)==0){
			break
		}		   
		id := <-Customer
		fmt.Println("2号台接收到顾客 商品id : ",id," 商品名称 :",Menu[id].name,"商品价格 : ",Menu[id].price,"$")
		time.Sleep(time.Duration(rand.Intn(4)*1000))
		fmt.Println("2号台结算结束")
        Sec = false
	}
	Wait.Done()
}

func go3(){
	for{
		if(!Sec){
			continue
		}
		if(len(Customer)==0){
			break
		}		   
		id := <-Customer
		fmt.Println("3号台接收到顾客 商品id : ",id," 商品名称 :",Menu[id].name,"商品价格 : ",Menu[id].price,"$")
		time.Sleep(time.Duration(rand.Intn(4)*1000))
		fmt.Println("3号台结算结束")
	}
	Wait.Done()
}

func main(){
	rand.Seed(time.Now().UnixNano())
    Wait.Add(3)	
	for i:=0;i<5;i++{
		Customer <- rand.Intn(4) + 1
	}
	go go1()
	go go2()
	go go3()
    Wait.Wait()
}