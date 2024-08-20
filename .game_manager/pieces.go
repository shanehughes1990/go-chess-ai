package gamemanager

import "github.com/notnil/chess"

// PieceImage is a map of GamePiece to the image file path.
var PieceImage = map[chess.Piece]string{
	chess.WhitePawn:   "assets/pawn.png",
	chess.WhiteRook:   "assets/rook.png",
	chess.WhiteKnight: "assets/knight.png",
	chess.WhiteBishop: "assets/bishop.png",
	chess.WhiteQueen:  "assets/queen.png",
	chess.WhiteKing:   "assets/king.png",
	chess.BlackPawn:   "assets/pawn1.png",
	chess.BlackRook:   "assets/rook1.png",
	chess.BlackKnight: "assets/knight1.png",
	chess.BlackBishop: "assets/bishop1.png",
	chess.BlackQueen:  "assets/queen1.png",
	chess.BlackKing:   "assets/king1.png",
}
