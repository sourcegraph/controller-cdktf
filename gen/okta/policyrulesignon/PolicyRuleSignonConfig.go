package policyrulesignon

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyRuleSignonConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#name PolicyRuleSignon#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Allow or deny access based on the rule conditions: ALLOW, DENY or CHALLENGE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#access PolicyRuleSignon#access}
	Access *string `field:"optional" json:"access" yaml:"access"`
	// Authentication entrypoint: ANY, RADIUS or LDAP_INTERFACE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#authtype PolicyRuleSignon#authtype}
	Authtype *string `field:"optional" json:"authtype" yaml:"authtype"`
	// List of behavior IDs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#behaviors PolicyRuleSignon#behaviors}
	Behaviors *[]*string `field:"optional" json:"behaviors" yaml:"behaviors"`
	// factor_sequence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#factor_sequence PolicyRuleSignon#factor_sequence}
	FactorSequence interface{} `field:"optional" json:"factorSequence" yaml:"factorSequence"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#id PolicyRuleSignon#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Apply rule based on the IdP used: ANY, OKTA or SPECIFIC_IDP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#identity_provider PolicyRuleSignon#identity_provider}
	IdentityProvider *string `field:"optional" json:"identityProvider" yaml:"identityProvider"`
	// When identity_provider is SPECIFIC_IDP then this is the list of IdP IDs to apply the rule on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#identity_provider_ids PolicyRuleSignon#identity_provider_ids}
	IdentityProviderIds *[]*string `field:"optional" json:"identityProviderIds" yaml:"identityProviderIds"`
	// Elapsed time before the next MFA challenge.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#mfa_lifetime PolicyRuleSignon#mfa_lifetime}
	MfaLifetime *float64 `field:"optional" json:"mfaLifetime" yaml:"mfaLifetime"`
	// Prompt for MFA based on the device used, a factor session lifetime, or every sign-on attempt: DEVICE, SESSION or ALWAYS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#mfa_prompt PolicyRuleSignon#mfa_prompt}
	MfaPrompt *string `field:"optional" json:"mfaPrompt" yaml:"mfaPrompt"`
	// Remember MFA device.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#mfa_remember_device PolicyRuleSignon#mfa_remember_device}
	MfaRememberDevice interface{} `field:"optional" json:"mfaRememberDevice" yaml:"mfaRememberDevice"`
	// Require MFA.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#mfa_required PolicyRuleSignon#mfa_required}
	MfaRequired interface{} `field:"optional" json:"mfaRequired" yaml:"mfaRequired"`
	// Network selection mode: ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#network_connection PolicyRuleSignon#network_connection}
	NetworkConnection *string `field:"optional" json:"networkConnection" yaml:"networkConnection"`
	// The zones to exclude.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#network_excludes PolicyRuleSignon#network_excludes}
	NetworkExcludes *[]*string `field:"optional" json:"networkExcludes" yaml:"networkExcludes"`
	// The zones to include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#network_includes PolicyRuleSignon#network_includes}
	NetworkIncludes *[]*string `field:"optional" json:"networkIncludes" yaml:"networkIncludes"`
	// Policy ID of the Rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#policyid PolicyRuleSignon#policyid}
	Policyid *string `field:"optional" json:"policyid" yaml:"policyid"`
	// Policy ID of the Rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#policy_id PolicyRuleSignon#policy_id}
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Primary factor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#primary_factor PolicyRuleSignon#primary_factor}
	PrimaryFactor *string `field:"optional" json:"primaryFactor" yaml:"primaryFactor"`
	// Policy Rule Priority, this attribute can be set to a valid priority.
	//
	// To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last (lowest) if not there.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#priority PolicyRuleSignon#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Risc level: ANY, LOW, MEDIUM or HIGH.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#risc_level PolicyRuleSignon#risc_level}
	RiscLevel *string `field:"optional" json:"riscLevel" yaml:"riscLevel"`
	// Max minutes a session can be idle.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#session_idle PolicyRuleSignon#session_idle}
	SessionIdle *float64 `field:"optional" json:"sessionIdle" yaml:"sessionIdle"`
	// Max minutes a session is active: Disable = 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#session_lifetime PolicyRuleSignon#session_lifetime}
	SessionLifetime *float64 `field:"optional" json:"sessionLifetime" yaml:"sessionLifetime"`
	// Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session cookies.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#session_persistent PolicyRuleSignon#session_persistent}
	SessionPersistent interface{} `field:"optional" json:"sessionPersistent" yaml:"sessionPersistent"`
	// Policy Rule Status: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#status PolicyRuleSignon#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Set of User IDs to Exclude.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon#users_excluded PolicyRuleSignon#users_excluded}
	UsersExcluded *[]*string `field:"optional" json:"usersExcluded" yaml:"usersExcluded"`
}

