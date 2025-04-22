package googlesiteverificationowner

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSiteVerificationOwnerConfig struct {
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
	// The email address of the owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_site_verification_owner#email GoogleSiteVerificationOwner#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// The id of the Web Resource to add this owner to, in the form "webResource/<web-resource-id>".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_site_verification_owner#web_resource_id GoogleSiteVerificationOwner#web_resource_id}
	WebResourceId *string `field:"required" json:"webResourceId" yaml:"webResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_site_verification_owner#id GoogleSiteVerificationOwner#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_site_verification_owner#timeouts GoogleSiteVerificationOwner#timeouts}
	Timeouts *GoogleSiteVerificationOwnerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

