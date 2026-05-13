package googledialogflowcxgenerativesettings


type GoogleDialogflowCxGenerativeSettingsLlmModelSettings struct {
	// The selected LLM model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#model GoogleDialogflowCxGenerativeSettings#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The custom prompt to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#prompt_text GoogleDialogflowCxGenerativeSettings#prompt_text}
	PromptText *string `field:"optional" json:"promptText" yaml:"promptText"`
}

