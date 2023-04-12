package trustedorigin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TrustedOriginConfig struct {
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
	// Unique name for this trusted origin.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/trusted_origin#name TrustedOrigin#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique origin URL for this trusted origin.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/trusted_origin#origin TrustedOrigin#origin}
	Origin *string `field:"required" json:"origin" yaml:"origin"`
	// Scopes of the Trusted Origin - can either be CORS or REDIRECT only.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/trusted_origin#scopes TrustedOrigin#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Whether the Trusted Origin is active or not - can only be issued post-creation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/trusted_origin#active TrustedOrigin#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/trusted_origin#id TrustedOrigin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

