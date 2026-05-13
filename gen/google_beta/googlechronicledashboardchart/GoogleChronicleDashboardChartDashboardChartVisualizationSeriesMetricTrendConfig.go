package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig struct {
	// Possible values: ["METRIC_DISPLAY_TREND_UNSPECIFIED", "METRIC_DISPLAY_TREND_ABSOLUTE_VALUE", "METRIC_DISPLAY_TREND_PERCENTAGE", "METRIC_DISPLAY_TREND_ABSOLUTE_VALUE_AND_PERCENTAGE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#metric_display_trend GoogleChronicleDashboardChart#metric_display_trend}
	MetricDisplayTrend *string `field:"optional" json:"metricDisplayTrend" yaml:"metricDisplayTrend"`
	// Possible values: ["METRIC_FORMAT_UNSPECIFIED", "METRIC_FORMAT_NUMBER", "METRIC_FORMAT_PLAIN_TEXT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#metric_format GoogleChronicleDashboardChart#metric_format}
	MetricFormat *string `field:"optional" json:"metricFormat" yaml:"metricFormat"`
	// Possible values: ["METRIC_TREND_TYPE_UNSPECIFIED", "METRIC_TREND_TYPE_REGULAR", "METRIC_TREND_TYPE_INVERSE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#metric_trend_type GoogleChronicleDashboardChart#metric_trend_type}
	MetricTrendType *string `field:"optional" json:"metricTrendType" yaml:"metricTrendType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#show_metric_trend GoogleChronicleDashboardChart#show_metric_trend}.
	ShowMetricTrend interface{} `field:"optional" json:"showMetricTrend" yaml:"showMetricTrend"`
}

