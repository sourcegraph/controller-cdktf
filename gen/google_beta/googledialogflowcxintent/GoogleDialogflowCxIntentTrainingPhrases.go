package googledialogflowcxintent


type GoogleDialogflowCxIntentTrainingPhrases struct {
	// parts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_intent#parts GoogleDialogflowCxIntent#parts}
	Parts interface{} `field:"required" json:"parts" yaml:"parts"`
	// Indicates how many times this example was added to the intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_intent#repeat_count GoogleDialogflowCxIntent#repeat_count}
	RepeatCount *float64 `field:"optional" json:"repeatCount" yaml:"repeatCount"`
}

