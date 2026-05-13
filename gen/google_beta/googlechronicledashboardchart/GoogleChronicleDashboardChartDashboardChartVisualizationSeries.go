package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationSeries struct {
	// area_style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#area_style GoogleChronicleDashboardChart#area_style}
	AreaStyle *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle `field:"optional" json:"areaStyle" yaml:"areaStyle"`
	// data_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#data_label GoogleChronicleDashboardChart#data_label}
	DataLabel *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel `field:"optional" json:"dataLabel" yaml:"dataLabel"`
	// encode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#encode GoogleChronicleDashboardChart#encode}
	Encode *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesEncode `field:"optional" json:"encode" yaml:"encode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#field GoogleChronicleDashboardChart#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// gauge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#gauge_config GoogleChronicleDashboardChart#gauge_config}
	GaugeConfig *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig `field:"optional" json:"gaugeConfig" yaml:"gaugeConfig"`
	// item_colors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#item_colors GoogleChronicleDashboardChart#item_colors}
	ItemColors *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemColors `field:"optional" json:"itemColors" yaml:"itemColors"`
	// item_style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#item_style GoogleChronicleDashboardChart#item_style}
	ItemStyle *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle `field:"optional" json:"itemStyle" yaml:"itemStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#label GoogleChronicleDashboardChart#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// metric_trend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#metric_trend_config GoogleChronicleDashboardChart#metric_trend_config}
	MetricTrendConfig *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig `field:"optional" json:"metricTrendConfig" yaml:"metricTrendConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#radius GoogleChronicleDashboardChart#radius}.
	Radius *[]*string `field:"optional" json:"radius" yaml:"radius"`
	// User specified series label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#series_name GoogleChronicleDashboardChart#series_name}
	SeriesName *string `field:"optional" json:"seriesName" yaml:"seriesName"`
	// Possible values: ["SAMESIGN", "ALL", "POSITIVE", "NEGATIVE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#series_stack_strategy GoogleChronicleDashboardChart#series_stack_strategy}
	SeriesStackStrategy *string `field:"optional" json:"seriesStackStrategy" yaml:"seriesStackStrategy"`
	// Possible values: ["LINE", "BAR", "PIE", "TEXT", "MAP", "GAUGE", "SCATTERPLOT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#series_type GoogleChronicleDashboardChart#series_type}
	SeriesType *string `field:"optional" json:"seriesType" yaml:"seriesType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#series_unique_value GoogleChronicleDashboardChart#series_unique_value}.
	SeriesUniqueValue *string `field:"optional" json:"seriesUniqueValue" yaml:"seriesUniqueValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#show_background GoogleChronicleDashboardChart#show_background}.
	ShowBackground interface{} `field:"optional" json:"showBackground" yaml:"showBackground"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#show_symbol GoogleChronicleDashboardChart#show_symbol}.
	ShowSymbol interface{} `field:"optional" json:"showSymbol" yaml:"showSymbol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#stack GoogleChronicleDashboardChart#stack}.
	Stack *string `field:"optional" json:"stack" yaml:"stack"`
}

