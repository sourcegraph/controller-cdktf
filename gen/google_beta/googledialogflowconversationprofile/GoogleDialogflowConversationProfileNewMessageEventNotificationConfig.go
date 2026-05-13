package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileNewMessageEventNotificationConfig struct {
	// Format of the message Possible values: ["MESSAGE_FORMAT_UNSPECIFIED", "PROTO", "JSON"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#message_format GoogleDialogflowConversationProfile#message_format}
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// Name of the Pub/Sub topic to publish conversation events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#topic GoogleDialogflowConversationProfile#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

