%%{init: {'theme': 'dark', 'themeVariables': { 'darkMode': true }}}%%
flowchart TD
    classDef blue stroke:#3b82f6,stroke-width:2px,fill:#1e293b,color:white,rx:10,ry:10;
    classDef purple stroke:#a855f7,stroke-width:2px,fill:#1e293b,color:white,rx:10,ry:10;
    classDef pink stroke:#ec4899,stroke-width:2px,fill:#1e293b,color:white,rx:10,ry:10;
    classDef external stroke:#64748b,stroke-width:1px,stroke-dasharray: 5 5,fill:#0f172a,color:#cbd5e1,rx:5,ry:5;

    subgraph data-science [Time: DATA-SCIENCE]
    direction TB
        ml.model-inference["<b>ml.model-inference</b><br/><span style='font-size:12px'>Sys: ml</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        analytics.customer-journey["<b>analytics.customer-journey</b><br/><span style='font-size:12px'>Sys: analytics</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        behavior.anomaly-detector["<b>behavior.anomaly-detector</b><br/><span style='font-size:12px'>Sys: behavior</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class ml.model-inference,analytics.customer-journey,behavior.anomaly-detector blue;

    subgraph customer-service [Time: CUSTOMER-SERVICE]
    direction TB
        dispute.management-api["<b>dispute.management-api</b><br/><span style='font-size:12px'>Sys: dispute</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        chargeback.processor["<b>chargeback.processor</b><br/><span style='font-size:12px'>Sys: chargeback</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        feedback.survey-api["<b>feedback.survey-api</b><br/><span style='font-size:12px'>Sys: feedback</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        rating.service-quality["<b>rating.service-quality</b><br/><span style='font-size:12px'>Sys: rating</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        chatbot.conversation-engine["<b>chatbot.conversation-engine</b><br/><span style='font-size:12px'>Sys: chatbot</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        help-desk.ticket-management["<b>help-desk.ticket-management</b><br/><span style='font-size:12px'>Sys: help-desk</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        knowledge-base.article-service["<b>knowledge-base.article-service</b><br/><span style='font-size:12px'>Sys: knowledge-base</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        warranty.claim-processor["<b>warranty.claim-processor</b><br/><span style='font-size:12px'>Sys: warranty</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class dispute.management-api,chargeback.processor,feedback.survey-api,rating.service-quality,chatbot.conversation-engine,help-desk.ticket-management,knowledge-base.article-service,warranty.claim-processor blue;

    subgraph hr [Time: HR]
    direction TB
        employee.registry["<b>employee.registry</b><br/><span style='font-size:12px'>Sys: employee</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        attendance.tracking-system["<b>attendance.tracking-system</b><br/><span style='font-size:12px'>Sys: attendance</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        leave.management-api["<b>leave.management-api</b><br/><span style='font-size:12px'>Sys: leave</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        performance.evaluation-system["<b>performance.evaluation-system</b><br/><span style='font-size:12px'>Sys: performance</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        training.learning-platform["<b>training.learning-platform</b><br/><span style='font-size:12px'>Sys: training</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        recruitment.job-portal["<b>recruitment.job-portal</b><br/><span style='font-size:12px'>Sys: recruitment</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class employee.registry,attendance.tracking-system,leave.management-api,performance.evaluation-system,training.learning-platform,recruitment.job-portal blue;

    subgraph procurement [Time: PROCUREMENT]
    direction TB
        vendor.management-api["<b>vendor.management-api</b><br/><span style='font-size:12px'>Sys: vendor</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        procurement.order-system["<b>procurement.order-system</b><br/><span style='font-size:12px'>Sys: procurement</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class vendor.management-api,procurement.order-system blue;

    subgraph fraud-prevention [Time: FRAUD-PREVENTION]
    direction TB
        fraud-engine.viper-api["<b>fraud-engine.viper-api</b><br/><span style='font-size:12px'>Sys: fraud-engine</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        fraud-engine.reference-data["<b>fraud-engine.reference-data</b><br/><span style='font-size:12px'>Sys: fraud-engine</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        fraud-solution.risk-location["<b>fraud-solution.risk-location</b><br/><span style='font-size:12px'>Sys: fraud-solution</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        fraud-engine.viper-bff["<b>fraud-engine.viper-bff</b><br/><span style='font-size:12px'>Sys: fraud-engine</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        transfer-antifraud.pix-out["<b>transfer-antifraud.pix-out</b><br/><span style='font-size:12px'>Sys: transfer-antifraud</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        fraud-engine.events-api["<b>fraud-engine.events-api</b><br/><span style='font-size:12px'>Sys: fraud-engine</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        velocity.check-engine["<b>velocity.check-engine</b><br/><span style='font-size:12px'>Sys: velocity</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class fraud-engine.viper-api,fraud-engine.reference-data,fraud-solution.risk-location,fraud-engine.viper-bff,transfer-antifraud.pix-out,fraud-engine.events-api,velocity.check-engine blue;

    subgraph compliance [Time: COMPLIANCE]
    direction TB
        compliance.kyc["<b>compliance.kyc</b><br/><span style='font-size:12px'>Sys: compliance</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        compliance.aml-check["<b>compliance.aml-check</b><br/><span style='font-size:12px'>Sys: compliance</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        security.audit-logger["<b>security.audit-logger</b><br/><span style='font-size:12px'>Sys: security</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        external.document-validator["<b>external.document-validator</b><br/><span style='font-size:12px'>Sys: external</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        external.sanctions-list["<b>external.sanctions-list</b><br/><span style='font-size:12px'>Sys: external</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class compliance.kyc,compliance.aml-check,security.audit-logger,external.document-validator,external.sanctions-list blue;

    subgraph data-engineering [Time: DATA-ENGINEERING]
    direction TB
        data-lake.customer-events["<b>data-lake.customer-events</b><br/><span style='font-size:12px'>Sys: data-lake</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        data-lake.transactions["<b>data-lake.transactions</b><br/><span style='font-size:12px'>Sys: data-lake</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        reporting.analytics-dashboard["<b>reporting.analytics-dashboard</b><br/><span style='font-size:12px'>Sys: reporting</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        batch.scheduled-jobs["<b>batch.scheduled-jobs</b><br/><span style='font-size:12px'>Sys: batch</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        migration.data-sync["<b>migration.data-sync</b><br/><span style='font-size:12px'>Sys: migration</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class data-lake.customer-events,data-lake.transactions,reporting.analytics-dashboard,batch.scheduled-jobs,migration.data-sync blue;

    subgraph wealth [Time: WEALTH]
    direction TB
        investment.portfolio-api["<b>investment.portfolio-api</b><br/><span style='font-size:12px'>Sys: investment</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        external.market-data["<b>external.market-data</b><br/><span style='font-size:12px'>Sys: external</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class investment.portfolio-api,external.market-data blue;

    subgraph operations [Time: OPERATIONS]
    direction TB
        asset.management-system["<b>asset.management-system</b><br/><span style='font-size:12px'>Sys: asset</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        facility.booking-api["<b>facility.booking-api</b><br/><span style='font-size:12px'>Sys: facility</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        office.expense-tracking["<b>office.expense-tracking</b><br/><span style='font-size:12px'>Sys: office</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        inventory.stock-management["<b>inventory.stock-management</b><br/><span style='font-size:12px'>Sys: inventory</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        logistics.shipment-tracking["<b>logistics.shipment-tracking</b><br/><span style='font-size:12px'>Sys: logistics</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class asset.management-system,facility.booking-api,office.expense-tracking,inventory.stock-management,logistics.shipment-tracking blue;

    subgraph credit-products [Time: CREDIT-PRODUCTS]
    direction TB
        credit.product-service["<b>credit.product-service</b><br/><span style='font-size:12px'>Sys: credit</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        credit.pricing-engine["<b>credit.pricing-engine</b><br/><span style='font-size:12px'>Sys: credit</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        credit.disbursement-api["<b>credit.disbursement-api</b><br/><span style='font-size:12px'>Sys: credit</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class credit.product-service,credit.pricing-engine,credit.disbursement-api blue;

    subgraph lending [Time: LENDING]
    direction TB
        loan.origination-api["<b>loan.origination-api</b><br/><span style='font-size:12px'>Sys: loan</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        loan.servicing-engine["<b>loan.servicing-engine</b><br/><span style='font-size:12px'>Sys: loan</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class loan.origination-api,loan.servicing-engine blue;

    subgraph finance [Time: FINANCE]
    direction TB
        accounting.ledger["<b>accounting.ledger</b><br/><span style='font-size:12px'>Sys: accounting</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        accounting.reconciliation["<b>accounting.reconciliation</b><br/><span style='font-size:12px'>Sys: accounting</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        reporting.statement-generator["<b>reporting.statement-generator</b><br/><span style='font-size:12px'>Sys: reporting</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        tax.calculation-engine["<b>tax.calculation-engine</b><br/><span style='font-size:12px'>Sys: tax</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        invoice.generation-api["<b>invoice.generation-api</b><br/><span style='font-size:12px'>Sys: invoice</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        expense.management-api["<b>expense.management-api</b><br/><span style='font-size:12px'>Sys: expense</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        payroll.processing-engine["<b>payroll.processing-engine</b><br/><span style='font-size:12px'>Sys: payroll</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class accounting.ledger,accounting.reconciliation,reporting.statement-generator,tax.calculation-engine,invoice.generation-api,expense.management-api,payroll.processing-engine blue;

    subgraph partnerships [Time: PARTNERSHIPS]
    direction TB
        merchant.onboarding-api["<b>merchant.onboarding-api</b><br/><span style='font-size:12px'>Sys: merchant</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        affiliate.commission-calculator["<b>affiliate.commission-calculator</b><br/><span style='font-size:12px'>Sys: affiliate</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class merchant.onboarding-api,affiliate.commission-calculator blue;

    subgraph marketing [Time: MARKETING]
    direction TB
        promotion.offers-api["<b>promotion.offers-api</b><br/><span style='font-size:12px'>Sys: promotion</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class promotion.offers-api blue;

    subgraph core-banking [Time: CORE-BANKING]
    direction TB
        account.dict["<b>account.dict</b><br/><span style='font-size:12px'>Sys: account</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        checking-account.account-service["<b>checking-account.account-service</b><br/><span style='font-size:12px'>Sys: checking-account</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        customer.profile-api["<b>customer.profile-api</b><br/><span style='font-size:12px'>Sys: customer</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        customer.preferences-service["<b>customer.preferences-service</b><br/><span style='font-size:12px'>Sys: customer</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class account.dict,checking-account.account-service,customer.profile-api,customer.preferences-service purple;

    subgraph platform [Time: PLATFORM]
    direction TB
        core.crypto["<b>core.crypto</b><br/><span style='font-size:12px'>Sys: core</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        core.auth-service["<b>core.auth-service</b><br/><span style='font-size:12px'>Sys: core</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        api-gateway.router["<b>api-gateway.router</b><br/><span style='font-size:12px'>Sys: api-gateway</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        monitoring.rate-limiter["<b>monitoring.rate-limiter</b><br/><span style='font-size:12px'>Sys: monitoring</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        monitoring.logging-service["<b>monitoring.logging-service</b><br/><span style='font-size:12px'>Sys: monitoring</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        data-store.redis-cache["<b>data-store.redis-cache</b><br/><span style='font-size:12px'>Sys: data-store</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        notification.email-service["<b>notification.email-service</b><br/><span style='font-size:12px'>Sys: notification</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        notification.sms-service["<b>notification.sms-service</b><br/><span style='font-size:12px'>Sys: notification</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        webhook.event-dispatcher["<b>webhook.event-dispatcher</b><br/><span style='font-size:12px'>Sys: webhook</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        integration.partner-api["<b>integration.partner-api</b><br/><span style='font-size:12px'>Sys: integration</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        security.secrets-manager["<b>security.secrets-manager</b><br/><span style='font-size:12px'>Sys: security</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        cache.invalidation-service["<b>cache.invalidation-service</b><br/><span style='font-size:12px'>Sys: cache</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        queue.message-broker["<b>queue.message-broker</b><br/><span style='font-size:12px'>Sys: queue</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        async.event-processor["<b>async.event-processor</b><br/><span style='font-size:12px'>Sys: async</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        search.customer-search["<b>search.customer-search</b><br/><span style='font-size:12px'>Sys: search</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        mobile.push-notification["<b>mobile.push-notification</b><br/><span style='font-size:12px'>Sys: mobile</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        document.generation-service["<b>document.generation-service</b><br/><span style='font-size:12px'>Sys: document</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        document.archive-service["<b>document.archive-service</b><br/><span style='font-size:12px'>Sys: document</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        translation.i18n-service["<b>translation.i18n-service</b><br/><span style='font-size:12px'>Sys: translation</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        backup.disaster-recovery["<b>backup.disaster-recovery</b><br/><span style='font-size:12px'>Sys: backup</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        test.environment-provisioner["<b>test.environment-provisioner</b><br/><span style='font-size:12px'>Sys: test</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class core.crypto,core.auth-service,api-gateway.router,monitoring.rate-limiter,monitoring.logging-service,data-store.redis-cache,notification.email-service,notification.sms-service,webhook.event-dispatcher,integration.partner-api,security.secrets-manager,cache.invalidation-service,queue.message-broker,async.event-processor,search.customer-search,mobile.push-notification,document.generation-service,document.archive-service,translation.i18n-service,backup.disaster-recovery,test.environment-provisioner blue;

    subgraph products [Time: PRODUCTS]
    direction TB
        savings.account-service["<b>savings.account-service</b><br/><span style='font-size:12px'>Sys: savings</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        savings.interest-calculator["<b>savings.interest-calculator</b><br/><span style='font-size:12px'>Sys: interest</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        insurance.policy-engine["<b>insurance.policy-engine</b><br/><span style='font-size:12px'>Sys: insurance</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        subscription.billing-engine["<b>subscription.billing-engine</b><br/><span style='font-size:12px'>Sys: subscription</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class savings.account-service,savings.interest-calculator,insurance.policy-engine,subscription.billing-engine blue;

    subgraph risk-analytics [Time: RISK-ANALYTICS]
    direction TB
        score.mule-account["<b>score.mule-account</b><br/><span style='font-size:12px'>Sys: score</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        risk.rating-service["<b>risk.rating-service</b><br/><span style='font-size:12px'>Sys: risk</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        risk.exposure-monitor["<b>risk.exposure-monitor</b><br/><span style='font-size:12px'>Sys: risk</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        limit.management-service["<b>limit.management-service</b><br/><span style='font-size:12px'>Sys: limit</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        network.graph-analysis["<b>network.graph-analysis</b><br/><span style='font-size:12px'>Sys: network</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class score.mule-account,risk.rating-service,risk.exposure-monitor,limit.management-service,network.graph-analysis blue;

    subgraph payments [Time: PAYMENTS]
    direction TB
        payment.pix-api["<b>payment.pix-api</b><br/><span style='font-size:12px'>Sys: payment</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        payment.ted-processor["<b>payment.ted-processor</b><br/><span style='font-size:12px'>Sys: payment</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        payment.settlement-api["<b>payment.settlement-api</b><br/><span style='font-size:12px'>Sys: payment</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        transfer.domestic-service["<b>transfer.domestic-service</b><br/><span style='font-size:12px'>Sys: transfer</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        transfer.international-service["<b>transfer.international-service</b><br/><span style='font-size:12px'>Sys: transfer</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        merchant.settlement-service["<b>merchant.settlement-service</b><br/><span style='font-size:12px'>Sys: merchant</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        reversal.transaction-service["<b>reversal.transaction-service</b><br/><span style='font-size:12px'>Sys: reversal</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        refund.processor["<b>refund.processor</b><br/><span style='font-size:12px'>Sys: refund</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        localization.currency-converter["<b>localization.currency-converter</b><br/><span style='font-size:12px'>Sys: localization</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class payment.pix-api,payment.ted-processor,payment.settlement-api,transfer.domestic-service,transfer.international-service,merchant.settlement-service,reversal.transaction-service,refund.processor,localization.currency-converter blue;

    subgraph card-platform [Time: CARD-PLATFORM]
    direction TB
        card.core-service["<b>card.core-service</b><br/><span style='font-size:12px'>Sys: card</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        card.tokenizer["<b>card.tokenizer</b><br/><span style='font-size:12px'>Sys: card</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
        card.authorization-engine["<b>card.authorization-engine</b><br/><span style='font-size:12px'>Sys: card</span><br/><br/><a href='' style='color:#60a5fa;text-decoration:none'>Github</a> &nbsp; <a href='' style='color:#60a5fa;text-decoration:none'>Confluence</a>"]
    end
    class card.core-service,card.tokenizer,card.authorization-engine blue;

    %% Dependências
    fraud-engine.viper-api --> fraud-engine.reference-data
    fraud-engine.viper-api --> fraud-solution.risk-location
    fraud-engine.viper-bff --> fraud-engine.viper-api
    fraud-engine.viper-bff --> fraud-engine.events-api
    transfer-antifraud.pix-out --> checking-account.account-service
    transfer-antifraud.pix-out --> account.dict
    transfer-antifraud.pix-out --> score.mule-account
    transfer-antifraud.pix-out --> fraud-engine.viper-api
    fraud-engine.events-api --> fraud-engine.reference-data
    checking-account.account-service --> account.dict
    checking-account.account-service --> core.crypto
    score.mule-account --> ml.model-inference
    score.mule-account --> data-lake.customer-events
    payment.pix-api --> fraud-engine.viper-api
    payment.pix-api --> transfer-antifraud.pix-out
    payment.pix-api --> checking-account.account-service
    payment.ted-processor --> payment.settlement-api
    payment.ted-processor --> compliance.aml-check
    payment.settlement-api --> accounting.ledger
    card.core-service --> account.dict
    card.core-service --> card.tokenizer
    card.tokenizer --> core.crypto
    card.authorization-engine --> fraud-engine.viper-api
    card.authorization-engine --> card.core-service
    credit.product-service --> account.dict
    credit.product-service --> credit.pricing-engine
    credit.product-service --> compliance.kyc
    credit.pricing-engine --> score.mule-account
    credit.pricing-engine --> risk.rating-service
    credit.disbursement-api --> credit.product-service
    credit.disbursement-api --> payment.settlement-api
    loan.origination-api --> credit.product-service
    loan.origination-api --> compliance.kyc
    loan.origination-api --> score.mule-account
    loan.servicing-engine --> loan.origination-api
    loan.servicing-engine --> accounting.ledger
    compliance.kyc --> account.dict
    compliance.kyc --> external.document-validator
    compliance.aml-check --> external.sanctions-list
    compliance.aml-check --> data-lake.customer-events
    risk.rating-service --> score.mule-account
    risk.rating-service --> data-lake.customer-events
    risk.exposure-monitor --> risk.rating-service
    risk.exposure-monitor --> accounting.ledger
    accounting.ledger --> core.crypto
    accounting.reconciliation --> accounting.ledger
    accounting.reconciliation --> data-lake.transactions
    core.auth-service --> core.crypto
    api-gateway.router --> core.auth-service
    api-gateway.router --> monitoring.rate-limiter
    monitoring.rate-limiter --> data-store.redis-cache
    ml.model-inference --> data-lake.customer-events
    ml.model-inference --> data-lake.transactions
    notification.email-service --> account.dict
    notification.sms-service --> account.dict
    transfer.domestic-service --> payment.pix-api
    transfer.domestic-service --> fraud-engine.viper-api
    transfer.international-service --> payment.ted-processor
    transfer.international-service --> compliance.aml-check
    customer.profile-api --> account.dict
    customer.profile-api --> notification.email-service
    customer.preferences-service --> customer.profile-api
    customer.preferences-service --> data-store.redis-cache
    savings.account-service --> account.dict
    savings.account-service --> interest.calculator
    class interest.calculator external;
    savings.interest-calculator --> accounting.ledger
    investment.portfolio-api --> account.dict
    investment.portfolio-api --> external.market-data
    insurance.policy-engine --> account.dict
    insurance.policy-engine --> risk.rating-service
    reporting.statement-generator --> accounting.ledger
    reporting.statement-generator --> customer.profile-api
    reporting.analytics-dashboard --> data-lake.transactions
    reporting.analytics-dashboard --> data-lake.customer-events
    webhook.event-dispatcher --> monitoring.logging-service
    webhook.event-dispatcher --> data-store.redis-cache
    integration.partner-api --> core.auth-service
    integration.partner-api --> webhook.event-dispatcher
    security.audit-logger --> monitoring.logging-service
    cache.invalidation-service --> data-store.redis-cache
    batch.scheduled-jobs --> accounting.reconciliation
    batch.scheduled-jobs --> data-lake.transactions
    async.event-processor --> queue.message-broker
    async.event-processor --> monitoring.logging-service
    merchant.onboarding-api --> compliance.kyc
    merchant.onboarding-api --> account.dict
    merchant.settlement-service --> payment.settlement-api
    merchant.settlement-service --> merchant.onboarding-api
    affiliate.commission-calculator --> accounting.ledger
    affiliate.commission-calculator --> data-lake.transactions
    subscription.billing-engine --> account.dict
    subscription.billing-engine --> payment.pix-api
    promotion.offers-api --> customer.preferences-service
    promotion.offers-api --> subscription.billing-engine
    analytics.customer-journey --> data-lake.customer-events
    analytics.customer-journey --> ml.model-inference
    search.customer-search --> customer.profile-api
    search.customer-search --> data-store.redis-cache
    mobile.push-notification --> notification.sms-service
    mobile.push-notification --> customer.profile-api
    document.generation-service --> customer.profile-api
    document.generation-service --> accounting.ledger
    document.archive-service --> document.generation-service
    dispute.management-api --> card.core-service
    dispute.management-api --> payment.pix-api
    dispute.management-api --> security.audit-logger
    chargeback.processor --> dispute.management-api
    chargeback.processor --> payment.settlement-api
    reversal.transaction-service --> payment.settlement-api
    reversal.transaction-service --> accounting.ledger
    refund.processor --> reversal.transaction-service
    refund.processor --> notification.email-service
    limit.management-service --> customer.profile-api
    limit.management-service --> score.mule-account
    limit.management-service --> accounting.ledger
    velocity.check-engine --> data-store.redis-cache
    velocity.check-engine --> fraud-engine.viper-api
    behavior.anomaly-detector --> ml.model-inference
    behavior.anomaly-detector --> data-lake.customer-events
    behavior.anomaly-detector --> fraud-engine.viper-api
    network.graph-analysis --> data-lake.transactions
    network.graph-analysis --> ml.model-inference
    feedback.survey-api --> customer.profile-api
    feedback.survey-api --> data-lake.customer-events
    rating.service-quality --> feedback.survey-api
    rating.service-quality --> customer.profile-api
    chatbot.conversation-engine --> customer.profile-api
    chatbot.conversation-engine --> ml.model-inference
    help-desk.ticket-management --> customer.profile-api
    help-desk.ticket-management --> security.audit-logger
    knowledge-base.article-service --> search.customer-search
    localization.currency-converter --> external.market-data
    localization.currency-converter --> data-store.redis-cache
    tax.calculation-engine --> accounting.ledger
    tax.calculation-engine --> customer.profile-api
    invoice.generation-api --> document.generation-service
    invoice.generation-api --> accounting.ledger
    expense.management-api --> accounting.ledger
    expense.management-api --> employee.registry
    payroll.processing-engine --> employee.registry
    payroll.processing-engine --> accounting.ledger
    attendance.tracking-system --> employee.registry
    leave.management-api --> employee.registry
    leave.management-api --> notification.email-service
    performance.evaluation-system --> employee.registry
    performance.evaluation-system --> feedback.survey-api
    training.learning-platform --> employee.registry
    training.learning-platform --> document.generation-service
    recruitment.job-portal --> employee.registry
    asset.management-system --> employee.registry
    asset.management-system --> accounting.ledger
    facility.booking-api --> employee.registry
    facility.booking-api --> notification.email-service
    office.expense-tracking --> expense.management-api
    office.expense-tracking --> employee.registry
    vendor.management-api --> compliance.kyc
    vendor.management-api --> accounting.ledger
    procurement.order-system --> vendor.management-api
    procurement.order-system --> inventory.stock-management
    inventory.stock-management --> accounting.ledger
    logistics.shipment-tracking --> vendor.management-api
    logistics.shipment-tracking --> notification.email-service
    warranty.claim-processor --> customer.profile-api
    warranty.claim-processor --> inventory.stock-management
    migration.data-sync --> data-lake.customer-events
    migration.data-sync --> data-lake.transactions
    test.environment-provisioner --> core.auth-service
