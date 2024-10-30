package authserverpolicyrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthServerPolicyRuleConfig struct {
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
	// Auth server ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#auth_server_id AuthServerPolicyRule#auth_server_id}
	AuthServerId *string `field:"required" json:"authServerId" yaml:"authServerId"`
	// Accepted grant type values, `authorization_code`, `implicit`, `password`, `client_credentials`, `urn:ietf:params:oauth:grant-type:saml2-bearer` (*Early Access Property*), `urn:ietf:params:oauth:grant-type:token-exchange` (*Early Access Property*),`urn:ietf:params:oauth:grant-type:device_code` (*Early Access Property*), `interaction_code` (*OIE only*).
	//
	// For `implicit` value either `user_whitelist` or `group_whitelist` should be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#grant_type_whitelist AuthServerPolicyRule#grant_type_whitelist}
	GrantTypeWhitelist *[]*string `field:"required" json:"grantTypeWhitelist" yaml:"grantTypeWhitelist"`
	// Auth server policy rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#name AuthServerPolicyRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Auth server policy ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#policy_id AuthServerPolicyRule#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Priority of the auth server policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#priority AuthServerPolicyRule#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Lifetime of access token. Can be set to a value between 5 and 1440 minutes. Default is `60`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#access_token_lifetime_minutes AuthServerPolicyRule#access_token_lifetime_minutes}
	AccessTokenLifetimeMinutes *float64 `field:"optional" json:"accessTokenLifetimeMinutes" yaml:"accessTokenLifetimeMinutes"`
	// Specifies a set of Groups whose Users are to be excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#group_blacklist AuthServerPolicyRule#group_blacklist}
	GroupBlacklist *[]*string `field:"optional" json:"groupBlacklist" yaml:"groupBlacklist"`
	// Specifies a set of Groups whose Users are to be included.
	//
	// Can be set to Group ID or to the following: `EVERYONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#group_whitelist AuthServerPolicyRule#group_whitelist}
	GroupWhitelist *[]*string `field:"optional" json:"groupWhitelist" yaml:"groupWhitelist"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#id AuthServerPolicyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the inline token to trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#inline_hook_id AuthServerPolicyRule#inline_hook_id}
	InlineHookId *string `field:"optional" json:"inlineHookId" yaml:"inlineHookId"`
	// Lifetime of refresh token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#refresh_token_lifetime_minutes AuthServerPolicyRule#refresh_token_lifetime_minutes}
	RefreshTokenLifetimeMinutes *float64 `field:"optional" json:"refreshTokenLifetimeMinutes" yaml:"refreshTokenLifetimeMinutes"`
	// Window in which a refresh token can be used.
	//
	// It can be a value between 5 and 2628000 (5 years) minutes. Default is `10080` (7 days).`refresh_token_window_minutes` must be between `access_token_lifetime_minutes` and `refresh_token_lifetime_minutes`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#refresh_token_window_minutes AuthServerPolicyRule#refresh_token_window_minutes}
	RefreshTokenWindowMinutes *float64 `field:"optional" json:"refreshTokenWindowMinutes" yaml:"refreshTokenWindowMinutes"`
	// Scopes allowed for this policy rule.
	//
	// They can be whitelisted by name or all can be whitelisted with ` * `
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#scope_whitelist AuthServerPolicyRule#scope_whitelist}
	ScopeWhitelist *[]*string `field:"optional" json:"scopeWhitelist" yaml:"scopeWhitelist"`
	// Default to `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#status AuthServerPolicyRule#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Auth server policy rule type, unlikely this will be anything other then the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#type AuthServerPolicyRule#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies a set of Users to be excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#user_blacklist AuthServerPolicyRule#user_blacklist}
	UserBlacklist *[]*string `field:"optional" json:"userBlacklist" yaml:"userBlacklist"`
	// Specifies a set of Users to be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/auth_server_policy_rule#user_whitelist AuthServerPolicyRule#user_whitelist}
	UserWhitelist *[]*string `field:"optional" json:"userWhitelist" yaml:"userWhitelist"`
}

