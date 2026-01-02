package main

import (
	"fmt"
)
type Book struct{
	name string
	years string
	coder string
}

type Library struct{
	num int
    book map[int]Book
}

func (u *Library)AddBook(name,years,coder string){//在末尾添加书籍
    u.num++
	u.book[u.num - 1]=Book{name,years,coder}
}

func (u *Library)RemoveBook(index int){//删除书籍 序号后的书籍前移补齐
	u.book[index]=Book{}
	for i:=index; i < u.num; i++{
		u.book[i] = u.book[i + 1]
	}
	u.num--
}

func (u Library)ShowAllBook(){//展示所有书
	fmt.Println("共有",u.num,"本书")
    for i := 0; i < u.num; i++{
		fmt.Printf("序号 : %d 名称 : %s 出版日期 : %s 作者 : %s\n",i,u.book[i].name,u.book[i].years,u.book[i].coder)
	}
}

func Solve() {
	
	a := make([]int,0) 
	for i := 0; i < 10; i++ { 
		for j := 1; j <= 10; j++ {
			a = append(a, i)
		}
	}

	temp := make(map[int]bool)

	var new []int
	
	for _, val := range a {
		if !temp[val] {
			temp[val] = true   
			new = append(new, val)
		}
	}

	fmt.Printf("去重后切片：%v\n", new) 
}

func main() {
	library := Library{//两种定义方式
		num:4,
		book:map[int]Book{
			0:{
				"西游记",
				"2012年1月1日",
				"吴承恩",
			},
			1:{
				"红楼梦",
				"1995年1月1日",
				"曹雪芹",
			},
			2:{
				"三国演义",
				"2010年1月1日",
				"罗贯中",
			},
			3:{
				"水浒传",
				"1524年-1530年",
				"施耐庵",
			},
		},
	}
    library.ShowAllBook()
	library.AddBook("我是天才","2026年1月2日","我")
	library.ShowAllBook()
	library.RemoveBook(2)
	library.ShowAllBook()
	Solve()
}



