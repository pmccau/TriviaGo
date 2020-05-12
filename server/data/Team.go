package data

import (
	"github.com/pmccau/TriviaGo/server/datamanagement"
)

// class: Team
// 	+ Score		+ GUID
// 	+ Name		+ Password
type Team struct {
	Name		string
	Guid		string
	Score		int
	Password	string
}

// NewTeam returns a new Team
func NewTeam(Name string, Password string) *Team {
	t := new(Team)
	t.Name = Name
	t.Password = Password
	t.Guid = datamanagement.GenerateGuid()
	return t
}