package dialogflowcxtoolversion


type DialogflowCxToolVersionTool struct {
	// High level description of the Tool and its usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#description DialogflowCxToolVersion#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The human-readable name of the tool, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#display_name DialogflowCxToolVersion#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// data_store_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#data_store_spec DialogflowCxToolVersion#data_store_spec}
	DataStoreSpec *DialogflowCxToolVersionToolDataStoreSpec `field:"optional" json:"dataStoreSpec" yaml:"dataStoreSpec"`
	// function_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#function_spec DialogflowCxToolVersion#function_spec}
	FunctionSpec *DialogflowCxToolVersionToolFunctionSpec `field:"optional" json:"functionSpec" yaml:"functionSpec"`
	// open_api_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#open_api_spec DialogflowCxToolVersion#open_api_spec}
	OpenApiSpec *DialogflowCxToolVersionToolOpenApiSpec `field:"optional" json:"openApiSpec" yaml:"openApiSpec"`
}

