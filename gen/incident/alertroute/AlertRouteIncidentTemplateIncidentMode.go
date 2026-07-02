package alertroute

type AlertRouteIncidentTemplateIncidentMode struct {
	// The array of literal or reference parameter values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#array_value AlertRoute#array_value}
	ArrayValue any `field:"optional" json:"arrayValue" yaml:"arrayValue"`
	// The literal or reference parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#value AlertRoute#value}
	Value *AlertRouteIncidentTemplateIncidentModeValue `field:"optional" json:"value" yaml:"value"`
}
