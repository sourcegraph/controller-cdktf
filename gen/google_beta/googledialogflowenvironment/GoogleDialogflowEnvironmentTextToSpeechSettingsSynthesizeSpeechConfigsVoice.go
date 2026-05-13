package googledialogflowenvironment


type GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoice struct {
	// The name of the voice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#name GoogleDialogflowEnvironment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The preferred gender of the voice. Possible values: ["SSML_VOICE_GENDER_UNSPECIFIED", "SSML_VOICE_GENDER_MALE", "SSML_VOICE_GENDER_FEMALE", "SSML_VOICE_GENDER_NEUTRAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#ssml_gender GoogleDialogflowEnvironment#ssml_gender}
	SsmlGender *string `field:"optional" json:"ssmlGender" yaml:"ssmlGender"`
}

