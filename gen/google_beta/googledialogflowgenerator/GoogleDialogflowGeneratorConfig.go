package googledialogflowgenerator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDialogflowGeneratorConfig struct {
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
	// desc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#location GoogleDialogflowGenerator#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// summarization_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#summarization_context GoogleDialogflowGenerator#summarization_context}
	SummarizationContext *GoogleDialogflowGeneratorSummarizationContext `field:"required" json:"summarizationContext" yaml:"summarizationContext"`
	// Optional. Human readable description of the generator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#description GoogleDialogflowGenerator#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional. The ID to use for the generator, which will become the final component of the generator's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#generator_id GoogleDialogflowGenerator#generator_id}
	GeneratorId *string `field:"optional" json:"generatorId" yaml:"generatorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#id GoogleDialogflowGenerator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inference_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#inference_parameter GoogleDialogflowGenerator#inference_parameter}
	InferenceParameter *GoogleDialogflowGeneratorInferenceParameter `field:"optional" json:"inferenceParameter" yaml:"inferenceParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#project GoogleDialogflowGenerator#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Optional.
	//
	// The published Large Language Model name. * To use the latest model version, specify the model name without version number. Example: text-bison * To use a stable model version, specify the version number as well. Example: text-bison@002.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#published_model GoogleDialogflowGenerator#published_model}
	PublishedModel *string `field:"optional" json:"publishedModel" yaml:"publishedModel"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#timeouts GoogleDialogflowGenerator#timeouts}
	Timeouts *GoogleDialogflowGeneratorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional.
	//
	// The trigger event of the generator. It defines when the generator is triggered in a conversation. Possible values: ["END_OF_UTTERANCE", "MANUAL_CALL", "CUSTOMER_MESSAGE", "AGENT_MESSAGE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#trigger_event GoogleDialogflowGenerator#trigger_event}
	TriggerEvent *string `field:"optional" json:"triggerEvent" yaml:"triggerEvent"`
}

