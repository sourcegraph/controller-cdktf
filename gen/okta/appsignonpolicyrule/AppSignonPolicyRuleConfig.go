package appsignonpolicyrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSignonPolicyRuleConfig struct {
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
	// Policy Rule Name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#name AppSignonPolicyRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of the policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#policy_id AppSignonPolicyRule#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Allow or deny access based on the rule conditions: ALLOW or DENY.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#access AppSignonPolicyRule#access}
	Access *string `field:"optional" json:"access" yaml:"access"`
	// An array that contains nested Authenticator Constraint objects that are organized by the Authenticator class.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#constraints AppSignonPolicyRule#constraints}
	Constraints *[]*string `field:"optional" json:"constraints" yaml:"constraints"`
	// This is an optional advanced setting.
	//
	// If the expression is formatted incorrectly or conflicts with conditions set above, the rule may not match any users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#custom_expression AppSignonPolicyRule#custom_expression}
	CustomExpression *string `field:"optional" json:"customExpression" yaml:"customExpression"`
	// If the device is managed.
	//
	// A device is managed if it's managed by a device management system. When managed is passed, registered must also be included and must be set to true.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#device_is_managed AppSignonPolicyRule#device_is_managed}
	DeviceIsManaged interface{} `field:"optional" json:"deviceIsManaged" yaml:"deviceIsManaged"`
	// If the device is registered.
	//
	// A device is registered if the User enrolls with Okta Verify that is installed on the device.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#device_is_registered AppSignonPolicyRule#device_is_registered}
	DeviceIsRegistered interface{} `field:"optional" json:"deviceIsRegistered" yaml:"deviceIsRegistered"`
	// The number of factors required to satisfy this assurance level.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#factor_mode AppSignonPolicyRule#factor_mode}
	FactorMode *string `field:"optional" json:"factorMode" yaml:"factorMode"`
	// List of group IDs to exclude.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#groups_excluded AppSignonPolicyRule#groups_excluded}
	GroupsExcluded *[]*string `field:"optional" json:"groupsExcluded" yaml:"groupsExcluded"`
	// List of group IDs to include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#groups_included AppSignonPolicyRule#groups_included}
	GroupsIncluded *[]*string `field:"optional" json:"groupsIncluded" yaml:"groupsIncluded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#id AppSignonPolicyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The inactivity duration after which the end user must re-authenticate.
	//
	// Use the ISO 8601 Period format for recurring time intervals.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#inactivity_period AppSignonPolicyRule#inactivity_period}
	InactivityPeriod *string `field:"optional" json:"inactivityPeriod" yaml:"inactivityPeriod"`
	// Network selection mode: ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#network_connection AppSignonPolicyRule#network_connection}
	NetworkConnection *string `field:"optional" json:"networkConnection" yaml:"networkConnection"`
	// The zones to exclude.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#network_excludes AppSignonPolicyRule#network_excludes}
	NetworkExcludes *[]*string `field:"optional" json:"networkExcludes" yaml:"networkExcludes"`
	// The zones to include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#network_includes AppSignonPolicyRule#network_includes}
	NetworkIncludes *[]*string `field:"optional" json:"networkIncludes" yaml:"networkIncludes"`
	// platform_include block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#platform_include AppSignonPolicyRule#platform_include}
	PlatformInclude interface{} `field:"optional" json:"platformInclude" yaml:"platformInclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#priority AppSignonPolicyRule#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The duration after which the end user must re-authenticate, regardless of user activity.
	//
	// Use the ISO 8601 Period format for recurring time intervals. PT0S - Every sign-in attempt, PT43800H - Once per session
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#re_authentication_frequency AppSignonPolicyRule#re_authentication_frequency}
	ReAuthenticationFrequency *string `field:"optional" json:"reAuthenticationFrequency" yaml:"reAuthenticationFrequency"`
	// Status of the rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#status AppSignonPolicyRule#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The Verification Method type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#type AppSignonPolicyRule#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Set of User IDs to exclude.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#users_excluded AppSignonPolicyRule#users_excluded}
	UsersExcluded *[]*string `field:"optional" json:"usersExcluded" yaml:"usersExcluded"`
	// Set of User IDs to include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#users_included AppSignonPolicyRule#users_included}
	UsersIncluded *[]*string `field:"optional" json:"usersIncluded" yaml:"usersIncluded"`
	// Set of User Type IDs to exclude.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#user_types_excluded AppSignonPolicyRule#user_types_excluded}
	UserTypesExcluded *[]*string `field:"optional" json:"userTypesExcluded" yaml:"userTypesExcluded"`
	// Set of User Type IDs to include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule#user_types_included AppSignonPolicyRule#user_types_included}
	UserTypesIncluded *[]*string `field:"optional" json:"userTypesIncluded" yaml:"userTypesIncluded"`
}

