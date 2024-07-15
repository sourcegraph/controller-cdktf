package googledialogflowcxwebhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDialogflowCxWebhookConfig struct {
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
	// The human-readable name of the webhook, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#display_name GoogleDialogflowCxWebhook#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Indicates whether the webhook is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#disabled GoogleDialogflowCxWebhook#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#enable_spell_correction GoogleDialogflowCxWebhook#enable_spell_correction}
	EnableSpellCorrection interface{} `field:"optional" json:"enableSpellCorrection" yaml:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#enable_stackdriver_logging GoogleDialogflowCxWebhook#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
	// generic_web_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#generic_web_service GoogleDialogflowCxWebhook#generic_web_service}
	GenericWebService *GoogleDialogflowCxWebhookGenericWebService `field:"optional" json:"genericWebService" yaml:"genericWebService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#id GoogleDialogflowCxWebhook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The agent to create a webhook for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#parent GoogleDialogflowCxWebhook#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#security_settings GoogleDialogflowCxWebhook#security_settings}
	SecuritySettings *string `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// service_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#service_directory GoogleDialogflowCxWebhook#service_directory}
	ServiceDirectory *GoogleDialogflowCxWebhookServiceDirectory `field:"optional" json:"serviceDirectory" yaml:"serviceDirectory"`
	// Webhook execution timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#timeout GoogleDialogflowCxWebhook#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#timeouts GoogleDialogflowCxWebhook#timeouts}
	Timeouts *GoogleDialogflowCxWebhookTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

