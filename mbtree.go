package merklebtree_go

type MerkleBtree struct {
	hashmap map[string]int64
	Root
}

type Root struct {
	Hash []byte
}

func (mbtree *MerkleBtree) BuildWithKeyValue(kv KeyVersion) {
	mbtree.hashmap[kv.Key] = kv.Version
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

	btree := MerkleBtree{hashmap: make(map[string]int64),Root:Root{Hash:nil}}
	return &btree
}
