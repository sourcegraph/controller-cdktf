package accesspolicy


type AccessPolicyExcludeExternalEvaluation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#evaluate_url AccessPolicy#evaluate_url}.
	EvaluateUrl *string `field:"optional" json:"evaluateUrl" yaml:"evaluateUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#keys_url AccessPolicy#keys_url}.
	KeysUrl *string `field:"optional" json:"keysUrl" yaml:"keysUrl"`
}

