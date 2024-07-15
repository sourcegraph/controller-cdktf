package googlebillingbudget


type GoogleBillingBudgetBudgetFilterCustomPeriodStartDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_budget#day GoogleBillingBudget#day}
	Day *float64 `field:"required" json:"day" yaml:"day"`
	// Month of a year. Must be from 1 to 12.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_budget#month GoogleBillingBudget#month}
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// Year of the date. Must be from 1 to 9999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_budget#year GoogleBillingBudget#year}
	Year *float64 `field:"required" json:"year" yaml:"year"`
}

