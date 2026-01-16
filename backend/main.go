package main

import (
	"fmt"
	"strings"
)

type App struct {
	Name             string
	System           string
	Team             string
	Domain           string
	Dependencies     []string
	RepositoryUrl    string
	DocumentationUrl string
}

func main() {
	apps := generateApps()
	mermaid := generateMermaid(apps)
	fmt.Println(mermaid)
}

func generateMermaid(apps []App) string {
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

	// agrupamento por times
	teamMap := make(map[string][]App)
	for _, app := range apps {
		teamMap[app.Team] = append(teamMap[app.Team], app)
	}

	// map colors to teams
	teamColors := map[string]string{
		"fraud-prevention": "blue",
		"core-banking":     "purple",
	}

	// 3. Gerar os Nós (Nodes) com HTML
	for team, appList := range teamMap {
		// Nome do subgrafo (Time)
		sb.WriteString(fmt.Sprintf("    subgraph %s [Time: %s]\n", team, strings.ToUpper(team)))
		sb.WriteString("    direction TB\n")

		for _, app := range appList {
			// Label com HTML
			label := fmt.Sprintf("\"<b>%s</b><br/><span style='font-size:12px'>Sys: %s</span><br/><br/><a href='%s' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='%s' style='color:#60a5fa;text-decoration:none'>Confluence</a>\"",
				app.Name, app.System, app.RepositoryUrl, app.DocumentationUrl)

			sb.WriteString(fmt.Sprintf("        %s[%s]\n", app.Name, label))
		}
		sb.WriteString("    end\n") // Fim do subgraph

		// Aplicar estilo aos nós deste domínio
		style := teamColors[team]
		if style == "" {
			style = "blue"
		}

		// Listar nomes para aplicar classe em lote
		var names []string
		for _, s := range appList {
			names = append(names, s.Name)
		}
		sb.WriteString(fmt.Sprintf("    class %s %s;\n\n", strings.Join(names, ","), style))
	}

	// 4. Gerar as Dependências (Edges)
	// Vamos identificar quais dependências são "externas" (não estão no JSON principal)
	knownApps := make(map[string]bool)
	for _, s := range apps {
		knownApps[s.Name] = true
	}

	sb.WriteString("    %% Dependências\n")
	for _, s := range apps {
		for _, dep := range s.Dependencies {
			sb.WriteString(fmt.Sprintf("    %s --> %s\n", s.Name, dep))

			// Se a dependência for externa, aplicamos um estilo diferente
			if !knownApps[dep] {
				sb.WriteString(fmt.Sprintf("    class %s external;\n", dep))
			}
		}
	}

	return sb.String()
}

func generateApps() []App {
	return []App{
		{
			Name:   "fraud-engine.viper-api",
			System: "fraud-engine",
			Team:   "fraud-prevention",
			Domain: "risk",
			Dependencies: []string{
				"fraud-engine.reference-data",
				"fraud-solution.risk-location",
			},
		},
		{
			Name:         "fraud-engine.reference-data",
			System:       "fraud-engine",
			Team:         "fraud-prevention",
			Domain:       "risk",
			Dependencies: []string{},
		},
		{
			Name:         "fraud-solution.risk-location",
			System:       "fraud-solution",
			Team:         "fraud-prevention",
			Domain:       "risk",
			Dependencies: []string{},
		},
		{
			Name:   "fraud-engine.viper-bff",
			System: "fraud-engine",
			Team:   "fraud-prevention",
			Domain: "risk",
			Dependencies: []string{
				"fraud-engine.viper-api",
				"fraud-engine.events-api",
			},
		},
		{
			Name:   "transfer-antifraud.pix-out",
			System: "transfer-antifraud",
			Team:   "fraud-prevention",
			Domain: "risk",
			Dependencies: []string{
				"checking-account.account-service",
				"account.dict",
				"score.mule-account",
				"fraud-engine.viper-api",
			},
		},
		{
			Name:         "account.dict",
			System:       "account",
			Team:         "core-banking",
			Domain:       "accounts",
			Dependencies: []string{},
		},
	}
}
