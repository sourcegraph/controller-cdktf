package alertsource


type AlertSourceTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditions struct {
	// The logical operation to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#operation AlertSource#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Bindings for the operation parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#param_bindings AlertSource#param_bindings}
	ParamBindings interface{} `field:"required" json:"paramBindings" yaml:"paramBindings"`
	// The subject of the condition, on which the operation is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#subject AlertSource#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

