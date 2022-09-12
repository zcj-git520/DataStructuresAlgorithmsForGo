package Trie

import (
	"fmt"
	"testing"
)

func TestInsertData(t *testing.T)  {
	trie := InitTrie()
	var testStr = []string{"zc", "zzc", "zcj","rr","好的", "zdd", "zcf", "找车估计","张伟"}
	for _, str := range testStr{
		err := trie.insertData(str)
		if err != nil {
			t.Error(err)
		}
	}
	//trie.Clear()
	dataS, err := trie.traverseData()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(dataS)
	ok := trie.FindData("zcj1")
	if !ok {
		fmt.Println("zcj is not  exist" )
	}else{
		fmt.Println("zcj is exist" )
	}
	//err = trie.DeleteData("rr")
	//if err != nil {
	//	t.Error(err)
	//}
	dataS, err = trie.traverseData()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(dataS)
	fmt.Println("****************************")
	dataS, err = trie.TrieData("z")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(dataS)

}
