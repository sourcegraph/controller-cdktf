package dialogflowconversationprofile


type DialogflowConversationProfileNewRecognitionResultNotificationConfig struct {
	// Format of message. Possible values: ["MESSAGE_FORMAT_UNSPECIFIED", "PROTO", "JSON"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_conversation_profile#message_format DialogflowConversationProfile#message_format}
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// Name of the Pub/Sub topic to publish conversation events like CONVERSATION_STARTED as serialized ConversationEvent protos.
	//
	// For telephony integration to receive notification, make sure either this topic is in the same project as the conversation or you grant service-<Conversation Project Number>@gcp-sa-dialogflow.iam.gserviceaccount.com the Dialogflow Service Agent role in the topic project.
	// For chat integration to receive notification, make sure API caller has been granted the Dialogflow Service Agent role for the topic.
	// Format: projects/<Project ID>/locations/<Location ID>/topics/<Topic ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_conversation_profile#topic DialogflowConversationProfile#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

