package chessgame

func (c *chess) IsMoveAvailable() bool {
	if c.TurnCount%2 == 0 {
		if c.Board[c.J1][c.I1] == WPawn {
			if c.IsCanMoveWPawn() {
				return true
			}
		}
	} else {
		if c.Board[c.J1][c.I1] == BPawn {
			if c.IsCanMoveBPawn() {
				return true
			}
		}
	}
	return false
}

func (c *chess) MovePiece() {
	c.Piece1 = c.Board[c.J1][c.I1]
	c.Piece2 = c.Board[c.J2][c.I2]
	c.Board[c.J2][c.I2] = c.Piece1
	c.Board[c.J1][c.I1] = NoPiece
	c.TurnCount++
}
