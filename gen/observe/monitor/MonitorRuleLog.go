package monitor


type MonitorRuleLog struct {
	// Comparison function used to compare the query result against the compare_values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#compare_function Monitor#compare_function}
	CompareFunction *string `field:"required" json:"compareFunction" yaml:"compareFunction"`
	// Amount of time to evaluate query and comparison over.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#lookback_time Monitor#lookback_time}
	LookbackTime *string `field:"required" json:"lookbackTime" yaml:"lookbackTime"`
	// Value(s) to compare the query(ies) against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#compare_values Monitor#compare_values}
	CompareValues *[]*float64 `field:"optional" json:"compareValues" yaml:"compareValues"`
	// Short summary or comment of how the data for monitor is queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#expression_summary Monitor#expression_summary}
	ExpressionSummary *string `field:"optional" json:"expressionSummary" yaml:"expressionSummary"`
	// An id of the stage that is used to generate logs for preview.
	//
	// This is usually a stage before aggregation.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#log_stage_id Monitor#log_stage_id}
	LogStageId *string `field:"optional" json:"logStageId" yaml:"logStageId"`
	// ID of the dataset that contains logs for preview.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#source_log_dataset Monitor#source_log_dataset}
	SourceLogDataset *string `field:"optional" json:"sourceLogDataset" yaml:"sourceLogDataset"`
}

