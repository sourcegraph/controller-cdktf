package ruleset


type RulesetRulesActionParametersMatchedData struct {
	// Public key to use within WAF Ruleset payload logging to view the HTTP request parameters.
	//
	// You can generate a public key [using the `matched-data-cli` command-line tool](https://developers.cloudflare.com/waf/managed-rulesets/payload-logging/command-line/generate-key-pair) or [in the Cloudflare dashboard](https://developers.cloudflare.com/waf/managed-rulesets/payload-logging/configure).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#public_key Ruleset#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
}

