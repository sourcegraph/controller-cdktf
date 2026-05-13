package googledialogflowcxtool


type GoogleDialogflowCxToolOpenApiSpecServiceDirectoryConfig struct {
	// The name of [Service Directory](https://cloud.google.com/service-directory/docs) service. Format: projects/<ProjectID>/locations/<LocationID>/namespaces/<NamespaceID>/services/<ServiceID>. LocationID of the service directory must be the same as the location of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#service GoogleDialogflowCxTool#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

