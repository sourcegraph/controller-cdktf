package cesapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CesAppConfig struct {
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
	// The ID to use for the app, which will become the final component of the app's resource name.
	//
	// If not provided, a unique ID will be
	// automatically assigned for the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#app_id CesApp#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Display name of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#display_name CesApp#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#location CesApp#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// audio_processing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#audio_processing_config CesApp#audio_processing_config}
	AudioProcessingConfig *CesAppAudioProcessingConfig `field:"optional" json:"audioProcessingConfig" yaml:"audioProcessingConfig"`
	// client_certificate_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#client_certificate_settings CesApp#client_certificate_settings}
	ClientCertificateSettings *CesAppClientCertificateSettings `field:"optional" json:"clientCertificateSettings" yaml:"clientCertificateSettings"`
	// data_store_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#data_store_settings CesApp#data_store_settings}
	DataStoreSettings *CesAppDataStoreSettings `field:"optional" json:"dataStoreSettings" yaml:"dataStoreSettings"`
	// default_channel_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#default_channel_profile CesApp#default_channel_profile}
	DefaultChannelProfile *CesAppDefaultChannelProfile `field:"optional" json:"defaultChannelProfile" yaml:"defaultChannelProfile"`
	// Human-readable description of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#description CesApp#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// evaluation_metrics_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#evaluation_metrics_thresholds CesApp#evaluation_metrics_thresholds}
	EvaluationMetricsThresholds *CesAppEvaluationMetricsThresholds `field:"optional" json:"evaluationMetricsThresholds" yaml:"evaluationMetricsThresholds"`
	// Instructions for all the agents in the app.
	//
	// You can use this instruction to set up a stable identity or personality
	// across all the agents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#global_instruction CesApp#global_instruction}
	GlobalInstruction *string `field:"optional" json:"globalInstruction" yaml:"globalInstruction"`
	// List of guardrails for the app. Format: 'projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#guardrails CesApp#guardrails}
	Guardrails *[]*string `field:"optional" json:"guardrails" yaml:"guardrails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#id CesApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// language_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#language_settings CesApp#language_settings}
	LanguageSettings *CesAppLanguageSettings `field:"optional" json:"languageSettings" yaml:"languageSettings"`
	// logging_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#logging_settings CesApp#logging_settings}
	LoggingSettings *CesAppLoggingSettings `field:"optional" json:"loggingSettings" yaml:"loggingSettings"`
	// Metadata about the app.
	//
	// This field can be used to store additional
	// information relevant to the app's details or intended usages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#metadata CesApp#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#model_settings CesApp#model_settings}
	ModelSettings *CesAppModelSettings `field:"optional" json:"modelSettings" yaml:"modelSettings"`
	// Whether the app is pinned in the app list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#pinned CesApp#pinned}
	Pinned interface{} `field:"optional" json:"pinned" yaml:"pinned"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#project CesApp#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The root agent is the entry point of the app. Format: 'projects/{project}/locations/{location}/apps/{app}/agents/{agent}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#root_agent CesApp#root_agent}
	RootAgent *string `field:"optional" json:"rootAgent" yaml:"rootAgent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#timeouts CesApp#timeouts}
	Timeouts *CesAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// time_zone_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#time_zone_settings CesApp#time_zone_settings}
	TimeZoneSettings *CesAppTimeZoneSettings `field:"optional" json:"timeZoneSettings" yaml:"timeZoneSettings"`
	// variable_declarations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#variable_declarations CesApp#variable_declarations}
	VariableDeclarations interface{} `field:"optional" json:"variableDeclarations" yaml:"variableDeclarations"`
}

