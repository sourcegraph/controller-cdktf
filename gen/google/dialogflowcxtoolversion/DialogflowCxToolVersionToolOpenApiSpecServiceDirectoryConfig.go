package dialogflowcxtoolversion


type DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfig struct {
	// The name of [Service Directory](https://cloud.google.com/service-directory/docs) service. Format: projects/<ProjectID>/locations/<LocationID>/namespaces/<NamespaceID>/services/<ServiceID>. LocationID of the service directory must be the same as the location of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#service DialogflowCxToolVersion#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

