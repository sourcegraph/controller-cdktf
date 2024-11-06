package adminorganizationsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdminOrganizationSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/admin_organization_settings#access_beta_tools AdminOrganizationSettings#access_beta_tools}.
	AccessBetaTools interface{} `field:"optional" json:"accessBetaTools" yaml:"accessBetaTools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/admin_organization_settings#global_module_sharing AdminOrganizationSettings#global_module_sharing}.
	GlobalModuleSharing interface{} `field:"optional" json:"globalModuleSharing" yaml:"globalModuleSharing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/admin_organization_settings#id AdminOrganizationSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/admin_organization_settings#module_sharing_consumer_organizations AdminOrganizationSettings#module_sharing_consumer_organizations}.
	ModuleSharingConsumerOrganizations *[]*string `field:"optional" json:"moduleSharingConsumerOrganizations" yaml:"moduleSharingConsumerOrganizations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/admin_organization_settings#organization AdminOrganizationSettings#organization}.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/admin_organization_settings#workspace_limit AdminOrganizationSettings#workspace_limit}.
	WorkspaceLimit *float64 `field:"optional" json:"workspaceLimit" yaml:"workspaceLimit"`
}

