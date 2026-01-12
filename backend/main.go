package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v81/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		fmt.Println("Variável de ambiente GITHUB_AUTH_TOKEN não definida. Pulando exemplo autenticado.")
		return
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts) // http.Client configurado para OAuth2

	clientWithAuth := github.NewClient(tc)

	// Listando repositórios do usuário autenticado (ou de outro, dependendo das permissões do token)
	authenticatedUserRepos, _, err := clientWithAuth.Repositories.List(ctx, "", nil) // "" para o usuário autenticado
	if err != nil {
		log.Fatalf("Erro ao listar repositórios com autenticação: %v", err)
	}

	fmt.Println("--- Repositórios do usuário autenticado (Autenticado) ---")
	for _, repo := range authenticatedUserRepos {
		fmt.Printf("- %s\n", *repo.Name)
	}
}
