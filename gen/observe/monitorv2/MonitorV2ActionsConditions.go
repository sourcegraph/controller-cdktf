package monitorv2


type MonitorV2ActionsConditions struct {
	// compare_terms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#compare_terms MonitorV2#compare_terms}
	CompareTerms interface{} `field:"required" json:"compareTerms" yaml:"compareTerms"`
	// Boolean operator to combine the list of compare terms. Can be "and" (default) or "or".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#operator MonitorV2#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

