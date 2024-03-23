package chessgame

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	isWindows = runtime.GOOS == "windows"
)

type Chess struct {
	Board     [SideLen][SideLen]byte
	TurnCount int
	I1        int
	J1        int
	I2        int
	J2        int
	Piece1    byte
	Piece2    byte
	Records   [RecordsLen]string
}

func New() *Chess {
	return &Chess{Board: DefaultBoardLayout()}
}

func (c *Chess) Start() {
	var (
		input    string
		err      error
		turnInfo string
	)

	for {
		clearScreen()
		c.drawBoard()
		c.printTurnMessage()

		input = ""

		if _, err = fmt.Scanln(&input); err != nil {
			fmt.Println("Error on scanning input:", err)
			fmt.Print(PressEnterContinue)
			fmt.Scanln()

			continue
		}

		if isExitInput(input) {
			fmt.Println("Goodbye!")

			os.Exit(0)
		}

		input = strings.ToUpper(input)

		if err = checkTurn(input); err != nil {
			fmt.Println("Error on checking turn:", err)
			fmt.Print(PressEnterContinue)
			fmt.Scanln()

			continue
		}

		c.convertTurn(input)

		if c.isMoveAvailable() {
			c.movePiece()

			turnInfo = fmt.Sprintf("%3d-%s. %s", c.TurnCount, input, c.turnInfo())

			c.addRecord(turnInfo)
		} else {
			fmt.Println("Error: unavailable move")
			fmt.Print(PressEnterContinue)
			fmt.Scanln()

			continue
		}
	}
}

func (c *Chess) isMoveAvailable() bool {
	if (c.TurnCount % 2) == 0 {
		if c.Board[c.J1][c.I1] == WPawn {
			if c.isCanMoveWPawn() {
				return true
			}
		}
	} else {
		if c.Board[c.J1][c.I1] == BPawn {
			if c.isCanMoveBPawn() {
				return true
			}
		}
	}

	return false
}

func (c *Chess) movePiece() {
	c.Piece1 = c.Board[c.J1][c.I1]
	c.Piece2 = c.Board[c.J2][c.I2]
	c.Board[c.J2][c.I2] = c.Piece1
	c.Board[c.J1][c.I1] = NoPiece
	c.TurnCount++
}

func clearScreen() {
	var cmd *exec.Cmd

	if isWindows {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("Error on screen clearing:", err)
	}
}

func isExitInput(input string) bool {
	if (input == "q") || (input == "quit") || (input == "exit") {
		return true
	}

	return false
}

func checkTurn(input string) error {
	if len(input) > 4 {
		return errors.New("more than 4 chars is entered")
	}

	if len(input) < 4 {
		return errors.New("less than 4 chars is entered")
	}

	for idx, val := range input {
		if idx%2 == 0 {
			if !(val >= 'A' && val <= 'H') {
				return fmt.Errorf("input wrong letter, position: %d", idx+1)
			}

			continue
		}

		if !(val >= '1' && val <= '8') {
			return fmt.Errorf("input wrong digit, position: %d", idx+1)
		}
	}

	return nil
}

func (c *Chess) drawBoard() {
	fmt.Println(BoardChars + "    Record Table")

	for j := SideLen - 1; j >= 0; j-- {
		for i := 0; i < SideLen; i++ {
			if i == 0 {
				fmt.Print(j + 1)
			}

			fmt.Print(string(c.Board[j][i]))

			if i == SideLen-1 {
				fmt.Print(j + 1)

				if c.Records[RecordsLen-j-1] == "" {
					fmt.Print("   ---\n")
				} else {
					fmt.Print("   " + c.Records[RecordsLen-j-1] + "\n")
				}
			}

			fmt.Print()
		}
	}

	fmt.Println(BoardChars)
	fmt.Println()
}

func (c *Chess) printTurnMessage() {
	if (c.TurnCount % 2) == 0 {
		fmt.Printf(PlayerColorMessage, WhitePlayer)

		return
	}

	fmt.Printf(PlayerColorMessage, BlackPlayer)
}

func (c *Chess) convertTurn(input string) {
	c.I1 = int((input)[0]) - 65
	c.J1 = int((input)[1]) - 49
	c.I2 = int((input)[2]) - 65
	c.J2 = int((input)[3]) - 49
}

func (c *Chess) turnInfo() string {
	if c.Piece2 == NoPiece {
		return pieceName(c.Piece1) + " moves"
	}

	return pieceName(c.Piece1) + " beats " + pieceName(c.Piece2)
}

func pieceName(piece byte) string {
	switch piece {
	case WPawn:
		return WhitePlayer + Space + PawnPiece
	case WKnight:
		return WhitePlayer + Space + KnightPiece
	case WBishop:
		return WhitePlayer + Space + BishopPiece
	case WRook:
		return WhitePlayer + Space + RookPiece
	case WQueen:
		return WhitePlayer + Space + QueenPiece
	case WKing:
		return WhitePlayer + Space + KingPiece
	case BPawn:
		return BlackPlayer + Space + PawnPiece
	case BKnight:
		return BlackPlayer + Space + KnightPiece
	case BBishop:
		return BlackPlayer + Space + BishopPiece
	case BRook:
		return BlackPlayer + Space + RookPiece
	case BQueen:
		return BlackPlayer + Space + QueenPiece
	case BKing:
		return BlackPlayer + Space + KingPiece
	default:
		return "Error piece"
	}
}

func (c *Chess) addRecord(recordLine string) {
	isArrayFull := true

	for i := 0; i < RecordsLen; i++ {
		if c.Records[i] == "" {
			isArrayFull = false

			break
		}
	}

	if isArrayFull {
		for i := 1; i < RecordsLen; i++ {
			c.Records[i-1] = c.Records[i]
		}

		c.Records[RecordsLen-1] = recordLine

		return
	}

	for i := 0; i < RecordsLen; i++ {
		if c.Records[i] == "" {
			c.Records[i] = recordLine

			break
		}
	}
}
