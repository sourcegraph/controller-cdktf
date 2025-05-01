package googlegeminirepositorygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGeminiRepositoryGroupConfig struct {
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
	// Required. Id of the Code Repository Index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#code_repository_index GoogleGeminiRepositoryGroup#code_repository_index}
	CodeRepositoryIndex *string `field:"required" json:"codeRepositoryIndex" yaml:"codeRepositoryIndex"`
	// The location of the Code Repository Index, for example 'us-central1'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#location GoogleGeminiRepositoryGroup#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// repositories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#repositories GoogleGeminiRepositoryGroup#repositories}
	Repositories interface{} `field:"required" json:"repositories" yaml:"repositories"`
	// Required. Id of the Repository Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#repository_group_id GoogleGeminiRepositoryGroup#repository_group_id}
	RepositoryGroupId *string `field:"required" json:"repositoryGroupId" yaml:"repositoryGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#id GoogleGeminiRepositoryGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#labels GoogleGeminiRepositoryGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#project GoogleGeminiRepositoryGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#timeouts GoogleGeminiRepositoryGroup#timeouts}
	Timeouts *GoogleGeminiRepositoryGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

