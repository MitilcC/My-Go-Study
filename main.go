package main  

import (
	"encoding/base64"
	"fmt"
)

func bm(){
   i:=0
   var base string

   for i < 2{
	fmt.Println("请输入编码内容 返回back")
	fmt.Scanf("%s",&base)
	if(base == "back"){
		break
	}
    fmt.Println(base64.StdEncoding.EncodeToString([]byte(base)))
   }
}

func jm(){

	i:=0
   var base string

	for i < 2{
		fmt.Println("请输入编码内容 返回back")
		fmt.Scanf("%s",&base)
		if(base == "back"){
			break
		}
		res,err:= base64.StdEncoding.DecodeString(base)
	    if(err != nil){
			fmt.Println("error")
			continue
		}
		fmt.Println(string(res))
	}
}

func main(){ 
	Select := 'n'
	Type := 0;
    i := 0
	for i < 1{
		fmt.Println("编码1 解码2 退出e")
        fmt.Scanf("%c",&Select)
		if(Select == '1'){
           Type = 0
		}else if(Select == '2'){
			Type = 1
		}else if(Select == 'e'){
			break
		}else{
			fmt.Println("error!")
			Type = 2;
		}

		if(Type < 2){
			if(Type == 0){
				bm()
			}else {
				jm()
			}
		}
	}


}
