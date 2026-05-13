package dialogflowcxgenerator


type DialogflowCxGeneratorLlmModelSettings struct {
	// The selected LLM model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_generator#model DialogflowCxGenerator#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The custom prompt to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_generator#prompt_text DialogflowCxGenerator#prompt_text}
	PromptText *string `field:"optional" json:"promptText" yaml:"promptText"`
}

