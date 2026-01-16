package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Service struct {
	Name         string   `json:"name"`
	Domain       string   `json:"domain"`
	Owner        string   `json:"owner"`
	Dependencies []string `json:"dependencies"`
}

func main() {

	jsonData := `[
		{
			"name": "temp_repo_linker", 
			"domain": "risk", 
			"owner": "mateus",
			"dependencies": ["data_processor", "notification_service"]
		},
		{
			"name": "analytics_dashboard", 
			"domain": "insights", 
			"owner": "sophia",
			"dependencies": ["data_collector", "report_generator"]
		},
		{
			"name": "data_processor", 
			"domain": "accounts", 
			"owner": "liam",
			"dependencies": ["auth_service", "profile_service"]
		},
		{
			"name": "auth_service", 
			"domain": "finance", 
			"owner": "olivia",
			"dependencies": []
		}
	]`

	var services []Service
	if err := json.Unmarshal([]byte(jsonData), &services); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(GenerateRichMermaid(services))
}

func GenerateRichMermaid(services []Service) string {
	var sb strings.Builder

	// 1. Configuração do Tema (Dark Mode)
	sb.WriteString("%%{init: {'theme': 'dark', 'themeVariables': { 'darkMode': true }}}%%\n")
	sb.WriteString("flowchart TD\n")

	// 2. Definição de Estilos (CSS)
	// Blue style
	sb.WriteString("    classDef blue stroke:#3b82f6,stroke-width:2px,fill:#1e293b,color:white,rx:10,ry:10;\n")
	// Purple style
	sb.WriteString("    classDef purple stroke:#a855f7,stroke-width:2px,fill:#1e293b,color:white,rx:10,ry:10;\n")
	// Pink/Red style
	sb.WriteString("    classDef pink stroke:#ec4899,stroke-width:2px,fill:#1e293b,color:white,rx:10,ry:10;\n")
	// Default style for external deps
	sb.WriteString("    classDef external stroke:#64748b,stroke-width:1px,stroke-dasharray: 5 5,fill:#0f172a,color:#cbd5e1,rx:5,ry:5;\n\n")

	// Agrupamento por Domínio
	domainMap := make(map[string][]Service)
	for _, s := range services {
		domainMap[s.Domain] = append(domainMap[s.Domain], s)
	}

	// Mapa para controlar cores baseadas no domínio
	domainColors := map[string]string{
		"risk":     "pink",
		"insights": "blue",
		"accounts": "purple",
		"finance":  "blue",
	}

	// 3. Gerar os Nós (Nodes) com HTML
	for domain, svcList := range domainMap {
		// Nome do subgrafo (Time)
		sb.WriteString(fmt.Sprintf("    subgraph %s [Time: %s]\n", domain, strings.ToUpper(domain)))
		sb.WriteString("    direction TB\n")

		for _, s := range svcList {
			// Simulação de URLs (podes vir do JSON no futuro)
			githubUrl := fmt.Sprintf("https://github.com/org/%s", s.Name)
			confUrl := "#"

			// Label HTML complexa
			// Usamos <br/> para quebra de linha e tags <a> para links
			label := fmt.Sprintf("\"<b>%s</b><br/><span style='font-size:12px'>Sys: %s</span><br/><br/><a href='%s' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='%s' style='color:#60a5fa;text-decoration:none'>Confluence</a>\"",
				s.Name, s.Owner, githubUrl, confUrl)

			sb.WriteString(fmt.Sprintf("        %s[%s]\n", s.Name, label))
		}
		sb.WriteString("    end\n") // Fim do subgraph

		// Aplicar estilo aos nós deste domínio
		style := domainColors[domain]
		if style == "" {
			style = "blue"
		}

		// Listar nomes para aplicar classe em lote
		var names []string
		for _, s := range svcList {
			names = append(names, s.Name)
		}
		sb.WriteString(fmt.Sprintf("    class %s %s;\n\n", strings.Join(names, ","), style))
	}

	// 4. Gerar Dependências (Edges)
	// Vamos identificar quais dependências são "externas" (não estão no JSON principal)
	knownServices := make(map[string]bool)
	for _, s := range services {
		knownServices[s.Name] = true
	}

	sb.WriteString("    %% Dependências\n")
	for _, s := range services {
		for _, dep := range s.Dependencies {
			sb.WriteString(fmt.Sprintf("    %s --> %s\n", s.Name, dep))

			// Se a dependência for externa, aplicamos um estilo diferente
			if !knownServices[dep] {
				sb.WriteString(fmt.Sprintf("    class %s external;\n", dep))
			}
		}
	}

	return sb.String()
}
