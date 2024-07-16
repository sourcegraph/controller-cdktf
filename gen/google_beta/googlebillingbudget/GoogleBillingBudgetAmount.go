package googlebillingbudget


type GoogleBillingBudgetAmount struct {
	// Configures a budget amount that is automatically set to 100% of last period's spend.
	//
	// Boolean. Set value to true to use. Do not set to false, instead
	// use the 'specified_amount' block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_budget#last_period_amount GoogleBillingBudget#last_period_amount}
	LastPeriodAmount interface{} `field:"optional" json:"lastPeriodAmount" yaml:"lastPeriodAmount"`
	// specified_amount block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_budget#specified_amount GoogleBillingBudget#specified_amount}
	SpecifiedAmount *GoogleBillingBudgetAmountSpecifiedAmount `field:"optional" json:"specifiedAmount" yaml:"specifiedAmount"`
}

