package googledialogflowcxgenerativesettings


type GoogleDialogflowCxGenerativeSettingsFallbackSettingsPromptTemplates struct {
	// Prompt name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#display_name GoogleDialogflowCxGenerativeSettings#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// If the flag is true, the prompt is frozen and cannot be modified by users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#frozen GoogleDialogflowCxGenerativeSettings#frozen}
	Frozen interface{} `field:"optional" json:"frozen" yaml:"frozen"`
	// Prompt text that is sent to a LLM on no-match default, placeholders are filled downstream.
	//
	// For example: "Here is a conversation $conversation, a response is: "
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#prompt_text GoogleDialogflowCxGenerativeSettings#prompt_text}
	PromptText *string `field:"optional" json:"promptText" yaml:"promptText"`
}

