package googlebillingbudget


type GoogleBillingBudgetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_billing_budget#create GoogleBillingBudget#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_billing_budget#delete GoogleBillingBudget#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_billing_budget#update GoogleBillingBudget#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

