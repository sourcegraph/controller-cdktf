package dataobservemonitor


type DataObserveMonitorRuleLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#compare_values DataObserveMonitor#compare_values}.
	CompareValues *[]*float64 `field:"optional" json:"compareValues" yaml:"compareValues"`
	// Short summary or comment of how the data for monitor is queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#expression_summary DataObserveMonitor#expression_summary}
	ExpressionSummary *string `field:"optional" json:"expressionSummary" yaml:"expressionSummary"`
	// An id of the stage that is used to generate logs for preview.
	//
	// This is usually a stage before aggregation.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#log_stage_id DataObserveMonitor#log_stage_id}
	LogStageId *string `field:"optional" json:"logStageId" yaml:"logStageId"`
	// ID of the dataset that contains logs for preview.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#source_log_dataset DataObserveMonitor#source_log_dataset}
	SourceLogDataset *string `field:"optional" json:"sourceLogDataset" yaml:"sourceLogDataset"`
}

