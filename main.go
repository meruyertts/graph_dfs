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
	recordInput()
}

func recordInput() {
	var size, node int
	var myMtrx [][]int
	fmt.Scanf("%d %d\n", &size, &node)
	count := 0
	for count < size {
		var myRow []int
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		for _, char := range strings.Split(scanner.Text(), " ") {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Println(err.Error())
				return
			}
			myRow = append(myRow, num)
		}
		myMtrx = append(myMtrx, myRow)
		count++
	}
	fmt.Println(dfs(myMtrx, node))
}

var visited []int

func dfs(myMtrx [][]int, node int) int {
	stuff := compNode(myMtrx, myMtrx[node-1], node)
	return stuff
}

func compNode(myMtrx [][]int, myArr []int, node int) int {
	visited = append(visited, node)
	count := 0
	for i, val := range myArr {
		if val == 1 && !nodeExists(i+1) {
			count += compNode(myMtrx, myMtrx[i], i+1)
		}
	}
	count++
	return count
}

func nodeExists(node int) bool {
	for _, val := range visited {
		if val == node {
			return true
		}
	}
	return false
}
