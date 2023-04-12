package resourceset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourceSetConfig struct {
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
	// A description of the Resource Set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/resource_set#description ResourceSet#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Unique name given to the Resource Set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/resource_set#label ResourceSet#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/resource_set#id ResourceSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The endpoints that reference the resources to be included in the new Resource Set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/resource_set#resources ResourceSet#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

