package googleidentityplatformprojectdefaultconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIdentityPlatformProjectDefaultConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_project_default_config#id GoogleIdentityPlatformProjectDefaultConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_project_default_config#project GoogleIdentityPlatformProjectDefaultConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// sign_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_project_default_config#sign_in GoogleIdentityPlatformProjectDefaultConfig#sign_in}
	SignIn *GoogleIdentityPlatformProjectDefaultConfigSignIn `field:"optional" json:"signIn" yaml:"signIn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_project_default_config#timeouts GoogleIdentityPlatformProjectDefaultConfig#timeouts}
	Timeouts *GoogleIdentityPlatformProjectDefaultConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

