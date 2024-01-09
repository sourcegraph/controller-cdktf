package alertpolicy


type AlertPolicyCondition struct {
	// One of `timeToBurnBudget` | `timeToBurnEntireBudget` | `burnRate` | `burnedBudget`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/alert_policy#measurement AlertPolicy#measurement}
	Measurement *string `field:"required" json:"measurement" yaml:"measurement"`
	// Indicates how long a given condition needs to be valid to mark the condition as true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/alert_policy#lasts_for AlertPolicy#lasts_for}
	LastsFor *string `field:"optional" json:"lastsFor" yaml:"lastsFor"`
	// For `averageBurnRate`, it indicates how fast the error budget is burning.
	//
	// For `burnedBudget`, it tells how much error budget is already burned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/alert_policy#value AlertPolicy#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
	// Used with `timeToBurnBudget` or `timeToBurnEntireBudget`, indicates when the budget would be exhausted.
	//
	// The expected value is a string in time duration string format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/alert_policy#value_string AlertPolicy#value_string}
	ValueString *string `field:"optional" json:"valueString" yaml:"valueString"`
}

