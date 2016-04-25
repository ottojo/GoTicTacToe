package main
import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	board := emptyBoard()
	player := true;
	for winner(board) == ""{
		print("\033[H\033[2J")
		drawBoard(board)
		if(player) {
			board = input(board, "x")
		}else {
			board = input(board, "o")
		}
		player = !player
	}
	drawBoard(board)
	println("Winner is: " + winner(board))
}

func emptyBoard() ([][]string){
	board := [][]string{}
	row1 := []string{" ", " ", " "}
	row2 := []string{" ", " ", " "}
	row3 := []string{" ", " ", " "}
	board = append(board, row1)
	board = append(board, row2)
	board = append(board, row3)
	return board
}

func drawBoard(board [][]string) {
	println("    0   1   2    x")
	println("  -------------")
	println("0 | " + board[0][0] + " | " + board[0][1] + " | " + board[0][2] + " |")
	println("  -------------")
	println("1 | " + board[1][0] + " | " + board[1][1] + " | " + board[1][2] + " |")
	println("  -------------")
	println("2 | " + board[2][0] + " | " + board[2][1] + " | " + board[2][2] + " |")
	println("  -------------")
	println("y")
}

func winner(board [][]string) (string) {
	if (testRows(board) == "x" || testColumns(board) == "x" || testDiagonal(board) == "x") {
		return "x"
	} else if (testRows(board) == "o" || testColumns(board) == "o" || testDiagonal(board) == "o") {
		return "o"
	}else {
		return ""
	}
}

func testRows(board [][]string) (string) {
	for i := 0; i < 3; i++ {

		if (board[i][0] == "x" && board[i][1] == "x" && board[i][2] == "x") {
			return "x"
		}else if (board[i][0] == "o" && board[i][1] == "o" && board[i][2] == "o") {
			return "o"
		}
	}
	return ""
}

func testColumns(board [][]string) (string) {
	for i := 0; i < 3; i++ {

		if (board[0][i] == "x" && board[1][i] == "x" && board[2][i] == "x") {
			return "x"
		}else if (board[0][i] == "o" && board[1][i] == "o" && board[2][i] == "o") {
			return "o"
		}
	}
	return ""
}

func testDiagonal(board [][]string) (string) {
	if (board[0][0] == "x" && board[1][1] == "x" && board[2][2] == "x") {
		return "x"
	}else if (board[0][0] == "o" && board[1][1] == "o" && board[2][2] == "o") {
		return "o"
	}else if (board[0][2] == "x" && board[1][1] == "x" && board[2][0] == "x") {
		return "x"
	}else if (board[0][2] == "o" && board[1][1] == "o" && board[2][0] == "o") {
		return "o"
	}else {
		return ""
	}
}

func input(board [][]string, player string) ([][]string){
	result := board
	println("Field [x y]: ")
	in := readLine()
	inArray := strings.Split(in, " ")
	x, _ := strconv.Atoi(inArray[0])
	y, _ := strconv.Atoi(inArray[1])
	if(board[y][x] == " "){
		result[y][x] = player
	}
	return result
}

func readLine() (string) {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	return line
}