package emptype

import "fmt"

type Lead struct {
	Name     string
	Age      string
	TeamSize string
}

func (m Lead) GetTeamSize() {
	fmt.Println("Lead's Team Size: ", m.TeamSize)
}
