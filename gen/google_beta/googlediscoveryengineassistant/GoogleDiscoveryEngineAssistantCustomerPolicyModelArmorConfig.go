package googlediscoveryengineassistant


type GoogleDiscoveryEngineAssistantCustomerPolicyModelArmorConfig struct {
	// The resource name of the Model Armor template for sanitizing assistant responses. Format: 'projects/{project}/locations/{location}/templates/{template_id}'.
	//
	// If not specified, no sanitization will be applied to the assistant
	// response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#response_template GoogleDiscoveryEngineAssistant#response_template}
	ResponseTemplate *string `field:"required" json:"responseTemplate" yaml:"responseTemplate"`
	// The resource name of the Model Armor template for sanitizing user prompts. Format: 'projects/{project}/locations/{location}/templates/{template_id}'.
	//
	// If not specified, no sanitization will be applied to the user prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#user_prompt_template GoogleDiscoveryEngineAssistant#user_prompt_template}
	UserPromptTemplate *string `field:"required" json:"userPromptTemplate" yaml:"userPromptTemplate"`
	// Defines the failure mode for Model Armor sanitization. The supported values: 'FAIL_OPEN', 'FAIL_CLOSED'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#failure_mode GoogleDiscoveryEngineAssistant#failure_mode}
	FailureMode *string `field:"optional" json:"failureMode" yaml:"failureMode"`
}

