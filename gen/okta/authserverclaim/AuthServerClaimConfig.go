package authserverclaim

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthServerClaimConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#auth_server_id AuthServerClaim#auth_server_id}
	AuthServerId *string `field:"required" json:"authServerId" yaml:"authServerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#claim_type AuthServerClaim#claim_type}.
	ClaimType *string `field:"required" json:"claimType" yaml:"claimType"`
	// Auth server claim name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#name AuthServerClaim#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#value AuthServerClaim#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#always_include_in_token AuthServerClaim#always_include_in_token}.
	AlwaysIncludeInToken interface{} `field:"optional" json:"alwaysIncludeInToken" yaml:"alwaysIncludeInToken"`
	// Required when value_type is GROUPS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#group_filter_type AuthServerClaim#group_filter_type}
	GroupFilterType *string `field:"optional" json:"groupFilterType" yaml:"groupFilterType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#id AuthServerClaim#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Auth server claim list of scopes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#scopes AuthServerClaim#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#status AuthServerClaim#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_claim#value_type AuthServerClaim#value_type}.
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

