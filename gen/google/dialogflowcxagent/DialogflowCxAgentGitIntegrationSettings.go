package dialogflowcxagent


type DialogflowCxAgentGitIntegrationSettings struct {
	// github_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_agent#github_settings DialogflowCxAgent#github_settings}
	GithubSettings *DialogflowCxAgentGitIntegrationSettingsGithubSettings `field:"optional" json:"githubSettings" yaml:"githubSettings"`
}

