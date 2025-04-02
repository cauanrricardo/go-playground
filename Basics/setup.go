package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Nome do projeto
	projectName := "Exercicio-"

	// Criando a pasta principal do projeto
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Println("Erro ao criar a pasta:", err)
		return
	}

	// Criando o arquivo README.md
	readmeContent := "# " + projectName + "\n\nEste é o projeto " + projectName + "."
	err = ioutil.WriteFile(projectName+"/README.md", []byte(readmeContent), 0644)
	if err != nil {
		fmt.Println("Erro ao criar o README.md:", err)
		return
	}

	// Criando o arquivo draft.go
	draftGoContent := `package main

import "fmt"

func main() {
    fmt.Println("Voce consegue Cauan")
}`
	err = ioutil.WriteFile(projectName+"/draft.go", []byte(draftGoContent), 0644)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo draft.go:", err)
		return
	}

	fmt.Println("Estrutura criada com sucesso!")
}
