package authserverpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthServerPolicyConfig struct {
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
	// The ID of the Auth Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy#auth_server_id AuthServerPolicy#auth_server_id}
	AuthServerId *string `field:"required" json:"authServerId" yaml:"authServerId"`
	// The clients to whitelist the policy for.
	//
	// `[ALL_CLIENTS]` is a special value that can be used to whitelist all clients, otherwise it is a list of client ids.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy#client_whitelist AuthServerPolicy#client_whitelist}
	ClientWhitelist *[]*string `field:"required" json:"clientWhitelist" yaml:"clientWhitelist"`
	// The description of the Auth Server Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy#description AuthServerPolicy#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the Auth Server Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy#name AuthServerPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Priority of the auth server policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy#priority AuthServerPolicy#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy#id AuthServerPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Default to `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy#status AuthServerPolicy#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

