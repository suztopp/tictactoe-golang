package main

import "fmt"

func main() {

	// init the game board with empty strings

	xoBoard := [3][3]string{}

	// var to carry the current player name

	player := "x"

	for {
		fmt.Println("player", player)

		// read row value

		var row int
		fmt.Println("Please enter a row number - choose either 0, 1, or 2 - Thank you")
		fmt.Scanln(&row)

		if row > 2 || row < 0 {
			fmt.Println("Invalid entry for row, please enter new number, either 0, 1 or 2 only")
			continue
		}

		// read column value

		var column int
		fmt.Println("Please enter a column number - choose either 0, 1 or 2 - Thank you")
		fmt.Scanln(&column)

		if column > 2 || column < 0 {
			fmt.Println("Invalid entry for column, please enter a new number, either 0, 1 or 2 only")
			continue
		}

		// set value into game board
		if xoBoard[row][column] == "" {
			xoBoard[row][column] = player
		} else {
			// index already has a value
			fmt.Println("Invalid Entry :", row, column, "value", xoBoard[row][column])
			continue
		}

		// display the game board after each move
		fmt.Println(xoBoard[0])
		fmt.Println(xoBoard[1])
		fmt.Println(xoBoard[2])

		// did someone win

		for i := 0; i < 3; i++ {
			// check row || column
			if xoBoard[i][0] == xoBoard[i][1] && xoBoard[i][1] == xoBoard[i][2] && xoBoard[i][2] == player ||
				xoBoard[0][i] == xoBoard[1][i] && xoBoard[1][i] == xoBoard[2][i] && xoBoard[2][i] == player {
				fmt.Println("Game Over - Winner is Player: ", player)
				// need to exit out of loop and game
				return
			}
		}

		if xoBoard[0][0] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][2] && xoBoard[2][2] == player ||
			xoBoard[0][2] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][0] && xoBoard[2][0] == player {
			fmt.Println("Game Over - Winner is Player: ", player)
			return
		}

		// swap players
		if player == "x" {
			player = "o"
		} else {
			player = "x"
		}
	}

}
