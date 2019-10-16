package main

import "fmt"
import (
	merklebtree "github.com/bradyjoestar/merklebtree-go"
)

func main() {
	println("test")
	hashmap := make(map[string]string)
	hashmap["wenbin"] = "123"

	value, existed := hashmap["wenbin1"]
	if existed {
		fmt.Println(existed)
		fmt.Println(value)
	} else {
		fmt.Println(existed)
		fmt.Println(value)
	}

	mbtree := merklebtree.NewMBTree()

	mbtree.BuildWithKeyValue(merklebtree.KeyVersion{Key: "wenbin", Version: 100})


	mbtree.BuildWithKeyValue(merklebtree.KeyVersion{Key: "wenbin123", Version: 1001})


}
