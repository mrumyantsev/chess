package chessgame

const (
	SideLen            = 8
	RecordsLen         = 8
	PressEnterContinue = "<Enter>\b\b\b\b"
	BoardChars         = " ABCDEFGH"
	WhitePlayer        = "White"
	BlackPlayer        = "Black"
	Space              = " "
	PawnPiece          = "Pawn"
	KnightPiece        = "Knight"
	BishopPiece        = "Bishop"
	RookPiece          = "Rook"
	QueenPiece         = "Queen"
	KingPiece          = "King"
	PlayerColorMessage = "%s turn:\n"

	WPawn   = '♙'
	WKnight = '♘'
	WBishop = '♗'
	WRook   = '♖'
	WQueen  = '♕'
	WKing   = '♔'
	BPawn   = '♟'
	BKnight = '♞'
	BBishop = '♝'
	BRook   = '♜'
	BQueen  = '♛'
	BKing   = '♚'
	NoPiece = ' '
)

func DefaultBoardLayout() [SideLen][SideLen]rune {
	return [SideLen][SideLen]rune{
		{WRook, WKnight, WBishop, WKing, WQueen, WBishop, WKnight, WRook},
		{WPawn, WPawn, WPawn, WPawn, WPawn, WPawn, WPawn, WPawn},
		{NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece},
		{NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece},
		{NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece},
		{NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece},
		{BPawn, BPawn, BPawn, BPawn, BPawn, BPawn, BPawn, BPawn},
		{BRook, BKnight, BBishop, BKing, BQueen, BBishop, BKnight, BRook},
	}
}
