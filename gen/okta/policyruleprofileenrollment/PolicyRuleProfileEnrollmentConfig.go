package policyruleprofileenrollment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyRuleProfileEnrollmentConfig struct {
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
	// ID of the policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#policy_id PolicyRuleProfileEnrollment#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Which action should be taken if this User is new.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#unknown_user_action PolicyRuleProfileEnrollment#unknown_user_action}
	UnknownUserAction *string `field:"required" json:"unknownUserAction" yaml:"unknownUserAction"`
	// Allow or deny access based on the rule conditions: ALLOW or DENY.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#access PolicyRuleProfileEnrollment#access}
	Access *string `field:"optional" json:"access" yaml:"access"`
	// Indicates whether email verification should occur before access is granted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#email_verification PolicyRuleProfileEnrollment#email_verification}
	EmailVerification interface{} `field:"optional" json:"emailVerification" yaml:"emailVerification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#id PolicyRuleProfileEnrollment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ID of a Registration Inline Hook.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#inline_hook_id PolicyRuleProfileEnrollment#inline_hook_id}
	InlineHookId *string `field:"optional" json:"inlineHookId" yaml:"inlineHookId"`
	// profile_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#profile_attributes PolicyRuleProfileEnrollment#profile_attributes}
	ProfileAttributes interface{} `field:"optional" json:"profileAttributes" yaml:"profileAttributes"`
	// The ID of a Group that this User should be added to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#target_group_id PolicyRuleProfileEnrollment#target_group_id}
	TargetGroupId *string `field:"optional" json:"targetGroupId" yaml:"targetGroupId"`
	// Value created by the backend. If present all policy updates must include this attribute/value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#ui_schema_id PolicyRuleProfileEnrollment#ui_schema_id}
	UiSchemaId *string `field:"optional" json:"uiSchemaId" yaml:"uiSchemaId"`
}

