package budgetadjustment


type BudgetAdjustmentFilters struct {
	// slos block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#slos BudgetAdjustment#slos}
	Slos interface{} `field:"required" json:"slos" yaml:"slos"`
}

