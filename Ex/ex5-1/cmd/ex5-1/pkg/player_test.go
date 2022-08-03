package pkg

import (
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestValidate(t *testing.T) {

	FirstName := "Rafael"
	LastName := "Rafael"
	Sport := "Cricket"
	Birthday :=  time.Date(1991, time.October, 10, 9, 12, 0, 0, time.UTC) 
	Genre := "M"
	TeamName:= "Junior"

	today := time.Now()
	tomorrow := today.Add(24 * time.Hour)

	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	if strings.TrimSpace(FirstName) == "" {
		t.Errorf("The first name was expected not to be a white space")
	}

	if !IsLetter(FirstName) {
		t.Errorf("First Name can't contain numbers or special characters")
	}

	if strings.TrimSpace(LastName) == "" {
		t.Errorf("The last name was expected not to be a white space")
	}

	if !IsLetter(LastName) {
		t.Errorf("Last Name can't contain numbers or special characters")
	}
	if strings.TrimSpace(Sport) == "" {
		t.Errorf("The sport name was expected not to be a white space")
	}

	if !IsLetter(Sport) {
		t.Errorf("Sport name can't contain numbers or special characters")
	}
	if Birthday.After(tomorrow) {
		t.Errorf("Birthday Must be a date and can't be a date in the future")
	}
	if Genre != "M" && Genre != "F" {
		t.Errorf("Gender can be only M or F")
	}
	if strings.TrimSpace(TeamName) == "" {
		t.Errorf("Team name can't be blank")
	}

	if !isLetter(TeamName) {
		t.Errorf("Team Name can't contain special characters")
	}
}
