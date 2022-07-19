package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Member struct {

	Id    int     `json:"ID"`
	Name  string  `json:"Name"`
	Sname string  `json:"Second name"`
	Note  string  `json:"Note"`
	Hora  float32 `json:"The meeting time is"`

}

func member() {

	m := Member{}
  
	
	fmt.Println("Enter the list number of this member (starting at 0) ")
	fmt.Scan(&m.Id)
	
	fmt.Println("Enter the name of the member ")
	fmt.Scan(&m.Name)
	
	fmt.Println("Enter", m.Name, "last name ")
	fmt.Scan(&m.Sname)
	
	fmt.Println("Enter a note for this meeting ")
	fmt.Scan(&m.Note)
	
	fmt.Println("Enter meeting time with ", m.Name, m.Sname)
	fmt.Scan(&m.Hora)
	

	if m.Hora > 24 {

		fmt.Println("Enter a hour valid with the format of 24 hours ")

		return

	}
	
		
		file_data, err := ioutil.ReadFile("./members.json")
		if err != nil {
			fmt.Println("error")
		}
		
		var members []Member
		
		errorrr := json.Unmarshal(file_data, &members)
		if errorrr != nil {
			fmt.Println("error")
		}
		
		members = append(members, m)
		
		dname, err := json.MarshalIndent(members, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Print(string(dname))
		
		error := ioutil.WriteFile("members.json", []byte(dname), 0644)
		if error != nil {
			log.Fatal(err)
		}
	}
	
func delete() {
	var eliminar int

	file_data, err := ioutil.ReadFile("./members.json")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Print(string(file_data))

	

	fmt.Println("Which member do you want to remove?")
	fmt.Scan(&eliminar)

	

	var members []Member

	errorrr := json.Unmarshal(file_data, &members)
	if errorrr != nil {
		fmt.Println("error")
	}

	
	members = append(members[:eliminar], members[eliminar+1:]...)
	
	dname, err := json.MarshalIndent(members, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(dname))
	
	error := ioutil.WriteFile("members.json", []byte(dname), 0644)
	if error != nil {
		log.Fatal(err)
	}

}

func main() {

	var a int
	for a != 4 {
		fmt.Println("What would you like to do?")

		fmt.Println("1. Create member")

		fmt.Println("2. Delete member")
 
		fmt.Println("3. List Members")

		fmt.Println("4. exit")
		fmt.Scan(&a)

		switch a {

		case 1:
			
			member()
			
		case 2:
			
			delete()
			
		case 3:
			
			file_data, err := ioutil.ReadFile("./members.json")

			if err != nil {
				fmt.Println("error")
			}
			fmt.Print(string(file_data))

		case 4:
			
			fmt.Println("Exit")
			break
			
		}
	}
}