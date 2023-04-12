package googledialogflowcxintent


type GoogleDialogflowCxIntentTrainingPhrases struct {
	// parts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dialogflow_cx_intent#parts GoogleDialogflowCxIntent#parts}
	Parts interface{} `field:"required" json:"parts" yaml:"parts"`
	// Indicates how many times this example was added to the intent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dialogflow_cx_intent#repeat_count GoogleDialogflowCxIntent#repeat_count}
	RepeatCount *float64 `field:"optional" json:"repeatCount" yaml:"repeatCount"`
}

