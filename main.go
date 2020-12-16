package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var gameBoardSize int
	var gameBoard = make(map[int]bool)

	fmt.Println("Enter Board size: ")
	fmt.Scanln(&gameBoardSize)
	println()
	numbersPool := generateBoard(gameBoardSize, gameBoard)
	printBoard(gameBoard)
	for len(numbersPool) > 0 {
		fmt.Println("Press enter to generate a Balote")
		fmt.Scanln()
		balote, newNumbersPool := baloteGenerator(numbersPool, gameBoard)
		result := fmt.Sprintf("Next Balote is: %v", balote)
		fmt.Println(result)
		println()
		printBoard(gameBoard)
		println()
		numbersPool = newNumbersPool
	}
}

func generateBoard(boardSize int, gameBoard map[int]bool) (numbersPool []int) {
	for index := 1; index < boardSize+1; index++ {
		gameBoard[index] = false
		numbersPool = append(numbersPool, index)
	}
	return
}

func printBoard(board map[int]bool) {
	limit := len(board) / 5

	row := "B:  "
	for index := 1; index < limit+1; index++ {
		coloredOutput := formatBalote(board, index)
		row = fmt.Sprintf("%v\t%v", row, coloredOutput)
	}
	fmt.Println(row)
	println()

	row = "I:  "
	for index := limit + 1; index < limit*2+1; index++ {
		coloredOutput := formatBalote(board, index)
		row = fmt.Sprintf("%v\t%v", row, coloredOutput)
	}
	fmt.Println(row)
	println()

	row = "N:  "
	for index := limit*2 + 1; index < limit*3+1; index++ {
		coloredOutput := formatBalote(board, index)
		row = fmt.Sprintf("%v\t%v", row, coloredOutput)
	}
	fmt.Println(row)
	println()

	row = "G:  "
	for index := limit*3 + 1; index < limit*4+1; index++ {
		coloredOutput := formatBalote(board, index)
		row = fmt.Sprintf("%v\t%v", row, coloredOutput)
	}
	fmt.Println(row)
	println()

	row = "O:  "
	for index := limit*4 + 1; index < limit*5+1; index++ {
		coloredOutput := formatBalote(board, index)
		row = fmt.Sprintf("%v\t%v", row, coloredOutput)
	}
	fmt.Println(row)
	println()
}

func formatBalote(board map[int]bool, value int) (number string) {
	number = ""
	if board[value] {
		number = "\033[36m"
	} else {
		number = "\033[33m"
	}
	number += strconv.Itoa(value) + "\033[0m"
	return
}

func baloteGenerator(numbersPool []int, gameBoard map[int]bool) (balote int, newPool []int) {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(numbersPool))
	balote = numbersPool[index]
	gameBoard[balote] = true
	newPool = removeFromSlice(numbersPool, index)
	return
}

func removeFromSlice(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
