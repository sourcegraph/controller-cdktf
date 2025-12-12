package alertsource


type AlertSourceTemplateAttributes struct {
	// ID of the alert attribute to set with this binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#alert_attribute_id AlertSource#alert_attribute_id}
	AlertAttributeId *string `field:"required" json:"alertAttributeId" yaml:"alertAttributeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#binding AlertSource#binding}.
	Binding *AlertSourceTemplateAttributesBinding `field:"required" json:"binding" yaml:"binding"`
}

