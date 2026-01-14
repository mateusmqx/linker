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
	authenticatedUserRepos, _, err := clientWithAuth.Repositories.ListByAuthenticatedUser(ctx, nil)
	if err != nil {
		log.Fatalf("Erro ao listar repositórios com autenticação: %v", err)
	}

	fmt.Println("--- Repositórios do usuário autenticado (Autenticado) ---")
	for _, repo := range authenticatedUserRepos {
		fmt.Printf("- %s\n", *repo.Name)
		owner := repo.GetOwner().GetLogin()
		repo := repo.GetName()
		path := "linker-info.yml"
		fileContent, directoryContent, _, err := clientWithAuth.Repositories.GetContents(ctx, owner, repo, path, nil)
		if err != nil {
			log.Printf("Erro ao obter conteúdo do linker-info.yml: %v", err)
			continue
		}

		// 2. Verificação do tipo de conteúdo retornado
		if directoryContent != nil {
			log.Printf("Esperava um arquivo, mas o caminho especificado é um diretório contendo %d itens.", len(directoryContent))
		}

		// Verifica se é realmente um arquivo e não um diretório
		if fileContent == nil {
			log.Printf("O caminho especificado é um diretório, não um arquivo.")
		}

		// 3. Decodificação do conteúdo
		// O conteúdo vem em Base64, mas a lib tem este helper que decodifica para string
		content, err := fileContent.GetContent()
		if err != nil {
			log.Printf("Erro ao decodificar conteúdo: %v", err)
		}

		fmt.Printf("Conteúdo do arquivo:\n%s\n", content)
	}
}
