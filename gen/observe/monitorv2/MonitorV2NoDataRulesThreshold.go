package monitorv2


type MonitorV2NoDataRulesThreshold struct {
	// The query aggregator (AllOf, AnyOf, AvgOf, Max, Min, SumOf) for the value monitor type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#aggregation MonitorV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Indicates which column in the input query has the value to apply the aggregation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#value_column_name MonitorV2#value_column_name}
	ValueColumnName *string `field:"required" json:"valueColumnName" yaml:"valueColumnName"`
	// compare_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#compare_groups MonitorV2#compare_groups}
	CompareGroups interface{} `field:"optional" json:"compareGroups" yaml:"compareGroups"`
	// compare_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#compare_values MonitorV2#compare_values}
	CompareValues interface{} `field:"optional" json:"compareValues" yaml:"compareValues"`
}

