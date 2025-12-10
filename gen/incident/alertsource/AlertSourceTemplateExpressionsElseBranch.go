package alertsource


type AlertSourceTemplateExpressionsElseBranch struct {
	// The result assumed if the else branch is reached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#result AlertSource#result}
	Result *AlertSourceTemplateExpressionsElseBranchResult `field:"required" json:"result" yaml:"result"`
}

