package slo


type SloObjectiveCountMetricsBadLogicMonitor struct {
	// Line.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#line Slo#line}
	Line *string `field:"required" json:"line" yaml:"line"`
	// Query type: device_metrics or website_metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#query_type Slo#query_type}
	QueryType *string `field:"required" json:"queryType" yaml:"queryType"`
	// Checkpoint ID. Used by Query type = website_metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#checkpoint_id Slo#checkpoint_id}
	CheckpointId *string `field:"optional" json:"checkpointId" yaml:"checkpointId"`
	// Device Datasource Instance ID. Used by Query type = device_metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#device_data_source_instance_id Slo#device_data_source_instance_id}
	DeviceDataSourceInstanceId *float64 `field:"optional" json:"deviceDataSourceInstanceId" yaml:"deviceDataSourceInstanceId"`
	// Graph ID. Used by Query type = device_metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#graph_id Slo#graph_id}
	GraphId *float64 `field:"optional" json:"graphId" yaml:"graphId"`
	// Graph Name. Used by Query type = website_metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#graph_name Slo#graph_name}
	GraphName *string `field:"optional" json:"graphName" yaml:"graphName"`
	// Website ID. Used by Query type = website_metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#website_id Slo#website_id}
	WebsiteId *string `field:"optional" json:"websiteId" yaml:"websiteId"`
}

