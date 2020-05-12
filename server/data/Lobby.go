package data

import (
	"github.com/pmccau/TriviaGo/server/datamanagement"
)

type Lobby struct {
	Teams					[]Team
	GUID					string
	Hosts					[]Team
	Code					string
	Password				string
	Categories				[]string
	StandardRoundsMax		int
	TieBreakRoundsMax		int
	TieBreakEnabled			bool
	CutThroatMode			bool
	Penalty					int
	MaxTeams				int
	QuestionsPerRound		int
}

func NewLobby(host Team, password string) *Lobby {
	l := new(Lobby)
	h := make([]Team, 1)
	h[0] = host
	l.Hosts = h
	l.GUID = datamanagement.GenerateGuid()
	l.Code = datamanagement.GenerateJoinCode()
	l.Password = password
	setDefaultRules(l)
	l.Categories = make([]string, l.StandardRoundsMax + l.TieBreakRoundsMax)
	return l
}

// setDefaultRules will populate the default rules
func setDefaultRules(l *Lobby) {
	l.MaxTeams = 8
	l.Penalty = 0
	l.CutThroatMode = false
	l.StandardRoundsMax = 5
	l.TieBreakRoundsMax = 3
	l.TieBreakEnabled = true
	l.QuestionsPerRound = 10
}

func UpdateRules(l *Lobby) {

}