package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// Default scanner is bufio.ScanLines. Lets use ScanWords.
	scanner.Split(bufio.ScanWords)

	success := scanner.Scan()
	if success == false {
		// False on error or EOF. Check error
		err = scanner.Err()
		if err == nil {
			// reached EOF
			//log.Println("Scan completed and reached EOF")
			log.Fatal("Didn't find a line in input.txt")
		} else {
			log.Fatal(err)
		}
	}
	var line = scanner.Text()
	fmt.Println("line = ", line)
	fmt.Println("\n")

	// Split on comma.
	slice := strings.Split(line, ",")

	// slice := strings.Split(scanner.Text(), ",")

	// fmt.Println(slice)

	var num []int
	fmt.Println("num", num)
	fmt.Println("\n")

	// loop through list of strings
	for _, s := range slice {
		// converting to numbers
		n, _ := strconv.Atoi(s)
		// append number to list
		num = append(num, n)
	}

	// OPCODES:
	// 99 	STOP
	// 1	ADD
	// 2	MUL

	// num[1] = 12
	// num[2] = 2

	// loop through numbers
	for i := 0; ; i += 4 {

		a := num[num[i+1]] // num[loc 1] = num 1
		b := num[num[i+2]] // num[loc 2] = num 2

		// check opcode
		if num[i] == 1 {
			// run func 1

			num[num[i+3]] = a + b // num[dest] = num 1 + num 2
		} else if num[i] == 2 {
			// run func 2
			num[num[i+3]] = a * b
		} else if num[i] == 3 {
			// opcode 3
			// should move the number 2 back because we only use 2 not 4
			// takes one input and saves it at the the address
			num[num[i+1]] = num[i+1]
			i-=2

		} else if num[i] == 4 {
			fmt.Println( num[i+1])
			// should move the number 
			// opcode 4


		} else if num[i] == 99 {

			break
		} else {
			fmt.Println("ERROR")
			fmt.Println(i)
			break
		}
	}
	fmt.Println(num[0])
	// fmt.Println("\n")
	// Display all elements.
	for true {
		//fmt.Println(code[codeIndex])
	}

}



func intCodeComputer(){
	// position in memory == address
	// value add
	//OPCODES: indicates what to do 
	// 1: 3 inputs 
	// 2: 3 inputs
	//  99: no inputs 
	// 3, 4
	//VERBS 
	//NOUNS
	//PARAMETER MODES:
	// 0,1
}