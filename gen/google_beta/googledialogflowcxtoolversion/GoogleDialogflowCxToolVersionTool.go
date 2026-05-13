package googledialogflowcxtoolversion


type GoogleDialogflowCxToolVersionTool struct {
	// High level description of the Tool and its usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#description GoogleDialogflowCxToolVersion#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The human-readable name of the tool, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#display_name GoogleDialogflowCxToolVersion#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// connector_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#connector_spec GoogleDialogflowCxToolVersion#connector_spec}
	ConnectorSpec *GoogleDialogflowCxToolVersionToolConnectorSpec `field:"optional" json:"connectorSpec" yaml:"connectorSpec"`
	// data_store_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#data_store_spec GoogleDialogflowCxToolVersion#data_store_spec}
	DataStoreSpec *GoogleDialogflowCxToolVersionToolDataStoreSpec `field:"optional" json:"dataStoreSpec" yaml:"dataStoreSpec"`
	// function_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#function_spec GoogleDialogflowCxToolVersion#function_spec}
	FunctionSpec *GoogleDialogflowCxToolVersionToolFunctionSpec `field:"optional" json:"functionSpec" yaml:"functionSpec"`
	// open_api_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#open_api_spec GoogleDialogflowCxToolVersion#open_api_spec}
	OpenApiSpec *GoogleDialogflowCxToolVersionToolOpenApiSpec `field:"optional" json:"openApiSpec" yaml:"openApiSpec"`
}

