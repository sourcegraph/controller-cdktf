package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformations struct {
	// primitive_transformation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_deidentify_template#primitive_transformation GoogleDataLossPreventionDeidentifyTemplate#primitive_transformation}
	PrimitiveTransformation *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation `field:"required" json:"primitiveTransformation" yaml:"primitiveTransformation"`
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_deidentify_template#info_types GoogleDataLossPreventionDeidentifyTemplate#info_types}
	InfoTypes interface{} `field:"optional" json:"infoTypes" yaml:"infoTypes"`
}

