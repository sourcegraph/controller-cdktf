package poller


type PollerGcpMonitoring struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#json_key Poller#json_key}.
	JsonKey *string `field:"required" json:"jsonKey" yaml:"jsonKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#project_id Poller#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#exclude_metric_type_prefixes Poller#exclude_metric_type_prefixes}.
	ExcludeMetricTypePrefixes *[]*string `field:"optional" json:"excludeMetricTypePrefixes" yaml:"excludeMetricTypePrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#include_metric_type_prefixes Poller#include_metric_type_prefixes}.
	IncludeMetricTypePrefixes *[]*string `field:"optional" json:"includeMetricTypePrefixes" yaml:"includeMetricTypePrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#rate_limit Poller#rate_limit}.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#total_limit Poller#total_limit}.
	TotalLimit *float64 `field:"optional" json:"totalLimit" yaml:"totalLimit"`
}

