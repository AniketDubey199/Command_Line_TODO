package main

import (
	"bufio"
	"fmt"
	"os"
)

func insertTask(val string, nums map[int]string, index int) {
	nums[index] = val
	fmt.Println("Task added Successfully")
}

func deleteTask(index int, nums map[int]string) {
	fmt.Println("Deleting: ", nums[index])
	delete(nums, index)
	fmt.Println("Task Deleted Successfully")
}

func printAll(nums map[int]string) {
	for i, k := range nums {
		fmt.Println(i, ".", k)
	}
}

func main() {
	mp := make(map[int]string)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Submit your task")
	s, _ := reader.ReadString('\n')

	index := 1

	insertTask(s, mp, index)
	printAll(mp)
	deleteTask(index, mp)

}
