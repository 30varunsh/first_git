package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type da struct {
	name string
	ti   int64
}

func main() {
	c := make(chan string)
	t := make(chan int64)

	go count(c, t)
	da1 := da{
		name: <-c, ti: <-t,
	}
	fmt.Printf("%s %v", da1.name, da1.ti)
	fmt.Println(da1)
	data, err := json.Marshal(da1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	// for i:=0;i<=3;i++{
	// 	d[i]=data{
	// 		name:<-c, ti:<-t,
	// 	}

	// for i:=0;i<=3;i++{
	// 	fmt.Printf("%s %v",d[i].name,d[i].ti)

	// }

}

func count(c chan string, t chan int64) {
	c <- "varun sharma"
	start := time.Now()
	fmt.Println("Hello World")
	elapsed := time.Since(start)
	t <- elapsed
}
