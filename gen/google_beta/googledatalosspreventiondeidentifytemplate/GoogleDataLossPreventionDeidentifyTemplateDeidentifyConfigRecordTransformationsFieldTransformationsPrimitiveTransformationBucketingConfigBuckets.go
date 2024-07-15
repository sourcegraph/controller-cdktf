package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets struct {
	// replacement_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#replacement_value GoogleDataLossPreventionDeidentifyTemplate#replacement_value}
	ReplacementValue *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue `field:"required" json:"replacementValue" yaml:"replacementValue"`
	// max block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#max GoogleDataLossPreventionDeidentifyTemplate#max}
	Max *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax `field:"optional" json:"max" yaml:"max"`
	// min block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#min GoogleDataLossPreventionDeidentifyTemplate#min}
	Min *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin `field:"optional" json:"min" yaml:"min"`
}

