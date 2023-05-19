package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformations struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_data_loss_prevention_deidentify_template#fields GoogleDataLossPreventionDeidentifyTemplate#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// primitive_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_data_loss_prevention_deidentify_template#primitive_transformation GoogleDataLossPreventionDeidentifyTemplate#primitive_transformation}
	PrimitiveTransformation *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation `field:"required" json:"primitiveTransformation" yaml:"primitiveTransformation"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_data_loss_prevention_deidentify_template#condition GoogleDataLossPreventionDeidentifyTemplate#condition}
	Condition *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition `field:"optional" json:"condition" yaml:"condition"`
}

