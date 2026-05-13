package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDocumentQuerySource struct {
	// Knowledge documents to query from. Format: projects/<Project ID>/locations/<Location ID>/knowledgeBases/<KnowledgeBase ID>/documents/<Document ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#documents GoogleDialogflowConversationProfile#documents}
	Documents *[]*string `field:"required" json:"documents" yaml:"documents"`
}

