package googledataproccluster


type GoogleDataprocClusterClusterConfigDataprocMetricConfigMetrics struct {
	// A source for the collection of Dataproc OSS metrics (see [available OSS metrics] (https://cloud.google.com//dataproc/docs/guides/monitoring#available_oss_metrics)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#metric_source GoogleDataprocCluster#metric_source}
	MetricSource *string `field:"required" json:"metricSource" yaml:"metricSource"`
	// Specify one or more [available OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics) to collect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#metric_overrides GoogleDataprocCluster#metric_overrides}
	MetricOverrides *[]*string `field:"optional" json:"metricOverrides" yaml:"metricOverrides"`
}

