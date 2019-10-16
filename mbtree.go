package merklebtree_go

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type MerkleBtree struct {
	hashmap map[string]int64
	Root
}

type Root struct {
	Hash []byte
}

func (mbtree *MerkleBtree) BuildWithKeyValue(kv KeyVersion) {
	var s string
	mbtree.hashmap[kv.Key] = kv.Version
	for key, _ := range mbtree.hashmap {
		signByte := []byte(key)
		hash := md5.New()
		hash.Write(signByte)
		s = s + hex.EncodeToString(hash.Sum(nil))
	}
	signByte := []byte(s)
	hash := md5.New()
	hash.Write(signByte)
}

func (mbtree *MerkleBtree) Delete(key string) {
	delete(mbtree.hashmap, key)
}

func (mbtree *MerkleBtree) Serach(key string) SearchResult {
	version, existed := mbtree.hashmap[key]
	return SearchResult{Key: key, Version: version, Existed: existed}
}

type SearchResult struct {
	Key     string
	Version int64
	Existed bool
}

type KeyVersion struct {
	Key     string
	Version int64
}

func NewMBTree() *MerkleBtree {
	btree := MerkleBtree{hashmap: make(map[string]int64), Root: Root{Hash: nil}}
	return &btree
}
