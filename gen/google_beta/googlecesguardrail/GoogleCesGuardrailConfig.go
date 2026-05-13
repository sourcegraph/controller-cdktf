package googlecesguardrail

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCesGuardrailConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#app GoogleCesGuardrail#app}
	App *string `field:"required" json:"app" yaml:"app"`
	// Display name of the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#display_name GoogleCesGuardrail#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID to use for the guardrail, which will become the final component of the guardrail's resource name.
	//
	// If not provided, a unique ID will be
	// automatically assigned for the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#guardrail_id GoogleCesGuardrail#guardrail_id}
	GuardrailId *string `field:"required" json:"guardrailId" yaml:"guardrailId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#location GoogleCesGuardrail#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#action GoogleCesGuardrail#action}
	Action *GoogleCesGuardrailAction `field:"optional" json:"action" yaml:"action"`
	// code_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#code_callback GoogleCesGuardrail#code_callback}
	CodeCallback *GoogleCesGuardrailCodeCallback `field:"optional" json:"codeCallback" yaml:"codeCallback"`
	// content_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#content_filter GoogleCesGuardrail#content_filter}
	ContentFilter *GoogleCesGuardrailContentFilter `field:"optional" json:"contentFilter" yaml:"contentFilter"`
	// Description of the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#description GoogleCesGuardrail#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the guardrail is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#enabled GoogleCesGuardrail#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#id GoogleCesGuardrail#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// llm_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#llm_policy GoogleCesGuardrail#llm_policy}
	LlmPolicy *GoogleCesGuardrailLlmPolicy `field:"optional" json:"llmPolicy" yaml:"llmPolicy"`
	// llm_prompt_security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#llm_prompt_security GoogleCesGuardrail#llm_prompt_security}
	LlmPromptSecurity *GoogleCesGuardrailLlmPromptSecurity `field:"optional" json:"llmPromptSecurity" yaml:"llmPromptSecurity"`
	// model_safety block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#model_safety GoogleCesGuardrail#model_safety}
	ModelSafety *GoogleCesGuardrailModelSafety `field:"optional" json:"modelSafety" yaml:"modelSafety"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#project GoogleCesGuardrail#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#timeouts GoogleCesGuardrail#timeouts}
	Timeouts *GoogleCesGuardrailTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

