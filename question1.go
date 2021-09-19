package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	findA();  
}


func findA() {
	inputArray := [11]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	//split arrays by whether they contain 'a'
	var withA []string
	var withoutA []string

	for i := 0; i < len(inputArray); i++ {
		if strings.ContainsAny(inputArray[i],"a"){
			withA= append(withA,inputArray[i])
		}else{
			withoutA= append(withoutA,inputArray[i])
		}
	}

	//create map and insert the number a 
	withAMap := make(map[string]int)
	for i := 0; i < len(withA); i ++ {
		withAMap[withA[i]] = strings.Count(withA[i], "a")
	}
	

    //sort by the number of a it contains
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

	//sort into array
	arr := []string{}
	for i := 0; i < len(sortingWithA); i++ { 
	   arr = append(arr, sortingWithA[i].Key)
	}

	//sort those that don't contain 'a' by length
	sort.Slice(withoutA, func(i, j int) bool {
        return len(withoutA[i]) > len(withoutA[j])
    })

	//add those that don't contain a after those that contain a
	for i := 0; i < len(withoutA); i++ { 
		arr = append(arr, withoutA[i])
	 }

	//print array
	
	fmt.Println(arr)

}

