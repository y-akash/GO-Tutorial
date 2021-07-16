package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	// here we will get float number in string format
	floatVar, _ := reader.ReadString('\n')
	// to convert that string into float
	// strconv.ParseFloat(variableWantToConvert, bit)
	// this method will retun two value one is convertable float and second one is error
	// if someone pass the alphabet instead of number so it will give error
	// bit for you want 64 or 32 bit float
	floatVar1, err := strconv.ParseFloat(strings.TrimSpace(floatVar), 64)
	// fmt.Printf("error =%v\n floatVar1=%v", err, floatVar1)
	fmt.Println("error", err)
	fmt.Println("value", floatVar1)

}
