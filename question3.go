package main

import (
	"fmt"
	"sort"
)

func main() {

	duplicate := []string{"apple","pie","apple","red","red","red"}

	mostFreq := (rankByWordCount(dup_count(duplicate))[0].Key)

	fmt.Println(mostFreq)
 
	 
}

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	
	return duplicate_frequency
}



func rankByWordCount(wordFrequencies map[string]int) PairList{

	// sort in reverse order by repeat value

	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
	  pl[i] = Pair{k, v}
	  i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
  }
  
  type Pair struct {
	Key string
	Value int
  }
  
  type PairList []Pair
  
  func (p PairList) Len() int { return len(p) }
  func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
  func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }


