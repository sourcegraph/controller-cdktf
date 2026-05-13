package dialogflowcxtoolversion


type DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig struct {
	// Optional.
	//
	// Indicate the auth token type generated from the Diglogflow service agent.
	// The generated token is sent in the Authorization header.
	// See [ServiceAgentAuth](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.tools#serviceagentauth) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#service_agent_auth DialogflowCxToolVersion#service_agent_auth}
	ServiceAgentAuth *string `field:"optional" json:"serviceAgentAuth" yaml:"serviceAgentAuth"`
}

