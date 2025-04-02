package budgetadjustment


type BudgetAdjustmentFiltersSlos struct {
	// slo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#slo BudgetAdjustment#slo}
	Slo interface{} `field:"required" json:"slo" yaml:"slo"`
}

