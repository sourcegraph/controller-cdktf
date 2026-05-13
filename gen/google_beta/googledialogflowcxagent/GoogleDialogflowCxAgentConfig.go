package googledialogflowcxagent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDialogflowCxAgentConfig struct {
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
	// The default language of the agent as a language tag.
	//
	// [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#default_language_code GoogleDialogflowCxAgent#default_language_code}
	DefaultLanguageCode *string `field:"required" json:"defaultLanguageCode" yaml:"defaultLanguageCode"`
	// The human-readable name of the agent, unique within the location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#display_name GoogleDialogflowCxAgent#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the location this agent is located in.
	//
	// ~> **Note:** The first time you are deploying an Agent in your project you must configure location settings.
	//  This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	//  Another options is to use global location so you don't need to manually configure location settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#location GoogleDialogflowCxAgent#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#time_zone GoogleDialogflowCxAgent#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// advanced_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#advanced_settings GoogleDialogflowCxAgent#advanced_settings}
	AdvancedSettings *GoogleDialogflowCxAgentAdvancedSettings `field:"optional" json:"advancedSettings" yaml:"advancedSettings"`
	// answer_feedback_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#answer_feedback_settings GoogleDialogflowCxAgent#answer_feedback_settings}
	AnswerFeedbackSettings *GoogleDialogflowCxAgentAnswerFeedbackSettings `field:"optional" json:"answerFeedbackSettings" yaml:"answerFeedbackSettings"`
	// The URI of the agent's avatar.
	//
	// Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#avatar_uri GoogleDialogflowCxAgent#avatar_uri}
	AvatarUri *string `field:"optional" json:"avatarUri" yaml:"avatarUri"`
	// client_certificate_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#client_certificate_settings GoogleDialogflowCxAgent#client_certificate_settings}
	ClientCertificateSettings *GoogleDialogflowCxAgentClientCertificateSettings `field:"optional" json:"clientCertificateSettings" yaml:"clientCertificateSettings"`
	// If set to 'true', Terraform will delete the chat engine associated with the agent when the agent is destroyed.
	//
	// Otherwise, the chat engine will persist.
	//
	// This virtual field addresses a critical dependency chain: 'agent' -> 'engine' -> 'data store'. The chat engine is automatically
	// provisioned when a data store is linked to the agent, meaning Terraform doesn't have direct control over its lifecycle as a managed
	// resource. This creates a problem when both the agent and data store are managed by Terraform and need to be destroyed. Without
	// delete_chat_engine_on_destroy set to true, the data store's deletion would fail because the unmanaged chat engine would still be
	// using it. This setting ensures that the entire dependency chain can be properly torn down.
	// See 'mmv1/templates/terraform/examples/dialogflowcx_tool_data_store.tf.tmpl' as an example.
	//
	// Data store can be linked to an agent through the 'knowledgeConnectorSettings' field of a [flow](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows#resource:-flow)
	// or a [page](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows.pages#resource:-page)
	// or the 'dataStoreSpec' field of a [tool](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.tools#resource:-tool).
	// The ID of the implicitly created engine is stored in the 'genAppBuilderSettings' field of the [agent](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents#resource:-agent).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#delete_chat_engine_on_destroy GoogleDialogflowCxAgent#delete_chat_engine_on_destroy}
	DeleteChatEngineOnDestroy interface{} `field:"optional" json:"deleteChatEngineOnDestroy" yaml:"deleteChatEngineOnDestroy"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#description GoogleDialogflowCxAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable training multi-lingual models for this agent.
	//
	// These models will be trained on all the languages supported by the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#enable_multi_language_training GoogleDialogflowCxAgent#enable_multi_language_training}
	EnableMultiLanguageTraining interface{} `field:"optional" json:"enableMultiLanguageTraining" yaml:"enableMultiLanguageTraining"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#enable_spell_correction GoogleDialogflowCxAgent#enable_spell_correction}
	EnableSpellCorrection interface{} `field:"optional" json:"enableSpellCorrection" yaml:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#enable_stackdriver_logging GoogleDialogflowCxAgent#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
	// gen_app_builder_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#gen_app_builder_settings GoogleDialogflowCxAgent#gen_app_builder_settings}
	GenAppBuilderSettings *GoogleDialogflowCxAgentGenAppBuilderSettings `field:"optional" json:"genAppBuilderSettings" yaml:"genAppBuilderSettings"`
	// git_integration_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#git_integration_settings GoogleDialogflowCxAgent#git_integration_settings}
	GitIntegrationSettings *GoogleDialogflowCxAgentGitIntegrationSettings `field:"optional" json:"gitIntegrationSettings" yaml:"gitIntegrationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#id GoogleDialogflowCxAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether the agent is locked for changes.
	//
	// If the agent is locked, modifications to the agent will be rejected except for [agents.restore][].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#locked GoogleDialogflowCxAgent#locked}
	Locked interface{} `field:"optional" json:"locked" yaml:"locked"`
	// personalization_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#personalization_settings GoogleDialogflowCxAgent#personalization_settings}
	PersonalizationSettings *GoogleDialogflowCxAgentPersonalizationSettings `field:"optional" json:"personalizationSettings" yaml:"personalizationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#project GoogleDialogflowCxAgent#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#security_settings GoogleDialogflowCxAgent#security_settings}
	SecuritySettings *string `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// speech_to_text_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#speech_to_text_settings GoogleDialogflowCxAgent#speech_to_text_settings}
	SpeechToTextSettings *GoogleDialogflowCxAgentSpeechToTextSettings `field:"optional" json:"speechToTextSettings" yaml:"speechToTextSettings"`
	// Name of the start playbook in this agent.
	//
	// A start playbook will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: **projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/playbooks/<PlaybookID>**. Currently only the default playbook with id "00000000-0000-0000-0000-000000000000" is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#start_playbook GoogleDialogflowCxAgent#start_playbook}
	StartPlaybook *string `field:"optional" json:"startPlaybook" yaml:"startPlaybook"`
	// The list of all languages supported by this agent (except for the default_language_code).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#supported_language_codes GoogleDialogflowCxAgent#supported_language_codes}
	SupportedLanguageCodes *[]*string `field:"optional" json:"supportedLanguageCodes" yaml:"supportedLanguageCodes"`
	// text_to_speech_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#text_to_speech_settings GoogleDialogflowCxAgent#text_to_speech_settings}
	TextToSpeechSettings *GoogleDialogflowCxAgentTextToSpeechSettings `field:"optional" json:"textToSpeechSettings" yaml:"textToSpeechSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#timeouts GoogleDialogflowCxAgent#timeouts}
	Timeouts *GoogleDialogflowCxAgentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

