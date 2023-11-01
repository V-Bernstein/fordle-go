package types 

type Color int

const (
    Gray Color = iota
    Yellow
    Green
)

type LetterCode struct {
    letter byte
    color Color
}

type GameResponse struct {
    letters []LetterCode
    won bool
}
