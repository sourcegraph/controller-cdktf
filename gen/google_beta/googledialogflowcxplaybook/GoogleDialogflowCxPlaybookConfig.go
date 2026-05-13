package googledialogflowcxplaybook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDialogflowCxPlaybookConfig struct {
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
	// The human-readable name of the playbook, unique within an agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#display_name GoogleDialogflowCxPlaybook#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// High level description of the goal the playbook intend to accomplish.
	//
	// A goal should be concise since it's visible to other playbooks that may reference this playbook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#goal GoogleDialogflowCxPlaybook#goal}
	Goal *string `field:"required" json:"goal" yaml:"goal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#id GoogleDialogflowCxPlaybook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instruction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#instruction GoogleDialogflowCxPlaybook#instruction}
	Instruction *GoogleDialogflowCxPlaybookInstruction `field:"optional" json:"instruction" yaml:"instruction"`
	// llm_model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#llm_model_settings GoogleDialogflowCxPlaybook#llm_model_settings}
	LlmModelSettings *GoogleDialogflowCxPlaybookLlmModelSettings `field:"optional" json:"llmModelSettings" yaml:"llmModelSettings"`
	// The agent to create a Playbook for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#parent GoogleDialogflowCxPlaybook#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// Type of the playbook. Possible values: ["PLAYBOOK_TYPE_UNSPECIFIED", "TASK", "ROUTINE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#playbook_type GoogleDialogflowCxPlaybook#playbook_type}
	PlaybookType *string `field:"optional" json:"playbookType" yaml:"playbookType"`
	// The resource name of tools referenced by the current playbook in the instructions.
	//
	// If not provided explicitly, they are will be implied using the tool being referenced in goal and steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#referenced_tools GoogleDialogflowCxPlaybook#referenced_tools}
	ReferencedTools *[]*string `field:"optional" json:"referencedTools" yaml:"referencedTools"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#timeouts GoogleDialogflowCxPlaybook#timeouts}
	Timeouts *GoogleDialogflowCxPlaybookTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

