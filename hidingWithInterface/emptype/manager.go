package emptype

import "fmt"

type Manager struct {
	Name     string
	Age      string
	TeamSize string
}

func (m Manager) GetTeamSize() {
	fmt.Println("Manager's Team Size: ", m.TeamSize)
}
