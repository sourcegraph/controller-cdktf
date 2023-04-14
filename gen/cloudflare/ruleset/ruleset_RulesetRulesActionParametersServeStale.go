package ruleset


type RulesetRulesActionParametersServeStale struct {
	// Disable stale while updating.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#disable_stale_while_updating Ruleset#disable_stale_while_updating}
	DisableStaleWhileUpdating interface{} `field:"optional" json:"disableStaleWhileUpdating" yaml:"disableStaleWhileUpdating"`
}

