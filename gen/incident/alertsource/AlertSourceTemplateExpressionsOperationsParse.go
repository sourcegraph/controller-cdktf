package alertsource


type AlertSourceTemplateExpressionsOperationsParse struct {
	// The return type of an operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#returns AlertSource#returns}
	Returns *AlertSourceTemplateExpressionsOperationsParseReturns `field:"required" json:"returns" yaml:"returns"`
	// The ES5 Javascript expression to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#source AlertSource#source}
	Source *string `field:"required" json:"source" yaml:"source"`
}

