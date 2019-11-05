package main


import "os"
import "html/template"

func main() {
	templateString := "Sample Template"
	t, _ := template.New("titleString").Parse(templateString)
	_ = t.Execute(os.Stdout, nil)
}