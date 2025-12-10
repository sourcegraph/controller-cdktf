package alertroute


type AlertRouteIncidentTemplateCustomFields struct {
	// Binding for the custom field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#binding AlertRoute#binding}
	Binding *AlertRouteIncidentTemplateCustomFieldsBinding `field:"required" json:"binding" yaml:"binding"`
	// ID of the custom field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#custom_field_id AlertRoute#custom_field_id}
	CustomFieldId *string `field:"required" json:"customFieldId" yaml:"customFieldId"`
	// The strategy to use when multiple alerts match this route. Possible values are: `first-wins`, `last-wins`, `append`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#merge_strategy AlertRoute#merge_strategy}
	MergeStrategy *string `field:"required" json:"mergeStrategy" yaml:"mergeStrategy"`
}

