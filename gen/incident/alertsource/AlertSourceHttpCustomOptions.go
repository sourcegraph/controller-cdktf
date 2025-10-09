package alertsource


type AlertSourceHttpCustomOptions struct {
	// JSON path to extract the deduplication key from the payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#deduplication_key_path AlertSource#deduplication_key_path}
	DeduplicationKeyPath *string `field:"required" json:"deduplicationKeyPath" yaml:"deduplicationKeyPath"`
	// JavaScript expression that returns an object with all alert fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#transform_expression AlertSource#transform_expression}
	TransformExpression *string `field:"required" json:"transformExpression" yaml:"transformExpression"`
}

