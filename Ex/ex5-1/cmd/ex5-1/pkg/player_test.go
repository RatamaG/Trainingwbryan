package pkg

import (
	"testing"
	"time"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		Player  Player
		Valid   bool
		Message string
	}{
		{
			Player: Player{
				FirstName: "Rafael",
				LastName:  "Perez",
				Birthday:  time.Now(),
				Genre:     "M",
				TeamName:  "Chelsea",
				Sport:     "Soccer",
			},
			Valid:   true,
			Message: "",
		},
		{
			Player: Player{
				FirstName: "",
				LastName:  "Perez",
				Sport:     "Soccer",
				Birthday:  time.Now(),
				Genre:     "M",
				TeamName:  "Chelsea",
			},
			Valid:   false,
			Message: "First Name can't be blank",
		},
		{
			Player: Player{
				FirstName: "Rafael@",
				LastName:  "Perez",
				Sport:     "Soccer",
				Birthday:  time.Now(),
				Genre:     "M",
				TeamName:  "Chelsea",
			},
			Valid:   false,
			Message: "First Name can't contain numbers or special characters",
		},
		{
			Player: Player{
				FirstName: "Rafael",
				LastName:  "",
				Sport:     "Soccer",
				Birthday:  time.Now(),
				Genre:     "M",
				TeamName:  "Chelsea",
			},
			Valid:   false,
			Message: "Last Name can't be blank",
		},
		{
			Player: Player{
				FirstName: "Rafael",
				LastName:  "Perez!5",
				Sport:     "Soccer",
				Birthday:  time.Now(),
				Genre:     "M",
				TeamName:  "Chelsea",
			},
			Valid:   false,
			Message: "Last Name can't contain numbers or special characters",
		},
		{
			Player: Player{
				FirstName: "Rafael",
				LastName:  "Pere",
				Sport:     "",
				Birthday:  time.Now(),
				Genre:     "M",
				TeamName:  "Chelsea",
			},
			Valid:   false,
			Message: "Sport name can't be blank",
		},
		{
			Player: Player{
				FirstName: "Rafael",
				LastName:  "Pere",
				Sport:     "Soccer90",
				Birthday:  time.Now(),
				Genre:     "M",
				TeamName:  "Chelsea",
			},
			Valid:   false,
			Message: "Sport name can't contain numbers or special characters",
		},
		{
			Player: Player{
				FirstName: "Rafael",
				LastName:  "Pere",
				Sport:     "Soccer",
				Birthday:  time.Now().Add(time.Hour * 24 * 10),
				Genre:     "M",
				TeamName:  "Chelsea",
			},
			Valid:   false,
			Message: "Birthday Must be a date and can't be a date in the future",
		},
		{
			Player: Player{
				FirstName: "Rafael",
				LastName:  "Pere",
				Sport:     "Soccer",
				Birthday:  time.Now(),
				Genre:     "K",
				TeamName:  "Chelsea",
			},
			Valid:   false,
			Message: "Gender can be only M or F",
		},
		{
			Player: Player{
				FirstName: "Rafael",
				LastName:  "Pere",
				Sport:     "Soccer",
				Birthday:  time.Now(),
				Genre:     "M",
				TeamName:  "",
			},
			Valid:   false,
			Message: "Team name can't be blank",
		},
		{
			Player: Player{
				FirstName: "Rafael",
				LastName:  "Pere",
				Sport:     "Soccer",
				Birthday:  time.Now(),
				Genre:     "M",
				TeamName:  "Chelsea!@",
			},
			Valid:   false,
			Message: "Team Name can't contain special characters",
		},

	}

	for index, tcase := range testCases {
		valid, message := tcase.Player.Validate()

		if valid != tcase.Valid {
			t.Errorf("[Valid] A validation is not well, index: %v", index)
		}

		if message != tcase.Message {
			t.Errorf("[Message] A validation is not well, index: %v", index)
		}
	}
}
