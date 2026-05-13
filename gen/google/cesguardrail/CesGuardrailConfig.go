package cesguardrail

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CesGuardrailConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#app CesGuardrail#app}
	App *string `field:"required" json:"app" yaml:"app"`
	// Display name of the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#display_name CesGuardrail#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID to use for the guardrail, which will become the final component of the guardrail's resource name.
	//
	// If not provided, a unique ID will be
	// automatically assigned for the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#guardrail_id CesGuardrail#guardrail_id}
	GuardrailId *string `field:"required" json:"guardrailId" yaml:"guardrailId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#location CesGuardrail#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#action CesGuardrail#action}
	Action *CesGuardrailAction `field:"optional" json:"action" yaml:"action"`
	// code_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#code_callback CesGuardrail#code_callback}
	CodeCallback *CesGuardrailCodeCallback `field:"optional" json:"codeCallback" yaml:"codeCallback"`
	// content_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#content_filter CesGuardrail#content_filter}
	ContentFilter *CesGuardrailContentFilter `field:"optional" json:"contentFilter" yaml:"contentFilter"`
	// Description of the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#description CesGuardrail#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the guardrail is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#enabled CesGuardrail#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#id CesGuardrail#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// llm_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#llm_policy CesGuardrail#llm_policy}
	LlmPolicy *CesGuardrailLlmPolicy `field:"optional" json:"llmPolicy" yaml:"llmPolicy"`
	// llm_prompt_security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#llm_prompt_security CesGuardrail#llm_prompt_security}
	LlmPromptSecurity *CesGuardrailLlmPromptSecurity `field:"optional" json:"llmPromptSecurity" yaml:"llmPromptSecurity"`
	// model_safety block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#model_safety CesGuardrail#model_safety}
	ModelSafety *CesGuardrailModelSafety `field:"optional" json:"modelSafety" yaml:"modelSafety"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#project CesGuardrail#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#timeouts CesGuardrail#timeouts}
	Timeouts *CesGuardrailTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

