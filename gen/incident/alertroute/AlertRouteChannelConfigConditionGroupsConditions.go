package alertroute


type AlertRouteChannelConfigConditionGroupsConditions struct {
	// The logical operation to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#operation AlertRoute#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Bindings for the operation parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#param_bindings AlertRoute#param_bindings}
	ParamBindings interface{} `field:"required" json:"paramBindings" yaml:"paramBindings"`
	// The subject of the condition, on which the operation is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#subject AlertRoute#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

