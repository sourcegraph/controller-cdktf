package alertroute


type AlertRouteIncidentTemplateCustomFieldsBindingArrayValue struct {
	// If set, this is the literal value of the step parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#literal AlertRoute#literal}
	Literal *string `field:"optional" json:"literal" yaml:"literal"`
	// If set, this is the reference into the trigger scope that is the value of this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#reference AlertRoute#reference}
	Reference *string `field:"optional" json:"reference" yaml:"reference"`
}

