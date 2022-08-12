package server

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RatamaG/docs/cmd/ex5-1/pkg"
)

func Test_users(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/players", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPlayers)

	handler.ServeHTTP(res, req)
	response := pkg.Response{}
	json.Unmarshal(res.Body.Bytes(), &response)

	if response.Status != http.StatusOK {
		t.Error("expected", http.StatusOK)
	}
	if len(response.Data) != len(pkg.Playerlist) {
		t.Error("expected", response.Status)
	}
	fmt.Println("", response.Status)
	fmt.Println("", response.Data)
}
func Test_user_get_player(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/players/967ef52a-898a-4a83-994f-300bdaaa2bac", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPlayer)

	handler.ServeHTTP(res, req)
	response := pkg.Response{}
	player := pkg.Player{}
	json.Unmarshal(res.Body.Bytes(), &response)

	if player.ID.String() != "967ef52a-898a-4a83-994f-300bdaaa2bac"{
		t.Error("They must be the same id")
	}
	if response.Status != http.StatusOK {
		t.Error("expected", http.StatusOK)
	}
	if len(response.Data) != 1 {
		t.Error("expected", response.Data)
	}
}

// func Test_user_Update(t *testing.T) {

// 	player := []byte(`{ ID:"47f1923a-0227-11ed-b939-0242ac120002",
// 	FirstName: "Ra",
// 	LastName:  "TamaG",
// 	Birthday:  time.Date(1998, time.June, 26, 12, 35, 0, 0, time.UTC),
// 	Genre:     "M",
// 	Sport:     "Cricket",
// 	TeamName:  "Caballeros de Kolkata", }`)

// 	req, err := http.NewRequest(http.MethodPut, "/update/04ef0cad-c366-486b-8a98-8480b7a57442", bytes.NewBuffer(player))

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	res := httptest.NewRecorder()
// 	handler := http.HandlerFunc(UpdatePlayer)

// 	handler.ServeHTTP(res, req)
// 	response := pkg.Response{}
// 	json.Unmarshal(res.Body.Bytes(), &response)

// 	if response.Status != http.StatusOK {
// 		t.Error("expected", http.StatusOK)
// 	}
// 	if response.Message == "" {
// 		t.Error("Message cant be blank")
// 	}
// 	if len(response.Data) != 1 {
// 		t.Error("expected", response.Data)
// 	}
// }

// func Test_user_Delete(t *testing.T) {

// 	req, err := http.NewRequest(http.MethodDelete, "/delete/{id}", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	res := httptest.NewRecorder()
// 	handler := http.HandlerFunc(DeletePlayer)

// 	handler.ServeHTTP(res, req)
// 	response := pkg.Response{}
// 	json.Unmarshal(res.Body.Bytes(), &response)

// 	if response.Status != http.StatusOK {
// 		t.Error("expected", http.StatusOK)
// 	}
// 	if response.Message == "" {
// 		t.Error("Message cant be blank")
// 	}
// 	if len(response.Data) != len(pkg.Playerlist)-1 {
// 		t.Error("expected", response.Data)
// 	}
// }

// func Test_user_NewPlayer(t *testing.T) {

// 	player := []byte(`{ ID:"47f1923a-0227-11ed-b939-0242ac120002",
// 	FirstName: "Ra",
// 	LastName:  "TamaG",
// 	Birthday:  time.Date(1998, time.June, 26, 12, 35, 0, 0, time.UTC),
// 	Genre:     "M",
// 	Sport:     "Cricket",
// 	TeamName:  "Caballeros de Kolkata", }`)

// 	req, err := http.NewRequest(http.MethodPost, "/newplayer", bytes.NewBuffer(player))

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	res := httptest.NewRecorder()
// 	handler := http.HandlerFunc(CreatePlayer)

// 	handler.ServeHTTP(res, req)
// 	response := pkg.Response{}
// 	json.Unmarshal(res.Body.Bytes(), &response)


// 	fmt.Println("1", response.Status)
// 	fmt.Println("2", response.Data)

// 	if response.Status != http.StatusOK {
// 		t.Error("expected", http.StatusOK)
// 	}
// 	if len(response.Data) != len(pkg.Playerlist)+1 {
// 		t.Error("expected", response.Data)
// 	}
// }
