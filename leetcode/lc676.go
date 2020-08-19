package main

import "fmt"

func main() {

	input := []string{"a", "b", "ab", "abc", "abcabacbababdbadbfaejfoiawfjaojfaojefaowjfoawjfoawj", "abcdefghijawefe", "aefawoifjowajfowafjeoawjfaow", "cba", "cas", "aaewfawi", "babcda", "bcd", "awefj"}
	dick := Constructor()
	dick.BuildDict(input)
	ret := dick.Search("abcd")
	fmt.Println(ret)
}

type MagicDictionary struct {
	end   bool
	nodes [26]*MagicDictionary
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, wd := range dict {
		tail := this
		for idx, ch := range wd {
			if tail.nodes[ch-'a'] == nil {
				tail.nodes[ch-'a'] = new(MagicDictionary)
			}
			tail = tail.nodes[ch-'a']
			if idx == len(wd)-1 {
				tail.end = true
			}
		}
	}
}

func (this *MagicDictionary) Search(word string) bool {
	for idx, n := range this.nodes {
		if n == nil {
			continue
		}
		mismatchPos := -1
		fmt.Printf("%+v \n", this)
		if true == dfsSearch(this.nodes[idx], idx, word, 0, mismatchPos) {
			return true
		}
		fmt.Printf("start over\n")
	}
	return false
}

func dfsSearch(level *MagicDictionary, nodeIdx int, word string, wdIdx int, mismatched int) bool {
	fmt.Printf("try to find %d, cadidates are:", word[wdIdx]-'a')
	for idx, otherPath := range level.nodes {
		if otherPath == nil {
			continue
		}
		fmt.Printf(" %d:%t /", idx, otherPath.end)
	}
	fmt.Printf(" foundMismatch:%d \n", mismatched)
	//fmt.Printf("\n%+v %d - %d %t\n", level, nodeIdx, wdIdx, mismatched)
	if nodeIdx != int(word[wdIdx]-'a') {
		if mismatched != -1 && mismatched != wdIdx {
			// end this route
			return false
		}
		mismatched = wdIdx
	}
	if wdIdx == len(word)-1 {
		return level.end == true && mismatched != -1
	}
	//DFS all the sub paths
	for idx, otherPath := range level.nodes {
		if otherPath == nil {
			continue
		}

		if true == dfsSearch(otherPath, idx, word, wdIdx+1, mismatched) {
			return true
		}

		fmt.Printf("roll back\n")
	}

	return false
}
