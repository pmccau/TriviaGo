/*
Package datamanagement is used for data manipulation:
	- MongoClient: interact with MongoDB to store/retrieve questions
	- Utility: Misc utility functions that didn't have a natural home elsewhere
*/
package datamanagement

import (
	"fmt"
	"crypto/rand"
	)

// GenerateGuid generates a simple GUID
// Credit: https://yourbasic.org/golang/generate-uuid-guid/
func GenerateGuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	Check(err)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return string(uuid)
}

// Check performs simple error checking
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Split function for csv field parsing.
// Credit: I didn't write this, but I can't find the stackoverflow question I pulled it from
func Split(r rune) bool {
	return r == ',' || r == '\n'
}

// GenerateJoinCode will return a 4 digit guid
func GenerateJoinCode() string {
	return GenerateGuid()[:4]
}