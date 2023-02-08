package authserverpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthServerPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_policy#auth_server_id AuthServerPolicy#auth_server_id}.
	AuthServerId *string `field:"required" json:"authServerId" yaml:"authServerId"`
	// Use ["ALL_CLIENTS"] when unsure.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_policy#client_whitelist AuthServerPolicy#client_whitelist}
	ClientWhitelist *[]*string `field:"required" json:"clientWhitelist" yaml:"clientWhitelist"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_policy#description AuthServerPolicy#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_policy#name AuthServerPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Priority of the auth server policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_policy#priority AuthServerPolicy#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_policy#id AuthServerPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_policy#status AuthServerPolicy#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Auth server policy type, unlikely this will be anything other then the default.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_policy#type AuthServerPolicy#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

