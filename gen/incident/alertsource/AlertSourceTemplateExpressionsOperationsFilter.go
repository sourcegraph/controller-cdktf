package alertsource


type AlertSourceTemplateExpressionsOperationsFilter struct {
	// Groups of prerequisite conditions. All conditions in at least one group must be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#condition_groups AlertSource#condition_groups}
	ConditionGroups interface{} `field:"required" json:"conditionGroups" yaml:"conditionGroups"`
}

