package resourceexplorer2view

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Resourceexplorer2ViewConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/resourceexplorer2_view#name Resourceexplorer2View#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/resourceexplorer2_view#default_view Resourceexplorer2View#default_view}.
	DefaultView any `field:"optional" json:"defaultView" yaml:"defaultView"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/resourceexplorer2_view#filters Resourceexplorer2View#filters}
	Filters any `field:"optional" json:"filters" yaml:"filters"`
	// included_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/resourceexplorer2_view#included_property Resourceexplorer2View#included_property}
	IncludedProperty any `field:"optional" json:"includedProperty" yaml:"includedProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/resourceexplorer2_view#tags Resourceexplorer2View#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}
