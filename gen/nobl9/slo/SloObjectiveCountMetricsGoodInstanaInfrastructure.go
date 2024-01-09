package slo


type SloObjectiveCountMetricsGoodInstanaInfrastructure struct {
	// Metric ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#metric_id Slo#metric_id}
	MetricId *string `field:"required" json:"metricId" yaml:"metricId"`
	// Metric retrieval method 'query' or 'snapshot'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#metric_retrieval_method Slo#metric_retrieval_method}
	MetricRetrievalMethod *string `field:"required" json:"metricRetrievalMethod" yaml:"metricRetrievalMethod"`
	// Plugin ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#plugin_id Slo#plugin_id}
	PluginId *string `field:"required" json:"pluginId" yaml:"pluginId"`
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#query Slo#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// Snapshot ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#snapshot_id Slo#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
}

