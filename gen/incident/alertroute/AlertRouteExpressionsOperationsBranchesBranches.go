package alertroute


type AlertRouteExpressionsOperationsBranchesBranches struct {
	// Groups of prerequisite conditions. All conditions in at least one group must be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#condition_groups AlertRoute#condition_groups}
	ConditionGroups interface{} `field:"required" json:"conditionGroups" yaml:"conditionGroups"`
	// The result assumed if the condition groups are satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#result AlertRoute#result}
	Result *AlertRouteExpressionsOperationsBranchesBranchesResult `field:"required" json:"result" yaml:"result"`
}

