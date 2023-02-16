package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to MyTime in GoLang")
	presentTime := time.Now()

	fmt.Println("Present time: ", presentTime)

	// 01-02-2006 15:04:05 Monday -> Default date, for date formatting
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.March, 9, 0, 0, 0, 0, time.UTC)

	fmt.Println("Created date ", createdDate.Format("01-02-2006 15:04:05 Monday"))

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Read string", input)

}
