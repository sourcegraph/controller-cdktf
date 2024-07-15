package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceDictionaryConfigWordListStruct struct {
	// Words or phrases defining the dictionary.
	//
	// The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#words GoogleDataLossPreventionDeidentifyTemplate#words}
	Words *[]*string `field:"required" json:"words" yaml:"words"`
}

