package datagoogleruntimeconfigconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleRuntimeconfigConfigConfig struct {
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
	// The name of the runtime config.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_runtimeconfig_config#name DataGoogleRuntimeconfigConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_runtimeconfig_config#id DataGoogleRuntimeconfigConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_runtimeconfig_config#project DataGoogleRuntimeconfigConfig#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}
