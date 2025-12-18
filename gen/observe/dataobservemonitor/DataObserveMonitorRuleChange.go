package dataobservemonitor


type DataObserveMonitorRuleChange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#compare_function DataObserveMonitor#compare_function}.
	CompareFunction *string `field:"required" json:"compareFunction" yaml:"compareFunction"`
}

