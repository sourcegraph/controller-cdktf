package googledialogflowcxagent


type GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettings struct {
	// The access token used to authenticate the access to the GitHub repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_agent#access_token GoogleDialogflowCxAgent#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// A list of branches configured to be used from Dialogflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_agent#branches GoogleDialogflowCxAgent#branches}
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// The unique repository display name for the GitHub repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_agent#display_name GoogleDialogflowCxAgent#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The GitHub repository URI related to the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_agent#repository_uri GoogleDialogflowCxAgent#repository_uri}
	RepositoryUri *string `field:"optional" json:"repositoryUri" yaml:"repositoryUri"`
	// The branch of the GitHub repository tracked for this agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_agent#tracking_branch GoogleDialogflowCxAgent#tracking_branch}
	TrackingBranch *string `field:"optional" json:"trackingBranch" yaml:"trackingBranch"`
}

