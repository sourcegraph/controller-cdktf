package registrymodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegistryModuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#id RegistryModule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#module_provider RegistryModule#module_provider}.
	ModuleProvider *string `field:"optional" json:"moduleProvider" yaml:"moduleProvider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#name RegistryModule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#namespace RegistryModule#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#no_code RegistryModule#no_code}.
	NoCode interface{} `field:"optional" json:"noCode" yaml:"noCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#organization RegistryModule#organization}.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#registry_name RegistryModule#registry_name}.
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// vcs_repo block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#vcs_repo RegistryModule#vcs_repo}
	VcsRepo *RegistryModuleVcsRepo `field:"optional" json:"vcsRepo" yaml:"vcsRepo"`
}

