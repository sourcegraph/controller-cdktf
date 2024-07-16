package googlebillingbudget


type GoogleBillingBudgetBudgetFilterCustomPeriod struct {
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_budget#start_date GoogleBillingBudget#start_date}
	StartDate *GoogleBillingBudgetBudgetFilterCustomPeriodStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_budget#end_date GoogleBillingBudget#end_date}
	EndDate *GoogleBillingBudgetBudgetFilterCustomPeriodEndDate `field:"optional" json:"endDate" yaml:"endDate"`
}

