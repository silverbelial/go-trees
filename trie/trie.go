package trie

const (
	HARD_CODE_ARRAY = iota
	LINKED_LIST

)

// Trie interface
type Trie interface {
	Initialize() error
	Insert(BitTrail, interface{})
	Search(BitTrail, bool) interface{}
}


//type VSTTrie struct {
//	ValueSetSize	int
//	RootNode 	*vSTTrieNode
//}
//
//type vSTTrieNode struct {
//	Value nodes.GeneralNode
//	childNodes []*vSTTrieNode
//}
//
//func(vst *VSTTrie) Initialize(size, depth int) error {
//	vst.RootNode = new(vSTTrieNode)
//	vst.ValueSetSize = size
//	//return vst.RootNode.Initialize(size, depth)
//	return nil
//}

