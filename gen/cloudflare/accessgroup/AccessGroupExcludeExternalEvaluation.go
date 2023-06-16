package accessgroup


type AccessGroupExcludeExternalEvaluation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#evaluate_url AccessGroup#evaluate_url}.
	EvaluateUrl *string `field:"optional" json:"evaluateUrl" yaml:"evaluateUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#keys_url AccessGroup#keys_url}.
	KeysUrl *string `field:"optional" json:"keysUrl" yaml:"keysUrl"`
}

