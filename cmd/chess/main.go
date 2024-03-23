package main

import "github.com/mrumyantsev/chess/internal/chessgame"

func main() {
	chess := chessgame.New()

	chess.Start()
}
