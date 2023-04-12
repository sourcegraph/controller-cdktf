package datatfegithubappinstallation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTfeGithubAppInstallationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/github_app_installation#installation_id DataTfeGithubAppInstallation#installation_id}.
	InstallationId *float64 `field:"optional" json:"installationId" yaml:"installationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/github_app_installation#name DataTfeGithubAppInstallation#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

