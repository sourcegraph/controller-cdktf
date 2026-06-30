package chronicledashboardchart

type ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig struct {
	// Possible values: ["METRIC_DISPLAY_TREND_UNSPECIFIED", "METRIC_DISPLAY_TREND_ABSOLUTE_VALUE", "METRIC_DISPLAY_TREND_PERCENTAGE", "METRIC_DISPLAY_TREND_ABSOLUTE_VALUE_AND_PERCENTAGE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#metric_display_trend ChronicleDashboardChart#metric_display_trend}
	MetricDisplayTrend *string `field:"optional" json:"metricDisplayTrend" yaml:"metricDisplayTrend"`
	// Possible values: ["METRIC_FORMAT_UNSPECIFIED", "METRIC_FORMAT_NUMBER", "METRIC_FORMAT_PLAIN_TEXT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#metric_format ChronicleDashboardChart#metric_format}
	MetricFormat *string `field:"optional" json:"metricFormat" yaml:"metricFormat"`
	// Possible values: ["METRIC_TREND_TYPE_UNSPECIFIED", "METRIC_TREND_TYPE_REGULAR", "METRIC_TREND_TYPE_INVERSE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#metric_trend_type ChronicleDashboardChart#metric_trend_type}
	MetricTrendType *string `field:"optional" json:"metricTrendType" yaml:"metricTrendType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#show_metric_trend ChronicleDashboardChart#show_metric_trend}.
	ShowMetricTrend any `field:"optional" json:"showMetricTrend" yaml:"showMetricTrend"`
}
