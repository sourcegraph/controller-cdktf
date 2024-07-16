package googledialogflowcxpage


type GoogleDialogflowCxPageAdvancedSettingsDtmfSettings struct {
	// If true, incoming audio is processed for DTMF (dual tone multi frequency) events.
	//
	// For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will detect the event (e.g. a "3" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_page#enabled GoogleDialogflowCxPage#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The digit that terminates a DTMF digit sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_page#finish_digit GoogleDialogflowCxPage#finish_digit}
	FinishDigit *string `field:"optional" json:"finishDigit" yaml:"finishDigit"`
	// Max length of DTMF digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_page#max_digits GoogleDialogflowCxPage#max_digits}
	MaxDigits *float64 `field:"optional" json:"maxDigits" yaml:"maxDigits"`
}

