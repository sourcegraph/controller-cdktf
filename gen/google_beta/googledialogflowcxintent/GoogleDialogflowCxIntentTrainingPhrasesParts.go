package googledialogflowcxintent


type GoogleDialogflowCxIntentTrainingPhrasesParts struct {
	// The text for this part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_intent#text GoogleDialogflowCxIntent#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// The parameter used to annotate this part of the training phrase.
	//
	// This field is required for annotated parts of the training phrase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_intent#parameter_id GoogleDialogflowCxIntent#parameter_id}
	ParameterId *string `field:"optional" json:"parameterId" yaml:"parameterId"`
}

