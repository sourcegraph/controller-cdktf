package alertsource


type AlertSourceTemplate struct {
	// Attributes to set on alerts coming from this source, with a binding describing how to set them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#attributes AlertSource#attributes}
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#description AlertSource#description}.
	Description *AlertSourceTemplateDescription `field:"required" json:"description" yaml:"description"`
	// The expressions to be prepared for use by steps and conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#expressions AlertSource#expressions}
	Expressions interface{} `field:"required" json:"expressions" yaml:"expressions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#title AlertSource#title}.
	Title *AlertSourceTemplateTitle `field:"required" json:"title" yaml:"title"`
}

