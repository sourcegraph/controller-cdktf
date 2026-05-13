package googlemodelarmortemplate


type GoogleModelArmorTemplateTemplateMetadata struct {
	// Indicates the custom error code set by the user to be returned to the end user if the LLM response trips Model Armor filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#custom_llm_response_safety_error_code GoogleModelArmorTemplate#custom_llm_response_safety_error_code}
	CustomLlmResponseSafetyErrorCode *float64 `field:"optional" json:"customLlmResponseSafetyErrorCode" yaml:"customLlmResponseSafetyErrorCode"`
	// Indicates the custom error message set by the user to be returned to the end user if the LLM response trips Model Armor filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#custom_llm_response_safety_error_message GoogleModelArmorTemplate#custom_llm_response_safety_error_message}
	CustomLlmResponseSafetyErrorMessage *string `field:"optional" json:"customLlmResponseSafetyErrorMessage" yaml:"customLlmResponseSafetyErrorMessage"`
	// Indicates the custom error code set by the user to be returned to the end user by the service extension if the prompt trips Model Armor filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#custom_prompt_safety_error_code GoogleModelArmorTemplate#custom_prompt_safety_error_code}
	CustomPromptSafetyErrorCode *float64 `field:"optional" json:"customPromptSafetyErrorCode" yaml:"customPromptSafetyErrorCode"`
	// Indicates the custom error message set by the user to be returned to the end user if the prompt trips Model Armor filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#custom_prompt_safety_error_message GoogleModelArmorTemplate#custom_prompt_safety_error_message}
	CustomPromptSafetyErrorMessage *string `field:"optional" json:"customPromptSafetyErrorMessage" yaml:"customPromptSafetyErrorMessage"`
	// Possible values: INSPECT_ONLY INSPECT_AND_BLOCK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#enforcement_type GoogleModelArmorTemplate#enforcement_type}
	EnforcementType *string `field:"optional" json:"enforcementType" yaml:"enforcementType"`
	// If true, partial detector failures should be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#ignore_partial_invocation_failures GoogleModelArmorTemplate#ignore_partial_invocation_failures}
	IgnorePartialInvocationFailures interface{} `field:"optional" json:"ignorePartialInvocationFailures" yaml:"ignorePartialInvocationFailures"`
	// If true, log sanitize operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#log_sanitize_operations GoogleModelArmorTemplate#log_sanitize_operations}
	LogSanitizeOperations interface{} `field:"optional" json:"logSanitizeOperations" yaml:"logSanitizeOperations"`
	// If true, log template crud operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#log_template_operations GoogleModelArmorTemplate#log_template_operations}
	LogTemplateOperations interface{} `field:"optional" json:"logTemplateOperations" yaml:"logTemplateOperations"`
	// multi_language_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#multi_language_detection GoogleModelArmorTemplate#multi_language_detection}
	MultiLanguageDetection *GoogleModelArmorTemplateTemplateMetadataMultiLanguageDetection `field:"optional" json:"multiLanguageDetection" yaml:"multiLanguageDetection"`
}

