package entity

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	PlayersIDs  []Player
}
type Player struct {
	ID      uint
	UserID  uint
	GameID  uint
	Score   uint
	Answers []uint
}
