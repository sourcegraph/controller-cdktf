package googledialogflowcxgenerativesettings


type GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettings struct {
	// banned_phrases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#banned_phrases GoogleDialogflowCxGenerativeSettings#banned_phrases}
	BannedPhrases interface{} `field:"optional" json:"bannedPhrases" yaml:"bannedPhrases"`
	// Optional. Default phrase match strategy for banned phrases. See [PhraseMatchStrategy](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/GenerativeSettings#phrasematchstrategy) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#default_banned_phrase_match_strategy GoogleDialogflowCxGenerativeSettings#default_banned_phrase_match_strategy}
	DefaultBannedPhraseMatchStrategy *string `field:"optional" json:"defaultBannedPhraseMatchStrategy" yaml:"defaultBannedPhraseMatchStrategy"`
}

