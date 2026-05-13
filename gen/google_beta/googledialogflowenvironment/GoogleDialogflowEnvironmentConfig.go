package googledialogflowenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDialogflowEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#environmentid GoogleDialogflowEnvironment#environmentid}.
	Environmentid *string `field:"required" json:"environmentid" yaml:"environmentid"`
	// The agent version loaded into this environment. Supported formats: - projects/<Project ID>/agent/versions/<Version ID> - projects/<Project ID>/locations/<Location ID>/agent/versions/<Version ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#agent_version GoogleDialogflowEnvironment#agent_version}
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// The developer-provided description for this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#description GoogleDialogflowEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#fulfillment GoogleDialogflowEnvironment#fulfillment}
	Fulfillment *GoogleDialogflowEnvironmentFulfillment `field:"optional" json:"fulfillment" yaml:"fulfillment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#id GoogleDialogflowEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#location GoogleDialogflowEnvironment#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#project GoogleDialogflowEnvironment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// text_to_speech_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#text_to_speech_settings GoogleDialogflowEnvironment#text_to_speech_settings}
	TextToSpeechSettings *GoogleDialogflowEnvironmentTextToSpeechSettings `field:"optional" json:"textToSpeechSettings" yaml:"textToSpeechSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#timeouts GoogleDialogflowEnvironment#timeouts}
	Timeouts *GoogleDialogflowEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

