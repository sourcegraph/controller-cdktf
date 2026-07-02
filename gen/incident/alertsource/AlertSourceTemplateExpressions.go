package alertsource

type AlertSourceTemplateExpressions struct {
	// The human readable label of the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#label AlertSource#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// The operations to execute in sequence for this expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#operations AlertSource#operations}
	Operations any `field:"required" json:"operations" yaml:"operations"`
	// A short ID that can be used to reference the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#reference AlertSource#reference}
	Reference *string `field:"required" json:"reference" yaml:"reference"`
	// The root reference for this expression (i.e. where the expression starts).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#root_reference AlertSource#root_reference}
	RootReference *string `field:"required" json:"rootReference" yaml:"rootReference"`
	// The else branch to resort to if all operations fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#else_branch AlertSource#else_branch}
	ElseBranch *AlertSourceTemplateExpressionsElseBranch `field:"optional" json:"elseBranch" yaml:"elseBranch"`
}
