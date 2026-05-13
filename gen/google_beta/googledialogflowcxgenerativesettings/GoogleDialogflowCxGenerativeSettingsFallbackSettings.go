package googledialogflowcxgenerativesettings


type GoogleDialogflowCxGenerativeSettingsFallbackSettings struct {
	// prompt_templates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#prompt_templates GoogleDialogflowCxGenerativeSettings#prompt_templates}
	PromptTemplates interface{} `field:"optional" json:"promptTemplates" yaml:"promptTemplates"`
	// Display name of the selected prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generative_settings#selected_prompt GoogleDialogflowCxGenerativeSettings#selected_prompt}
	SelectedPrompt *string `field:"optional" json:"selectedPrompt" yaml:"selectedPrompt"`
}

