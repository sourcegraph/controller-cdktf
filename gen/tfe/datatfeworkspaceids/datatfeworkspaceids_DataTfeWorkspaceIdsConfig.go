package datatfeworkspaceids

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTfeWorkspaceIdsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/workspace_ids#organization DataTfeWorkspaceIds#organization}.
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/workspace_ids#exclude_tags DataTfeWorkspaceIds#exclude_tags}.
	ExcludeTags *[]*string `field:"optional" json:"excludeTags" yaml:"excludeTags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/workspace_ids#id DataTfeWorkspaceIds#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/workspace_ids#names DataTfeWorkspaceIds#names}.
	Names *[]*string `field:"optional" json:"names" yaml:"names"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/workspace_ids#tag_names DataTfeWorkspaceIds#tag_names}.
	TagNames *[]*string `field:"optional" json:"tagNames" yaml:"tagNames"`
}

