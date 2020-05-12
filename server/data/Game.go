package data

import (
	"github.com/pmccau/TriviaGo/server/datamanagement"
)

type Game struct {
	GUID 				string
	Teams				[]Team
	StartTime			int
	Round				int
	RoundsCompleted		int
	MaxRounds			int
	TieBreaker			bool
	Questions			[][]Question
	Lobby				Lobby
}

// Constructor for Game
func NewGame(l Lobby) *Game {
	g := new(Game)
	//g.Teams = Team
	g.GUID = datamanagement.GenerateGuid()
	g.Round = 1
	g.RoundsCompleted = 0
	//g.Questions = QueryQuestions(l)
	return g
}

//func QueryQuestions(Lobby) []Question {
//	return make([]Question, )
//}

