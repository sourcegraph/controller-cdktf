package slo


type SloObjectiveCountMetricsGoodCloudwatch struct {
	// Region of the CloudWatch instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#region Slo#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// AccountID used with cross-account observability feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#account_id Slo#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#dimensions Slo#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// JSON query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#json Slo#json}
	Json *string `field:"optional" json:"json" yaml:"json"`
	// Metric name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#metric_name Slo#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Namespace of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#namespace Slo#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// SQL query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#sql Slo#sql}
	Sql *string `field:"optional" json:"sql" yaml:"sql"`
	// Metric data aggregations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#stat Slo#stat}
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
}

