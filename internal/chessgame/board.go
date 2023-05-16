package chessgame

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type playableGame interface {
	play()
}

func PlayGame(p playableGame) {
	p.play()
}

type chess struct {
	Board     [LenOfSide][LenOfSide]byte
	TurnCount int
	I1        int
	J1        int
	I2        int
	J2        int
	Piece1    byte
	Piece2    byte
	Records   [8]string
}

func NewChess() *chess {
	return &chess{
		Board: [LenOfSide][LenOfSide]byte{
			{WRook, WKnight, WBishop, WKing, WQueen, WBishop, WKnight, WRook},
			{WPawn, WPawn, WPawn, WPawn, WPawn, WPawn, WPawn, WPawn},
			{NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece},
			{NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece},
			{NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece},
			{NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece},
			{BPawn, BPawn, BPawn, BPawn, BPawn, BPawn, BPawn, BPawn},
			{BRook, BKnight, BBishop, BKing, BQueen, BBishop, BKnight, BRook},
		},
	}
}

func (c *chess) play() {
	var input string
	var err error
	var turnTelling string

	for {
		clearScreen()
		c.DrawBoard()
		fmt.Println()
		c.PrintTurnMessage()
		input = ""
		err = nil
		_, err = fmt.Scan(&input)
		fmt.Scanln()

		if err != nil {
			fmt.Println("Error on scanning input:", err)
			fmt.Print(PressEnterContinue)
			fmt.Scanln()
			continue
		}

		if isExitInput(&input) {
			fmt.Println("Have a nice day!")
			os.Exit(0)
		}

		input = strings.ToUpper(input)
		err = checkTurn(&input)

		if err != nil {
			fmt.Println("Error on checking turn:", err)
			fmt.Print(PressEnterContinue)
			fmt.Scanln()
			continue
		}

		c.ConvertTurn(&input)

		if c.IsMoveAvailable() {
			c.MovePiece()
			turnTelling = fmt.Sprintf("%3d-%s. %s", c.TurnCount, input, c.TellAboutTurn())
			c.AddNewRecord(turnTelling)
		} else {
			fmt.Println("Error: unavailable move")
			fmt.Print(PressEnterContinue)
			fmt.Scanln()
			continue
		}
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error on screen clearing:", err)
	}
}

func isExitInput(input *string) bool {
	if *input == "q" || *input == "quit" || *input == "exit" {
		return true
	}

	return false
}

func checkTurn(input *string) error {
	if input == nil {
		return errors.New("nil is accepted")
	}

	if len(*input) > 4 {
		return errors.New("more than 4 chars is entered")
	}

	if len(*input) < 4 {
		return errors.New("less than 4 chars is entered")
	}

	for idx, val := range *input {
		if idx%2 == 0 {
			if !(val >= 'A' && val <= 'H') {
				return errors.New(fmt.Sprintf("input wrong letter, position: %d", idx+1))
			}

			continue
		}

		if !(val >= '1' && val <= '8') {
			return errors.New(fmt.Sprintf("input wrong digit, position: %d", idx+1))
		}
	}

	return nil
}

func (c *chess) DrawBoard() {
	fmt.Println(BoardChars + "    Record Table")
	for j := LenOfSide - 1; j >= 0; j-- {
		for i := 0; i < LenOfSide; i++ {
			if i == 0 {
				fmt.Print(j + 1)
			}
			fmt.Print(string(c.Board[j][i]))
			if i == LenOfSide-1 {
				fmt.Print(j + 1)
				if c.Records[LenOfRecords-j-1] == "" {
					fmt.Print("   ---\n")
				} else {
					fmt.Print("   " + c.Records[LenOfRecords-j-1] + "\n")
				}
			}
			fmt.Print()
		}
	}
	fmt.Println(BoardChars)
}

func (c *chess) PrintTurnMessage() {
	if c.TurnCount%2 == 0 {
		fmt.Printf(PlayerColorMessage, WhitePlayer)
		return
	}
	fmt.Printf(PlayerColorMessage, BlackPlayer)
}

func (c *chess) ConvertTurn(input *string) {
	c.I1 = int((*input)[0]) - 65
	c.J1 = int((*input)[1]) - 49
	c.I2 = int((*input)[2]) - 65
	c.J2 = int((*input)[3]) - 49
}

func (c *chess) TellAboutTurn() string {
	if c.Piece2 == NoPiece {
		return getPieceName(c.Piece1) + " moves"
	}
	return getPieceName(c.Piece1) + " beats " + getPieceName(c.Piece2)
}

func getPieceName(piece byte) string {
	if piece == WPawn {
		return WhitePlayer + Space + PawnPiece
	}
	if piece == WKnight {
		return WhitePlayer + Space + KnightPiece
	}
	if piece == WBishop {
		return WhitePlayer + Space + BishopPiece
	}
	if piece == WRook {
		return WhitePlayer + Space + RookPiece
	}
	if piece == WQueen {
		return WhitePlayer + Space + QueenPiece
	}
	if piece == WKing {
		return WhitePlayer + Space + KingPiece
	}
	if piece == BPawn {
		return BlackPlayer + Space + PawnPiece
	}
	if piece == BKnight {
		return BlackPlayer + Space + KnightPiece
	}
	if piece == BBishop {
		return BlackPlayer + Space + BishopPiece
	}
	if piece == BRook {
		return BlackPlayer + Space + RookPiece
	}
	if piece == BQueen {
		return BlackPlayer + Space + QueenPiece
	}
	if piece == BKing {
		return BlackPlayer + Space + KingPiece
	}
	return "Error piece"
}

func (c *chess) AddNewRecord(recordLine string) {
	var isArrayFull bool = true
	for i := 0; i < LenOfRecords; i++ {
		if c.Records[i] == "" {
			isArrayFull = false
			break
		}
	}
	if isArrayFull {
		for i := 1; i < LenOfRecords; i++ {
			c.Records[i-1] = c.Records[i]
		}
		c.Records[LenOfRecords-1] = recordLine
		return
	}
	for i := 0; i < LenOfRecords; i++ {
		if c.Records[i] == "" {
			c.Records[i] = recordLine
			break
		}
	}
}
