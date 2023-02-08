package passwordpolicyrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PasswordPolicyRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#name PasswordPolicyRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#id PasswordPolicyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Network selection mode: ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#network_connection PasswordPolicyRule#network_connection}
	NetworkConnection *string `field:"optional" json:"networkConnection" yaml:"networkConnection"`
	// The zones to exclude.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#network_excludes PasswordPolicyRule#network_excludes}
	NetworkExcludes *[]*string `field:"optional" json:"networkExcludes" yaml:"networkExcludes"`
	// The zones to include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#network_includes PasswordPolicyRule#network_includes}
	NetworkIncludes *[]*string `field:"optional" json:"networkIncludes" yaml:"networkIncludes"`
	// Allow or deny a user to change their password: ALLOW or DENY. Default = ALLOW.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#password_change PasswordPolicyRule#password_change}
	PasswordChange *string `field:"optional" json:"passwordChange" yaml:"passwordChange"`
	// Allow or deny a user to reset their password: ALLOW or DENY. Default = ALLOW.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#password_reset PasswordPolicyRule#password_reset}
	PasswordReset *string `field:"optional" json:"passwordReset" yaml:"passwordReset"`
	// Allow or deny a user to unlock. Default = DENY.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#password_unlock PasswordPolicyRule#password_unlock}
	PasswordUnlock *string `field:"optional" json:"passwordUnlock" yaml:"passwordUnlock"`
	// Policy ID of the Rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#policyid PasswordPolicyRule#policyid}
	Policyid *string `field:"optional" json:"policyid" yaml:"policyid"`
	// Policy ID of the Rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#policy_id PasswordPolicyRule#policy_id}
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Policy Rule Priority, this attribute can be set to a valid priority.
	//
	// To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last (lowest) if not there.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#priority PasswordPolicyRule#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Policy Rule Status: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#status PasswordPolicyRule#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Set of User IDs to Exclude.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy_rule#users_excluded PasswordPolicyRule#users_excluded}
	UsersExcluded *[]*string `field:"optional" json:"usersExcluded" yaml:"usersExcluded"`
}

