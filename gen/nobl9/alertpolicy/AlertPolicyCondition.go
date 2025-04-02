package alertpolicy


type AlertPolicyCondition struct {
	// One of `timeToBurnBudget` | `timeToBurnEntireBudget` | `averageBurnRate` | `burnedBudget` | `budgetDrop`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/alert_policy#measurement AlertPolicy#measurement}
	Measurement *string `field:"required" json:"measurement" yaml:"measurement"`
	// Duration over which the burn rate is evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/alert_policy#alerting_window AlertPolicy#alerting_window}
	AlertingWindow *string `field:"optional" json:"alertingWindow" yaml:"alertingWindow"`
	// Indicates how long a given condition needs to be valid to mark the condition as true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/alert_policy#lasts_for AlertPolicy#lasts_for}
	LastsFor *string `field:"optional" json:"lastsFor" yaml:"lastsFor"`
	// A mathematical inequality operator. One of `lt` | `lte` | `gt` | `gte`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/alert_policy#op AlertPolicy#op}
	Op *string `field:"optional" json:"op" yaml:"op"`
	// For `averageBurnRate`, it indicates how fast the error budget is burning.
	//
	// For `burnedBudget`, it tells how much error budget is already burned. For `budgetDrop`, it tells how much budget dropped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/alert_policy#value AlertPolicy#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
	// Used with `timeToBurnBudget` or `timeToBurnEntireBudget`, indicates when the budget would be exhausted.
	//
	// The expected value is a string in time duration string format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/alert_policy#value_string AlertPolicy#value_string}
	ValueString *string `field:"optional" json:"valueString" yaml:"valueString"`
}

