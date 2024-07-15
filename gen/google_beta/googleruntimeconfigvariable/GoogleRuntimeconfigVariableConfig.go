package googleruntimeconfigvariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleRuntimeconfigVariableConfig struct {
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
	// The name of the variable to manage. Note that variable names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_runtimeconfig_variable#name GoogleRuntimeconfigVariable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the RuntimeConfig resource containing this variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_runtimeconfig_variable#parent GoogleRuntimeconfigVariable#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_runtimeconfig_variable#id GoogleRuntimeconfigVariable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_runtimeconfig_variable#project GoogleRuntimeconfigVariable#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_runtimeconfig_variable#text GoogleRuntimeconfigVariable#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_runtimeconfig_variable#value GoogleRuntimeconfigVariable#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

