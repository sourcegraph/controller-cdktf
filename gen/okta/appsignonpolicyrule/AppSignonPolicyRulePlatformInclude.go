package appsignonpolicyrule


type AppSignonPolicyRulePlatformInclude struct {
	// Only available with OTHER OS type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_signon_policy_rule#os_expression AppSignonPolicyRule#os_expression}
	OsExpression *string `field:"optional" json:"osExpression" yaml:"osExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_signon_policy_rule#os_type AppSignonPolicyRule#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_signon_policy_rule#type AppSignonPolicyRule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

