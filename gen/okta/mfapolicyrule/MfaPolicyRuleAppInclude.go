package mfapolicyrule


type MfaPolicyRuleAppInclude struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy_rule#type MfaPolicyRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy_rule#id MfaPolicyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy_rule#name MfaPolicyRule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

