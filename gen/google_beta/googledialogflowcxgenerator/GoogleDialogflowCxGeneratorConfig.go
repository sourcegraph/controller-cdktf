package googledialogflowcxgenerator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDialogflowCxGeneratorConfig struct {
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
	// The human-readable name of the generator, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#display_name GoogleDialogflowCxGenerator#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// prompt_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#prompt_text GoogleDialogflowCxGenerator#prompt_text}
	PromptText *GoogleDialogflowCxGeneratorPromptText `field:"required" json:"promptText" yaml:"promptText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#id GoogleDialogflowCxGenerator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The language to create generators for the following fields: * Generator.prompt_text.text If not specified, the agent's default language is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#language_code GoogleDialogflowCxGenerator#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// llm_model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#llm_model_settings GoogleDialogflowCxGenerator#llm_model_settings}
	LlmModelSettings *GoogleDialogflowCxGeneratorLlmModelSettings `field:"optional" json:"llmModelSettings" yaml:"llmModelSettings"`
	// model_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#model_parameter GoogleDialogflowCxGenerator#model_parameter}
	ModelParameter *GoogleDialogflowCxGeneratorModelParameter `field:"optional" json:"modelParameter" yaml:"modelParameter"`
	// The agent to create a Generator for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#parent GoogleDialogflowCxGenerator#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// placeholders block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#placeholders GoogleDialogflowCxGenerator#placeholders}
	Placeholders interface{} `field:"optional" json:"placeholders" yaml:"placeholders"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#timeouts GoogleDialogflowCxGenerator#timeouts}
	Timeouts *GoogleDialogflowCxGeneratorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

