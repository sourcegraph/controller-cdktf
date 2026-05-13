package googledialogflowcxgenerativesettings


type GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsBannedPhrases struct {
	// Language code of the phrase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#language_code GoogleDialogflowCxGenerativeSettings#language_code}
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Text input which can be used for prompt or banned phrases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#text GoogleDialogflowCxGenerativeSettings#text}
	Text *string `field:"required" json:"text" yaml:"text"`
}

