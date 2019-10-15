package main

import "fmt"

func main()  {
	println("test")
	hashmap := make(map[string]string)
	hashmap["wenbin"] = "123"

	value,existed := hashmap["wenbin1"]
	if existed{
		fmt.Println(existed);
		fmt.Println(value);
	}else{
		fmt.Println(existed);
		fmt.Println(value);
	}
}
