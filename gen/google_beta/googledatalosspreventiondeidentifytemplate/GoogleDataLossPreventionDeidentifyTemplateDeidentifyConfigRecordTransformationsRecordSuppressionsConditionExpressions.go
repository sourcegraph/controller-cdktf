package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#conditions GoogleDataLossPreventionDeidentifyTemplate#conditions}
	Conditions *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// The operator to apply to the result of conditions.
	//
	// Default and currently only supported value is AND. Default value: "AND" Possible values: ["AND"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#logical_operator GoogleDataLossPreventionDeidentifyTemplate#logical_operator}
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
}

