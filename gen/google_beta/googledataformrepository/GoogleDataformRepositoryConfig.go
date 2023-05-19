package googledataformrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataformRepositoryConfig struct {
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
	// The repository's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataform_repository#name GoogleDataformRepository#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// git_remote_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataform_repository#git_remote_settings GoogleDataformRepository#git_remote_settings}
	GitRemoteSettings *GoogleDataformRepositoryGitRemoteSettings `field:"optional" json:"gitRemoteSettings" yaml:"gitRemoteSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataform_repository#id GoogleDataformRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataform_repository#project GoogleDataformRepository#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A reference to the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataform_repository#region GoogleDataformRepository#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataform_repository#timeouts GoogleDataformRepository#timeouts}
	Timeouts *GoogleDataformRepositoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

