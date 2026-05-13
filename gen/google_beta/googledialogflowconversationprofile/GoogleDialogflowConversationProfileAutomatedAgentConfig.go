package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileAutomatedAgentConfig struct {
	// ID of the Dialogflow agent environment to use. Expects the format "projects/<Project ID>/locations/<Location ID>/agent/environments/<EnvironmentID>".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#agent GoogleDialogflowConversationProfile#agent}
	Agent *string `field:"required" json:"agent" yaml:"agent"`
	// Configure lifetime of the Dialogflow session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#session_ttl GoogleDialogflowConversationProfile#session_ttl}
	SessionTtl *string `field:"optional" json:"sessionTtl" yaml:"sessionTtl"`
}

