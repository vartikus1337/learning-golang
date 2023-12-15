package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file1, _ := os.Create("text.txt")
	file1.WriteString("1 строка \n")
	info1, _ := file1.Stat()
	fmt.Printf("Size: %d\n", info1.Size())
	os.Rename("text.txt", "new_text.txt")
	os.Remove("new_text.txt")

	var sum int
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        num, _  := strconv.Atoi(s.Text())
        sum += num 
    }
    io.WriteString(os.Stdout, strconv.Itoa(sum))

}

func anotherSolution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var sum int
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		num, _ := strconv.Atoi(string(line))
		sum += num
	}

	writer.WriteString(strconv.Itoa(sum))
	writer.Flush()
}


