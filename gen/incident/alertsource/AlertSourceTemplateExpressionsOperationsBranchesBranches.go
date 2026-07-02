package alertsource

type AlertSourceTemplateExpressionsOperationsBranchesBranches struct {
	// Groups of prerequisite conditions. All conditions in at least one group must be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#condition_groups AlertSource#condition_groups}
	ConditionGroups any `field:"required" json:"conditionGroups" yaml:"conditionGroups"`
	// The result assumed if the condition groups are satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#result AlertSource#result}
	Result *AlertSourceTemplateExpressionsOperationsBranchesBranchesResult `field:"required" json:"result" yaml:"result"`
}
