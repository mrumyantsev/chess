package chessgame

func (c *Chess) isPiece2(a ...byte) bool {
	for _, val := range a {
		if c.Board[c.J2][c.I2] == val {
			return true
		}
	}

	return false
}

func (c *Chess) isCanMoveWPawn() bool {
	if c.isPiece2(NoPiece) && (c.I1 == c.I2) {
		if (c.J1 + 1) == c.J2 {
			return true
		}
		if ((c.J1 + 2) == c.J2) && (c.J1 == 1) {
			return true
		}
	}
	if !c.isPiece2(NoPiece, WPawn, WKnight, WBishop, WRook, WQueen, WKing) && (c.J1+1 == c.J2) {
		if (c.I1 + 1) == c.I2 {
			return true
		}
		if (c.I1 - 1) == c.I2 {
			return true
		}
	}

	return false
}

func (c *Chess) isCanMoveBPawn() bool {
	if c.isPiece2(NoPiece) && (c.I1 == c.I2) {
		if (c.J1 - 1) == c.J2 {
			return true
		}
		if ((c.J1 - 2) == c.J2) && (c.J1 == 6) {
			return true
		}
	}
	if !c.isPiece2(NoPiece, BPawn, BKnight, BBishop, BRook, BQueen, BKing) && (c.J1-1 == c.J2) {
		if (c.I1 + 1) == c.I2 {
			return true
		}
		if (c.I1 - 1) == c.I2 {
			return true
		}
	}

	return false
}
