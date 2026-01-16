package main

import (
	"fmt"
	"os"
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

	// exportar em um arquivo
	err := os.WriteFile("architecture_diagram.md", []byte(mermaid), 0644)
	if err != nil {
		fmt.Println("Erro ao escrever o arquivo:", err)
		return
	}
	fmt.Println("Diagrama gerado com sucesso em architecture_diagram.md")
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
		{
			Name:   "fraud-engine.events-api",
			System: "fraud-engine",
			Team:   "fraud-prevention",
			Domain: "risk",
			Dependencies: []string{
				"fraud-engine.reference-data",
			},
		},
		{
			Name:   "checking-account.account-service",
			System: "checking-account",
			Team:   "core-banking",
			Domain: "accounts",
			Dependencies: []string{
				"account.dict",
				"core.crypto",
			},
		},
		{
			Name:   "score.mule-account",
			System: "score",
			Team:   "risk-analytics",
			Domain: "risk",
			Dependencies: []string{
				"ml.model-inference",
				"data-lake.customer-events",
			},
		},
		{
			Name:   "payment.pix-api",
			System: "payment",
			Team:   "payments",
			Domain: "payments",
			Dependencies: []string{
				"fraud-engine.viper-api",
				"transfer-antifraud.pix-out",
				"checking-account.account-service",
			},
		},
		{
			Name:   "payment.ted-processor",
			System: "payment",
			Team:   "payments",
			Domain: "payments",
			Dependencies: []string{
				"payment.settlement-api",
				"compliance.aml-check",
			},
		},
		{
			Name:   "payment.settlement-api",
			System: "payment",
			Team:   "payments",
			Domain: "payments",
			Dependencies: []string{
				"accounting.ledger",
			},
		},
		{
			Name:   "card.core-service",
			System: "card",
			Team:   "card-platform",
			Domain: "cards",
			Dependencies: []string{
				"account.dict",
				"card.tokenizer",
			},
		},
		{
			Name:   "card.tokenizer",
			System: "card",
			Team:   "card-platform",
			Domain: "cards",
			Dependencies: []string{
				"core.crypto",
			},
		},
		{
			Name:   "card.authorization-engine",
			System: "card",
			Team:   "card-platform",
			Domain: "cards",
			Dependencies: []string{
				"fraud-engine.viper-api",
				"card.core-service",
			},
		},
		{
			Name:   "credit.product-service",
			System: "credit",
			Team:   "credit-products",
			Domain: "credit",
			Dependencies: []string{
				"account.dict",
				"credit.pricing-engine",
				"compliance.kyc",
			},
		},
		{
			Name:   "credit.pricing-engine",
			System: "credit",
			Team:   "credit-products",
			Domain: "credit",
			Dependencies: []string{
				"score.mule-account",
				"risk.rating-service",
			},
		},
		{
			Name:   "credit.disbursement-api",
			System: "credit",
			Team:   "credit-products",
			Domain: "credit",
			Dependencies: []string{
				"credit.product-service",
				"payment.settlement-api",
			},
		},
		{
			Name:   "loan.origination-api",
			System: "loan",
			Team:   "lending",
			Domain: "lending",
			Dependencies: []string{
				"credit.product-service",
				"compliance.kyc",
				"score.mule-account",
			},
		},
		{
			Name:   "loan.servicing-engine",
			System: "loan",
			Team:   "lending",
			Domain: "lending",
			Dependencies: []string{
				"loan.origination-api",
				"accounting.ledger",
			},
		},
		{
			Name:   "compliance.kyc",
			System: "compliance",
			Team:   "compliance",
			Domain: "compliance",
			Dependencies: []string{
				"account.dict",
				"external.document-validator",
			},
		},
		{
			Name:   "compliance.aml-check",
			System: "compliance",
			Team:   "compliance",
			Domain: "compliance",
			Dependencies: []string{
				"external.sanctions-list",
				"data-lake.customer-events",
			},
		},
		{
			Name:   "risk.rating-service",
			System: "risk",
			Team:   "risk-analytics",
			Domain: "risk",
			Dependencies: []string{
				"score.mule-account",
				"data-lake.customer-events",
			},
		},
		{
			Name:   "risk.exposure-monitor",
			System: "risk",
			Team:   "risk-analytics",
			Domain: "risk",
			Dependencies: []string{
				"risk.rating-service",
				"accounting.ledger",
			},
		},
		{
			Name:   "accounting.ledger",
			System: "accounting",
			Team:   "finance",
			Domain: "finance",
			Dependencies: []string{
				"core.crypto",
			},
		},
		{
			Name:   "accounting.reconciliation",
			System: "accounting",
			Team:   "finance",
			Domain: "finance",
			Dependencies: []string{
				"accounting.ledger",
				"data-lake.transactions",
			},
		},
		{
			Name:         "core.crypto",
			System:       "core",
			Team:         "platform",
			Domain:       "security",
			Dependencies: []string{},
		},
		{
			Name:   "core.auth-service",
			System: "core",
			Team:   "platform",
			Domain: "security",
			Dependencies: []string{
				"core.crypto",
			},
		},
		{
			Name:   "api-gateway.router",
			System: "api-gateway",
			Team:   "platform",
			Domain: "platform",
			Dependencies: []string{
				"core.auth-service",
				"monitoring.rate-limiter",
			},
		},
		{
			Name:   "monitoring.rate-limiter",
			System: "monitoring",
			Team:   "platform",
			Domain: "platform",
			Dependencies: []string{
				"data-store.redis-cache",
			},
		},
		{
			Name:         "monitoring.logging-service",
			System:       "monitoring",
			Team:         "platform",
			Domain:       "platform",
			Dependencies: []string{},
		},
		{
			Name:         "data-store.redis-cache",
			System:       "data-store",
			Team:         "platform",
			Domain:       "platform",
			Dependencies: []string{},
		},
		{
			Name:         "data-lake.customer-events",
			System:       "data-lake",
			Team:         "data-engineering",
			Domain:       "data",
			Dependencies: []string{},
		},
		{
			Name:         "data-lake.transactions",
			System:       "data-lake",
			Team:         "data-engineering",
			Domain:       "data",
			Dependencies: []string{},
		},
		{
			Name:   "ml.model-inference",
			System: "ml",
			Team:   "data-science",
			Domain: "analytics",
			Dependencies: []string{
				"data-lake.customer-events",
				"data-lake.transactions",
			},
		},
		{
			Name:   "notification.email-service",
			System: "notification",
			Team:   "platform",
			Domain: "communication",
			Dependencies: []string{
				"account.dict",
			},
		},
		{
			Name:   "notification.sms-service",
			System: "notification",
			Team:   "platform",
			Domain: "communication",
			Dependencies: []string{
				"account.dict",
			},
		},
		{
			Name:   "transfer.domestic-service",
			System: "transfer",
			Team:   "payments",
			Domain: "payments",
			Dependencies: []string{
				"payment.pix-api",
				"fraud-engine.viper-api",
			},
		},
		{
			Name:   "transfer.international-service",
			System: "transfer",
			Team:   "payments",
			Domain: "payments",
			Dependencies: []string{
				"payment.ted-processor",
				"compliance.aml-check",
			},
		},
		{
			Name:   "customer.profile-api",
			System: "customer",
			Team:   "core-banking",
			Domain: "customers",
			Dependencies: []string{
				"account.dict",
				"notification.email-service",
			},
		},
		{
			Name:   "customer.preferences-service",
			System: "customer",
			Team:   "core-banking",
			Domain: "customers",
			Dependencies: []string{
				"customer.profile-api",
				"data-store.redis-cache",
			},
		},
		{
			Name:   "savings.account-service",
			System: "savings",
			Team:   "products",
			Domain: "products",
			Dependencies: []string{
				"account.dict",
				"interest.calculator",
			},
		},
		{
			Name:   "savings.interest-calculator",
			System: "interest",
			Team:   "products",
			Domain: "products",
			Dependencies: []string{
				"accounting.ledger",
			},
		},
		{
			Name:   "investment.portfolio-api",
			System: "investment",
			Team:   "wealth",
			Domain: "investment",
			Dependencies: []string{
				"account.dict",
				"external.market-data",
			},
		},
		{
			Name:   "insurance.policy-engine",
			System: "insurance",
			Team:   "products",
			Domain: "products",
			Dependencies: []string{
				"account.dict",
				"risk.rating-service",
			},
		},
		{
			Name:   "reporting.statement-generator",
			System: "reporting",
			Team:   "finance",
			Domain: "reporting",
			Dependencies: []string{
				"accounting.ledger",
				"customer.profile-api",
			},
		},
		{
			Name:   "reporting.analytics-dashboard",
			System: "reporting",
			Team:   "data-engineering",
			Domain: "reporting",
			Dependencies: []string{
				"data-lake.transactions",
				"data-lake.customer-events",
			},
		},
		{
			Name:   "webhook.event-dispatcher",
			System: "webhook",
			Team:   "platform",
			Domain: "integration",
			Dependencies: []string{
				"monitoring.logging-service",
				"data-store.redis-cache",
			},
		},
		{
			Name:   "integration.partner-api",
			System: "integration",
			Team:   "platform",
			Domain: "integration",
			Dependencies: []string{
				"core.auth-service",
				"webhook.event-dispatcher",
			},
		},
		{
			Name:         "security.secrets-manager",
			System:       "security",
			Team:         "platform",
			Domain:       "security",
			Dependencies: []string{},
		},
		{
			Name:   "security.audit-logger",
			System: "security",
			Team:   "compliance",
			Domain: "compliance",
			Dependencies: []string{
				"monitoring.logging-service",
			},
		},
		{
			Name:   "cache.invalidation-service",
			System: "cache",
			Team:   "platform",
			Domain: "platform",
			Dependencies: []string{
				"data-store.redis-cache",
			},
		},
		{
			Name:   "batch.scheduled-jobs",
			System: "batch",
			Team:   "data-engineering",
			Domain: "data",
			Dependencies: []string{
				"accounting.reconciliation",
				"data-lake.transactions",
			},
		},
		{
			Name:         "queue.message-broker",
			System:       "queue",
			Team:         "platform",
			Domain:       "platform",
			Dependencies: []string{},
		},
		{
			Name:   "async.event-processor",
			System: "async",
			Team:   "platform",
			Domain: "platform",
			Dependencies: []string{
				"queue.message-broker",
				"monitoring.logging-service",
			},
		},
		{
			Name:   "merchant.onboarding-api",
			System: "merchant",
			Team:   "partnerships",
			Domain: "merchants",
			Dependencies: []string{
				"compliance.kyc",
				"account.dict",
			},
		},
		{
			Name:   "merchant.settlement-service",
			System: "merchant",
			Team:   "payments",
			Domain: "merchants",
			Dependencies: []string{
				"payment.settlement-api",
				"merchant.onboarding-api",
			},
		},
		{
			Name:   "affiliate.commission-calculator",
			System: "affiliate",
			Team:   "partnerships",
			Domain: "merchants",
			Dependencies: []string{
				"accounting.ledger",
				"data-lake.transactions",
			},
		},
		{
			Name:   "subscription.billing-engine",
			System: "subscription",
			Team:   "products",
			Domain: "products",
			Dependencies: []string{
				"account.dict",
				"payment.pix-api",
			},
		},
		{
			Name:   "promotion.offers-api",
			System: "promotion",
			Team:   "marketing",
			Domain: "marketing",
			Dependencies: []string{
				"customer.preferences-service",
				"subscription.billing-engine",
			},
		},
		{
			Name:   "analytics.customer-journey",
			System: "analytics",
			Team:   "data-science",
			Domain: "analytics",
			Dependencies: []string{
				"data-lake.customer-events",
				"ml.model-inference",
			},
		},
		{
			Name:   "search.customer-search",
			System: "search",
			Team:   "platform",
			Domain: "platform",
			Dependencies: []string{
				"customer.profile-api",
				"data-store.redis-cache",
			},
		},
		{
			Name:   "mobile.push-notification",
			System: "mobile",
			Team:   "platform",
			Domain: "communication",
			Dependencies: []string{
				"notification.sms-service",
				"customer.profile-api",
			},
		},
		{
			Name:   "document.generation-service",
			System: "document",
			Team:   "platform",
			Domain: "document",
			Dependencies: []string{
				"customer.profile-api",
				"accounting.ledger",
			},
		},
		{
			Name:   "document.archive-service",
			System: "document",
			Team:   "platform",
			Domain: "document",
			Dependencies: []string{
				"document.generation-service",
			},
		},
		{
			Name:   "dispute.management-api",
			System: "dispute",
			Team:   "customer-service",
			Domain: "customer-service",
			Dependencies: []string{
				"card.core-service",
				"payment.pix-api",
				"security.audit-logger",
			},
		},
		{
			Name:   "chargeback.processor",
			System: "chargeback",
			Team:   "customer-service",
			Domain: "customer-service",
			Dependencies: []string{
				"dispute.management-api",
				"payment.settlement-api",
			},
		},
		{
			Name:   "reversal.transaction-service",
			System: "reversal",
			Team:   "payments",
			Domain: "payments",
			Dependencies: []string{
				"payment.settlement-api",
				"accounting.ledger",
			},
		},
		{
			Name:   "refund.processor",
			System: "refund",
			Team:   "payments",
			Domain: "payments",
			Dependencies: []string{
				"reversal.transaction-service",
				"notification.email-service",
			},
		},
		{
			Name:   "limit.management-service",
			System: "limit",
			Team:   "risk-analytics",
			Domain: "risk",
			Dependencies: []string{
				"customer.profile-api",
				"score.mule-account",
				"accounting.ledger",
			},
		},
		{
			Name:   "velocity.check-engine",
			System: "velocity",
			Team:   "fraud-prevention",
			Domain: "risk",
			Dependencies: []string{
				"data-store.redis-cache",
				"fraud-engine.viper-api",
			},
		},
		{
			Name:   "behavior.anomaly-detector",
			System: "behavior",
			Team:   "data-science",
			Domain: "analytics",
			Dependencies: []string{
				"ml.model-inference",
				"data-lake.customer-events",
				"fraud-engine.viper-api",
			},
		},
		{
			Name:   "network.graph-analysis",
			System: "network",
			Team:   "risk-analytics",
			Domain: "risk",
			Dependencies: []string{
				"data-lake.transactions",
				"ml.model-inference",
			},
		},
		{
			Name:         "external.document-validator",
			System:       "external",
			Team:         "compliance",
			Domain:       "compliance",
			Dependencies: []string{},
		},
		{
			Name:         "external.sanctions-list",
			System:       "external",
			Team:         "compliance",
			Domain:       "compliance",
			Dependencies: []string{},
		},
		{
			Name:         "external.market-data",
			System:       "external",
			Team:         "wealth",
			Domain:       "investment",
			Dependencies: []string{},
		},
		{
			Name:   "feedback.survey-api",
			System: "feedback",
			Team:   "customer-service",
			Domain: "customer-service",
			Dependencies: []string{
				"customer.profile-api",
				"data-lake.customer-events",
			},
		},
		{
			Name:   "rating.service-quality",
			System: "rating",
			Team:   "customer-service",
			Domain: "customer-service",
			Dependencies: []string{
				"feedback.survey-api",
				"customer.profile-api",
			},
		},
		{
			Name:   "chatbot.conversation-engine",
			System: "chatbot",
			Team:   "customer-service",
			Domain: "customer-service",
			Dependencies: []string{
				"customer.profile-api",
				"ml.model-inference",
			},
		},
		{
			Name:   "help-desk.ticket-management",
			System: "help-desk",
			Team:   "customer-service",
			Domain: "customer-service",
			Dependencies: []string{
				"customer.profile-api",
				"security.audit-logger",
			},
		},
		{
			Name:   "knowledge-base.article-service",
			System: "knowledge-base",
			Team:   "customer-service",
			Domain: "customer-service",
			Dependencies: []string{
				"search.customer-search",
			},
		},
		{
			Name:         "translation.i18n-service",
			System:       "translation",
			Team:         "platform",
			Domain:       "platform",
			Dependencies: []string{},
		},
		{
			Name:   "localization.currency-converter",
			System: "localization",
			Team:   "payments",
			Domain: "payments",
			Dependencies: []string{
				"external.market-data",
				"data-store.redis-cache",
			},
		},
		{
			Name:   "tax.calculation-engine",
			System: "tax",
			Team:   "finance",
			Domain: "finance",
			Dependencies: []string{
				"accounting.ledger",
				"customer.profile-api",
			},
		},
		{
			Name:   "invoice.generation-api",
			System: "invoice",
			Team:   "finance",
			Domain: "finance",
			Dependencies: []string{
				"document.generation-service",
				"accounting.ledger",
			},
		},
		{
			Name:   "expense.management-api",
			System: "expense",
			Team:   "finance",
			Domain: "finance",
			Dependencies: []string{
				"accounting.ledger",
				"employee.registry",
			},
		},
		{
			Name:         "employee.registry",
			System:       "employee",
			Team:         "hr",
			Domain:       "hr",
			Dependencies: []string{},
		},
		{
			Name:   "payroll.processing-engine",
			System: "payroll",
			Team:   "finance",
			Domain: "finance",
			Dependencies: []string{
				"employee.registry",
				"accounting.ledger",
			},
		},
		{
			Name:   "attendance.tracking-system",
			System: "attendance",
			Team:   "hr",
			Domain: "hr",
			Dependencies: []string{
				"employee.registry",
			},
		},
		{
			Name:   "leave.management-api",
			System: "leave",
			Team:   "hr",
			Domain: "hr",
			Dependencies: []string{
				"employee.registry",
				"notification.email-service",
			},
		},
		{
			Name:   "performance.evaluation-system",
			System: "performance",
			Team:   "hr",
			Domain: "hr",
			Dependencies: []string{
				"employee.registry",
				"feedback.survey-api",
			},
		},
		{
			Name:   "training.learning-platform",
			System: "training",
			Team:   "hr",
			Domain: "hr",
			Dependencies: []string{
				"employee.registry",
				"document.generation-service",
			},
		},
		{
			Name:   "recruitment.job-portal",
			System: "recruitment",
			Team:   "hr",
			Domain: "hr",
			Dependencies: []string{
				"employee.registry",
			},
		},
		{
			Name:   "asset.management-system",
			System: "asset",
			Team:   "operations",
			Domain: "operations",
			Dependencies: []string{
				"employee.registry",
				"accounting.ledger",
			},
		},
		{
			Name:   "facility.booking-api",
			System: "facility",
			Team:   "operations",
			Domain: "operations",
			Dependencies: []string{
				"employee.registry",
				"notification.email-service",
			},
		},
		{
			Name:   "office.expense-tracking",
			System: "office",
			Team:   "operations",
			Domain: "operations",
			Dependencies: []string{
				"expense.management-api",
				"employee.registry",
			},
		},
		{
			Name:   "vendor.management-api",
			System: "vendor",
			Team:   "procurement",
			Domain: "procurement",
			Dependencies: []string{
				"compliance.kyc",
				"accounting.ledger",
			},
		},
		{
			Name:   "procurement.order-system",
			System: "procurement",
			Team:   "procurement",
			Domain: "procurement",
			Dependencies: []string{
				"vendor.management-api",
				"inventory.stock-management",
			},
		},
		{
			Name:   "inventory.stock-management",
			System: "inventory",
			Team:   "operations",
			Domain: "operations",
			Dependencies: []string{
				"accounting.ledger",
			},
		},
		{
			Name:   "logistics.shipment-tracking",
			System: "logistics",
			Team:   "operations",
			Domain: "operations",
			Dependencies: []string{
				"vendor.management-api",
				"notification.email-service",
			},
		},
		{
			Name:   "warranty.claim-processor",
			System: "warranty",
			Team:   "customer-service",
			Domain: "customer-service",
			Dependencies: []string{
				"customer.profile-api",
				"inventory.stock-management",
			},
		},
		{
			Name:   "migration.data-sync",
			System: "migration",
			Team:   "data-engineering",
			Domain: "data",
			Dependencies: []string{
				"data-lake.customer-events",
				"data-lake.transactions",
			},
		},
		{
			Name:         "backup.disaster-recovery",
			System:       "backup",
			Team:         "platform",
			Domain:       "platform",
			Dependencies: []string{},
		},
		{
			Name:   "test.environment-provisioner",
			System: "test",
			Team:   "platform",
			Domain: "platform",
			Dependencies: []string{
				"core.auth-service",
			},
		},
	}
}
