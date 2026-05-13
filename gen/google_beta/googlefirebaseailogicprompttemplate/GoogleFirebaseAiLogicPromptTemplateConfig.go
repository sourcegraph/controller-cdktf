package googlefirebaseailogicprompttemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseAiLogicPromptTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_prompt_template#location GoogleFirebaseAiLogicPromptTemplate#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The unique ID of the PromptTemplate, which is the final component of the PromptTemplate's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_prompt_template#template_id GoogleFirebaseAiLogicPromptTemplate#template_id}
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// The [Dotprompt](https://google.github.io/dotprompt/getting-started) raw template string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_prompt_template#template_string GoogleFirebaseAiLogicPromptTemplate#template_string}
	TemplateString *string `field:"required" json:"templateString" yaml:"templateString"`
	// The display name of the PromptTemplate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_prompt_template#display_name GoogleFirebaseAiLogicPromptTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_prompt_template#id GoogleFirebaseAiLogicPromptTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_prompt_template#project GoogleFirebaseAiLogicPromptTemplate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_prompt_template#timeouts GoogleFirebaseAiLogicPromptTemplate#timeouts}
	Timeouts *GoogleFirebaseAiLogicPromptTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

