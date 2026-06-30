package chronicledashboardchart

type ChronicleDashboardChartDashboardChartVisualizationSeries struct {
	// area_style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#area_style ChronicleDashboardChart#area_style}
	AreaStyle *ChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle `field:"optional" json:"areaStyle" yaml:"areaStyle"`
	// data_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#data_label ChronicleDashboardChart#data_label}
	DataLabel *ChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel `field:"optional" json:"dataLabel" yaml:"dataLabel"`
	// encode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#encode ChronicleDashboardChart#encode}
	Encode *ChronicleDashboardChartDashboardChartVisualizationSeriesEncode `field:"optional" json:"encode" yaml:"encode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#field ChronicleDashboardChart#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// gauge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#gauge_config ChronicleDashboardChart#gauge_config}
	GaugeConfig *ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig `field:"optional" json:"gaugeConfig" yaml:"gaugeConfig"`
	// item_colors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#item_colors ChronicleDashboardChart#item_colors}
	ItemColors *ChronicleDashboardChartDashboardChartVisualizationSeriesItemColors `field:"optional" json:"itemColors" yaml:"itemColors"`
	// item_style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#item_style ChronicleDashboardChart#item_style}
	ItemStyle *ChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle `field:"optional" json:"itemStyle" yaml:"itemStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#label ChronicleDashboardChart#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// metric_trend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#metric_trend_config ChronicleDashboardChart#metric_trend_config}
	MetricTrendConfig *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig `field:"optional" json:"metricTrendConfig" yaml:"metricTrendConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#radius ChronicleDashboardChart#radius}.
	Radius *[]*string `field:"optional" json:"radius" yaml:"radius"`
	// User specified series label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#series_name ChronicleDashboardChart#series_name}
	SeriesName *string `field:"optional" json:"seriesName" yaml:"seriesName"`
	// Possible values: ["SAMESIGN", "ALL", "POSITIVE", "NEGATIVE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#series_stack_strategy ChronicleDashboardChart#series_stack_strategy}
	SeriesStackStrategy *string `field:"optional" json:"seriesStackStrategy" yaml:"seriesStackStrategy"`
	// Possible values: ["LINE", "BAR", "PIE", "TEXT", "MAP", "GAUGE", "SCATTERPLOT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#series_type ChronicleDashboardChart#series_type}
	SeriesType *string `field:"optional" json:"seriesType" yaml:"seriesType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#series_unique_value ChronicleDashboardChart#series_unique_value}.
	SeriesUniqueValue *string `field:"optional" json:"seriesUniqueValue" yaml:"seriesUniqueValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#show_background ChronicleDashboardChart#show_background}.
	ShowBackground any `field:"optional" json:"showBackground" yaml:"showBackground"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#show_symbol ChronicleDashboardChart#show_symbol}.
	ShowSymbol any `field:"optional" json:"showSymbol" yaml:"showSymbol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#stack ChronicleDashboardChart#stack}.
	Stack *string `field:"optional" json:"stack" yaml:"stack"`
}
