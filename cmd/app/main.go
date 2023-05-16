package main

import cg "app/internal/chessgame"

func main() {
	chess := cg.NewChess()

	cg.PlayGame(chess)
}
