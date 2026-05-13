package googledialogflowcxtool


type GoogleDialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfig struct {
	// Optional.
	//
	// Indicate the auth token type generated from the Diglogflow service agent.
	// The generated token is sent in the Authorization header.
	// See [ServiceAgentAuth](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.tools#serviceagentauth) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#service_agent_auth GoogleDialogflowCxTool#service_agent_auth}
	ServiceAgentAuth *string `field:"optional" json:"serviceAgentAuth" yaml:"serviceAgentAuth"`
}

