package googledialogflowconversationprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDialogflowConversationProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Required. Human readable name for this profile. Max length 1024 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#display_name GoogleDialogflowConversationProfile#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// desc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#location GoogleDialogflowConversationProfile#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// automated_agent_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#automated_agent_config GoogleDialogflowConversationProfile#automated_agent_config}
	AutomatedAgentConfig *GoogleDialogflowConversationProfileAutomatedAgentConfig `field:"optional" json:"automatedAgentConfig" yaml:"automatedAgentConfig"`
	// human_agent_assistant_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#human_agent_assistant_config GoogleDialogflowConversationProfile#human_agent_assistant_config}
	HumanAgentAssistantConfig *GoogleDialogflowConversationProfileHumanAgentAssistantConfig `field:"optional" json:"humanAgentAssistantConfig" yaml:"humanAgentAssistantConfig"`
	// human_agent_handoff_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#human_agent_handoff_config GoogleDialogflowConversationProfile#human_agent_handoff_config}
	HumanAgentHandoffConfig *GoogleDialogflowConversationProfileHumanAgentHandoffConfig `field:"optional" json:"humanAgentHandoffConfig" yaml:"humanAgentHandoffConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#id GoogleDialogflowConversationProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Language code for the conversation profile. This should be a BCP-47 language tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#language_code GoogleDialogflowConversationProfile#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#logging_config GoogleDialogflowConversationProfile#logging_config}
	LoggingConfig *GoogleDialogflowConversationProfileLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// new_message_event_notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#new_message_event_notification_config GoogleDialogflowConversationProfile#new_message_event_notification_config}
	NewMessageEventNotificationConfig *GoogleDialogflowConversationProfileNewMessageEventNotificationConfig `field:"optional" json:"newMessageEventNotificationConfig" yaml:"newMessageEventNotificationConfig"`
	// new_recognition_result_notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#new_recognition_result_notification_config GoogleDialogflowConversationProfile#new_recognition_result_notification_config}
	NewRecognitionResultNotificationConfig *GoogleDialogflowConversationProfileNewRecognitionResultNotificationConfig `field:"optional" json:"newRecognitionResultNotificationConfig" yaml:"newRecognitionResultNotificationConfig"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#notification_config GoogleDialogflowConversationProfile#notification_config}
	NotificationConfig *GoogleDialogflowConversationProfileNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#project GoogleDialogflowConversationProfile#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Name of the CX SecuritySettings reference for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#security_settings GoogleDialogflowConversationProfile#security_settings}
	SecuritySettings *string `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// stt_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#stt_config GoogleDialogflowConversationProfile#stt_config}
	SttConfig *GoogleDialogflowConversationProfileSttConfig `field:"optional" json:"sttConfig" yaml:"sttConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#timeouts GoogleDialogflowConversationProfile#timeouts}
	Timeouts *GoogleDialogflowConversationProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The time zone of this conversational profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#time_zone GoogleDialogflowConversationProfile#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// tts_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#tts_config GoogleDialogflowConversationProfile#tts_config}
	TtsConfig *GoogleDialogflowConversationProfileTtsConfig `field:"optional" json:"ttsConfig" yaml:"ttsConfig"`
	// Optional. Whether to use the bidi streaming API in telephony integration for the conversation profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#use_bidi_streaming GoogleDialogflowConversationProfile#use_bidi_streaming}
	UseBidiStreaming interface{} `field:"optional" json:"useBidiStreaming" yaml:"useBidiStreaming"`
}

