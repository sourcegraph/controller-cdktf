package alertroute


type AlertRouteExpressionsOperationsParse struct {
	// The return type of an operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#returns AlertRoute#returns}
	Returns *AlertRouteExpressionsOperationsParseReturns `field:"required" json:"returns" yaml:"returns"`
	// The ES5 Javascript expression to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#source AlertRoute#source}
	Source *string `field:"required" json:"source" yaml:"source"`
}

