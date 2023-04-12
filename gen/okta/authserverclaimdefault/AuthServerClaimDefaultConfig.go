package authserverclaimdefault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthServerClaimDefaultConfig struct {
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
	// Auth server ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim_default#auth_server_id AuthServerClaimDefault#auth_server_id}
	AuthServerId *string `field:"required" json:"authServerId" yaml:"authServerId"`
	// Default auth server claim name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim_default#name AuthServerClaimDefault#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim_default#always_include_in_token AuthServerClaimDefault#always_include_in_token}.
	AlwaysIncludeInToken interface{} `field:"optional" json:"alwaysIncludeInToken" yaml:"alwaysIncludeInToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim_default#id AuthServerClaimDefault#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim_default#value AuthServerClaimDefault#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

