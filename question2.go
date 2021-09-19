package main

import (
	"fmt"
	"strconv"
)

func main() {

initial := int(9)
var recursive func(int)
    recursive = func(input int) {

	number := int64(input)
	//Convert number to binary string
	binary := strconv.FormatInt(number, 2)
	//Add necessary 0s in the beginning if missing
	for i := len(binary); i < 4; i++ {
		binary="0"+binary
	}
	
	//Check the equivalance of the last 2 bits of the binary number
	//if equal then add a "1" in the end of the string
	//if not add a 0 in the end of the string
	if string(binary[2]) == string(binary[3]){
		binary=binary+"1"
	}else{
		binary=binary+"0"
	}

	//Take the last 4 bits of the current binary number as the new binary number (left shift by one)
	newBinary := string(binary[1]) + string(binary[2]) + string(binary[3]) + string(binary[4])
	
	//Back To Decimal
	if newDecimal, err := strconv.ParseInt(newBinary , 2, 64); err != nil {
        	fmt.Println(err)
    	} else {
        	fmt.Println(newDecimal)
		//if we reach to the initial value, the code should stop
		if int(newDecimal) == initial {
		return;
		}
		recursive(int(newDecimal))
    	}
    }
    recursive(initial)
}