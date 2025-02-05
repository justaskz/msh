package pack

// import (
// 	"html/template"
// 	"log"
// 	"os/exec"
// )

// func Install(url string, location string) error {
// 	tmpl := `
// 	{{.Name}}
// 	`
// 	foo := template.Must(template.New("person").Parse(tmpl))

// 	output, err := exec.Command("bash", "-c", foo).Output()

// 	log.Println(string(output))

// 	if err != nil {
// 		log.Println("could not run command: ", err)
// 	}

// 	return err
// }
