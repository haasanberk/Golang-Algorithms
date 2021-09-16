package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findA() {
	inputArray := [11]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}	
	var withA []string
	var withoutA []string

	for i := 0; i < len(inputArray); i++ {
		if strings.ContainsAny(inputArray[i],"a"){
			withA= append(withA,inputArray[i])
		}else{
			withoutA= append(withoutA,inputArray[i])
		}
	}

	withAMap := make(map[string]int)
	for i := 0; i < len(withA); i ++ {
		withAMap[withA[i]] = strings.Count(withA[i], "a")
	}
	
	fmt.Println(withAMap)
	type kv struct {
        Key   string
        Value int
    }

    var sortingWithA []kv
	
    for k, v := range withAMap {
        sortingWithA = append(sortingWithA, kv{k, v})
    }

    sort.Slice(sortingWithA, func(i, j int) bool {
        return sortingWithA[i].Value > sortingWithA[j].Value
    })

	fmt.Println(sortingWithA)
	
	arr := []string{}
	for i := 0; i < len(sortingWithA); i++ { 
	   arr = append(arr, sortingWithA[i].Key)
	}

	fmt.Println(arr)

	sort.Slice(withoutA, func(i, j int) bool {
        return len(withoutA[i]) > len(withoutA[j])
    })


	for i := 0; i < len(withoutA); i++ { 
		arr = append(arr, withoutA[i])
	 }

	

	 fmt.Println("başlaa")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}








func toBinary(input int) string{
	last := '0';
	first := '0';
	
	var a = (string(first)+(strconv.FormatInt(int64(input), 2))+string(last)) // 1111011
	return a
}
func toDecimal(binary string) int{
 	output, err := strconv.ParseInt(binary, 2, 64)  
	 if err != nil {
 
		panic(err)
	 
	  }

  	return int(output)  
 	
}

func circularShift ( input string ) string {
    var strFirstChar string = input[:1]

	var newStr string = (input[1:]+strFirstChar)	
	return newStr 
}




func recursive(input string){
	var binarySayı string = input

	fmt.Println(binarySayı)

	a:=circularShift(binarySayı)
	toShow := a[1:5]

	fmt.Println(toShow)

	fmt.Println(toDecimal(toShow) )
	fmt.Println(a )
	
	a=circularShift(a)

	toShow = a[1:5]

	fmt.Println(toShow)

	fmt.Println(toDecimal(toShow) )

	fmt.Println(a )
	
	a=circularShift(a)

	toShow = a[1:5]

	fmt.Println(toShow)

	fmt.Println(toDecimal(toShow) )
}




func printslice(slice []string) {
	fmt.Println("slice = ", slice)
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

func main() {
	//findA();  
	//fmt.Println(toBinary(9))
	
	//recursive(toBinary(9))
	duplicate := []string{"apple","pie","apple","red","red","red"}


 	dup_map := dup_count(duplicate)

 	fmt.Println(dup_map)
	fmt.Println(rankByWordCount(dup_map)[0].Key)
 
	 
}
