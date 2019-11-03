package privatepackage

type OtherEmployeeDetails struct {
	name        string
	Age         string
	Designation string
}

func CreateEmployee(name string, age string, designation string) *OtherEmployeeDetails {
	return &OtherEmployeeDetails{name, age, designation}
}

func (emp OtherEmployeeDetails) UserName() string {
	return emp.name
}
