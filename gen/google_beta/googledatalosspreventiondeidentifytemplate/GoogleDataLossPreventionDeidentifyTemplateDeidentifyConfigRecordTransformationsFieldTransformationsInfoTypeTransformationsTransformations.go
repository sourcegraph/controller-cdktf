package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations struct {
	// primitive_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#primitive_transformation GoogleDataLossPreventionDeidentifyTemplate#primitive_transformation}
	PrimitiveTransformation *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation `field:"required" json:"primitiveTransformation" yaml:"primitiveTransformation"`
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#info_types GoogleDataLossPreventionDeidentifyTemplate#info_types}
	InfoTypes interface{} `field:"optional" json:"infoTypes" yaml:"infoTypes"`
}

