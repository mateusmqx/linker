package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v81/github"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

type Service struct {
	Owner    string
	RepoName string
	Info     ServiceInfo
}

type ServiceInfo struct {
	Dependencies []string `yaml:"dependencies"`
}

func main() {
	ctx := context.Background()

	githubClient := connectGitHub(ctx)

	_ = getServices(ctx, githubClient)
}

func connectGitHub(ctx context.Context) *github.Client {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatalf("Variável de ambiente GITHUB_AUTH_TOKEN não definida.")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts) // http.Client configurado para OAuth2

	client := github.NewClient(tc)
	return client
}

func getServices(ctx context.Context, githubClient *github.Client) []Service {
	repos, _, err := githubClient.Repositories.ListByAuthenticatedUser(ctx, nil)
	if err != nil {
		log.Fatalf("Erro ao listar repositórios: %v", err)
	}

	var services []Service

	for _, repo := range repos {
		owner := repo.GetOwner().GetLogin()
		repoName := repo.GetName()

		service := Service{
			Owner:    owner,
			RepoName: repoName,
		}

		path := "linker-info.yml"
		fileContent, _, _, err := githubClient.Repositories.GetContents(ctx, owner, repoName, path, nil)
		if err != nil {
			log.Printf("Erro ao obter conteúdo do linker-info.yml em %s/%s: %v", owner, repoName, err)
			services = append(services, service)
			continue
		}

		if fileContent == nil {
			log.Printf("O caminho especificado é um diretório, não um arquivo em %s/%s.", owner, repoName)
			services = append(services, service)
			continue
		}
		content, err := fileContent.GetContent()
		if err != nil {
			log.Printf("Erro ao decodificar conteúdo do arquivo em %s/%s: %v", owner, repoName, err)
			services = append(services, service)
			continue
		}

		err = yaml.Unmarshal([]byte(content), &service.Info)
		if err != nil {
			log.Printf("Erro ao fazer unmarshal do conteúdo YAML em %s/%s: %v", owner, repoName, err)
			services = append(services, service)
			continue
		}

		services = append(services, service)
	}

	return services
}

func generateMermaidDiagram(services []Service) string {
	var sb strings.Builder

	sb.WriteString("graph TD\n")

	// 1. organizar serviços por owner
	ownerMap := make(map[string][]Service)
	for _, service := range services {
		ownerMap[service.Owner] = append(ownerMap[service.Owner], service)
	}

	// 2. organizar por domínio
	for owner, svcList := range ownerMap {
		sb.WriteString("  subgraph " + owner + "\n")
		for _, svc := range svcList {
			sb.WriteString("    " + svc.RepoName + "[" + svc.RepoName + "]\n")
		}
		sb.WriteString("  end\n")
	}

	return sb.String()
}
