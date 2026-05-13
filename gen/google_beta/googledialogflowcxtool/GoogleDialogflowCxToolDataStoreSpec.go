package googledialogflowcxtool


type GoogleDialogflowCxToolDataStoreSpec struct {
	// data_store_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#data_store_connections GoogleDialogflowCxTool#data_store_connections}
	DataStoreConnections interface{} `field:"required" json:"dataStoreConnections" yaml:"dataStoreConnections"`
	// fallback_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#fallback_prompt GoogleDialogflowCxTool#fallback_prompt}
	FallbackPrompt *GoogleDialogflowCxToolDataStoreSpecFallbackPrompt `field:"required" json:"fallbackPrompt" yaml:"fallbackPrompt"`
}

