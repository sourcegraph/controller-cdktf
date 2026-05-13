package googledialogflowcxgenerator


type GoogleDialogflowCxGeneratorLlmModelSettings struct {
	// The selected LLM model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#model GoogleDialogflowCxGenerator#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The custom prompt to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#prompt_text GoogleDialogflowCxGenerator#prompt_text}
	PromptText *string `field:"optional" json:"promptText" yaml:"promptText"`
}

