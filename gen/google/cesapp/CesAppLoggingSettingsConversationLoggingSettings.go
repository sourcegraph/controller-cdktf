package cesapp

type CesAppLoggingSettingsConversationLoggingSettings struct {
	// Whether to disable conversation logging for the sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/ces_app#disable_conversation_logging CesApp#disable_conversation_logging}
	DisableConversationLogging any `field:"optional" json:"disableConversationLogging" yaml:"disableConversationLogging"`
}
