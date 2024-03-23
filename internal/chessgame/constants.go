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

	WPawn   = 'p'
	WKnight = 'k'
	WBishop = 'b'
	WRook   = 'r'
	WQueen  = 'q'
	WKing   = 'c'
	BPawn   = 'P'
	BKnight = 'K'
	BBishop = 'B'
	BRook   = 'R'
	BQueen  = 'Q'
	BKing   = 'C'
	NoPiece = ' '
)

func DefaultBoardLayout() [SideLen][SideLen]byte {
	return [SideLen][SideLen]byte{
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
