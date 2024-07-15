package googlefirebaserulesruleset


type GoogleFirebaserulesRulesetSourceFiles struct {
	// Textual Content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebaserules_ruleset#content GoogleFirebaserulesRuleset#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// File name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebaserules_ruleset#name GoogleFirebaserulesRuleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Fingerprint (e.g. github sha) associated with the `File`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebaserules_ruleset#fingerprint GoogleFirebaserulesRuleset#fingerprint}
	Fingerprint *string `field:"optional" json:"fingerprint" yaml:"fingerprint"`
}

