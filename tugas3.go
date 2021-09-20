package main

import "fmt"

func main(){
  var element=[]string{"apel","jeruk","anggur","melon"}
  element=append(element,"pepaya")
  fmt.Println("Jumlah Element : ",len(element))
  fmt.Println("Isi Element :",element)
  for i:=0;i<len(element);i++{
    fmt.Println("Element ke- :",i,element[i])
  }
}
