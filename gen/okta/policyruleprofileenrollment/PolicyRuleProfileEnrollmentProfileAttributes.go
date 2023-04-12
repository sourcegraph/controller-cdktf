package policyruleprofileenrollment


type PolicyRuleProfileEnrollmentProfileAttributes struct {
	// A display-friendly label for this property.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#label PolicyRuleProfileEnrollment#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// The name of a User Profile property.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#name PolicyRuleProfileEnrollment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates if this property is required for enrollment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment#required PolicyRuleProfileEnrollment#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

