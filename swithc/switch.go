package main

import (
"fmt"
"time"
)

func main(){
i:=2
fmt.Println("Write ", i, " as ")
switch i{
case 1:
fmt.Println("one")
case 2:
fmt.Println("two")
case 3:
fmt.Println("three")
}

switch time.Now().Weekday() {
case time.Saturday, time.Sunday :
fmt.Println("its weekend")
default :
fmt.Println("its weekday")
}

t := time.Now()
switch {
case t.Hour()<12:
fmt.Println("it's before noon")
default:
fmt.Println("its afternoon")
}

whatAmI :=  func(i interface{}){
switch i.(type){
case bool:
fmt.Println("this is a bool")
case string:
fmt.Println("THis is a string")
case int:
fmt.Println("This is an integer")
}
}
whatAmI(true)
whatAmI("string")
whatAmI(23)
}
