package alertroute


type AlertRouteExpressions struct {
	// The human readable label of the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#label AlertRoute#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// The operations to execute in sequence for this expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#operations AlertRoute#operations}
	Operations interface{} `field:"required" json:"operations" yaml:"operations"`
	// A short ID that can be used to reference the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#reference AlertRoute#reference}
	Reference *string `field:"required" json:"reference" yaml:"reference"`
	// The root reference for this expression (i.e. where the expression starts).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#root_reference AlertRoute#root_reference}
	RootReference *string `field:"required" json:"rootReference" yaml:"rootReference"`
	// The else branch to resort to if all operations fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#else_branch AlertRoute#else_branch}
	ElseBranch *AlertRouteExpressionsElseBranch `field:"optional" json:"elseBranch" yaml:"elseBranch"`
}

