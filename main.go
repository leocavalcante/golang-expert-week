package main

import (
	"fmt"

	"github.com/leocavalcante/golang-expert-week/git"
)

func main() {
	fmt.Println("#VAAAAAI")
	tag := "v1.26.2"
	donoDoRepositorio := "kubernetes"
	repo := "kubernetes"

	buscador := git.BuscadorGit{}
	buscador.BuscaGitTag(tag, donoDoRepositorio, repo)
}
