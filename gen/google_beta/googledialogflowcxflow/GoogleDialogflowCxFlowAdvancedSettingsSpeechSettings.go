package googledialogflowcxflow


type GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings struct {
	// Sensitivity of the speech model that detects the end of speech. Scale from 0 to 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dialogflow_cx_flow#endpointer_sensitivity GoogleDialogflowCxFlow#endpointer_sensitivity}
	EndpointerSensitivity *float64 `field:"optional" json:"endpointerSensitivity" yaml:"endpointerSensitivity"`
	// Mapping from language to Speech-to-Text model.
	//
	// The mapped Speech-to-Text model will be selected for requests from its corresponding language. For more information, see [Speech models](https://cloud.google.com/dialogflow/cx/docs/concept/speech-models).
	// An object containing a list of **"key": value** pairs. Example: **{ "name": "wrench", "mass": "1.3kg", "count": "3" }**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dialogflow_cx_flow#models GoogleDialogflowCxFlow#models}
	Models *map[string]*string `field:"optional" json:"models" yaml:"models"`
	// Timeout before detecting no speech. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dialogflow_cx_flow#no_speech_timeout GoogleDialogflowCxFlow#no_speech_timeout}
	NoSpeechTimeout *string `field:"optional" json:"noSpeechTimeout" yaml:"noSpeechTimeout"`
	// Use timeout based endpointing, interpreting endpointer sensitivity as seconds of timeout value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dialogflow_cx_flow#use_timeout_based_endpointing GoogleDialogflowCxFlow#use_timeout_based_endpointing}
	UseTimeoutBasedEndpointing interface{} `field:"optional" json:"useTimeoutBasedEndpointing" yaml:"useTimeoutBasedEndpointing"`
}

