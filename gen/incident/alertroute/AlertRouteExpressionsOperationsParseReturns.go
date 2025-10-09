package alertroute


type AlertRouteExpressionsOperationsParseReturns struct {
	// Whether the return value should be single or multi-value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#array AlertRoute#array}
	Array interface{} `field:"required" json:"array" yaml:"array"`
	// Expected return type of this expression (what to try casting the result to).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#type AlertRoute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

