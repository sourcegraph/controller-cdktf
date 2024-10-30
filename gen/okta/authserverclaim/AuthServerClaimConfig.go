package authserverclaim

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthServerClaimConfig struct {
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
	// ID of the authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#auth_server_id AuthServerClaim#auth_server_id}
	AuthServerId *string `field:"required" json:"authServerId" yaml:"authServerId"`
	// Specifies whether the claim is for an access token `RESOURCE` or ID token `IDENTITY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#claim_type AuthServerClaim#claim_type}
	ClaimType *string `field:"required" json:"claimType" yaml:"claimType"`
	// The name of the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#name AuthServerClaim#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#value AuthServerClaim#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Specifies whether to include claims in token, by default it is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#always_include_in_token AuthServerClaim#always_include_in_token}
	AlwaysIncludeInToken interface{} `field:"optional" json:"alwaysIncludeInToken" yaml:"alwaysIncludeInToken"`
	// Specifies the type of group filter if `value_type` is `GROUPS`.
	//
	// Can be set to one of the following `STARTS_WITH`, `EQUALS`, `CONTAINS`, `REGEX`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#group_filter_type AuthServerClaim#group_filter_type}
	GroupFilterType *string `field:"optional" json:"groupFilterType" yaml:"groupFilterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#id AuthServerClaim#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The list of scopes the auth server claim is tied to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#scopes AuthServerClaim#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Default to `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#status AuthServerClaim#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The type of value of the claim. It can be set to `EXPRESSION` or `GROUPS`. It defaults to `EXPRESSION`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_claim#value_type AuthServerClaim#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

