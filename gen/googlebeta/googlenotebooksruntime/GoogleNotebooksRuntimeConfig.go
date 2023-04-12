package googlenotebooksruntime

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNotebooksRuntimeConfig struct {
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
	// A reference to the zone where the machine resides.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#location GoogleNotebooksRuntime#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name specified for the Notebook runtime.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#name GoogleNotebooksRuntime#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#access_config GoogleNotebooksRuntime#access_config}
	AccessConfig *GoogleNotebooksRuntimeAccessConfig `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#id GoogleNotebooksRuntime#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#project GoogleNotebooksRuntime#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// software_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#software_config GoogleNotebooksRuntime#software_config}
	SoftwareConfig *GoogleNotebooksRuntimeSoftwareConfig `field:"optional" json:"softwareConfig" yaml:"softwareConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#timeouts GoogleNotebooksRuntime#timeouts}
	Timeouts *GoogleNotebooksRuntimeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#virtual_machine GoogleNotebooksRuntime#virtual_machine}
	VirtualMachine *GoogleNotebooksRuntimeVirtualMachine `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
}

