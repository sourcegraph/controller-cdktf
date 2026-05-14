package chronicledashboardchart


type ChronicleDashboardChartDashboardChartVisualization struct {
	// button block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#button ChronicleDashboardChart#button}
	Button *ChronicleDashboardChartDashboardChartVisualizationButton `field:"optional" json:"button" yaml:"button"`
	// column_defs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#column_defs ChronicleDashboardChart#column_defs}
	ColumnDefs interface{} `field:"optional" json:"columnDefs" yaml:"columnDefs"`
	// google_maps_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#google_maps_config ChronicleDashboardChart#google_maps_config}
	GoogleMapsConfig *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig `field:"optional" json:"googleMapsConfig" yaml:"googleMapsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#grouping_type ChronicleDashboardChart#grouping_type}.
	GroupingType *string `field:"optional" json:"groupingType" yaml:"groupingType"`
	// legends block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#legends ChronicleDashboardChart#legends}
	Legends interface{} `field:"optional" json:"legends" yaml:"legends"`
	// markdown block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#markdown ChronicleDashboardChart#markdown}
	Markdown *ChronicleDashboardChartDashboardChartVisualizationMarkdown `field:"optional" json:"markdown" yaml:"markdown"`
	// series block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#series ChronicleDashboardChart#series}
	Series interface{} `field:"optional" json:"series" yaml:"series"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#series_column ChronicleDashboardChart#series_column}.
	SeriesColumn *[]*string `field:"optional" json:"seriesColumn" yaml:"seriesColumn"`
	// table_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#table_config ChronicleDashboardChart#table_config}
	TableConfig *ChronicleDashboardChartDashboardChartVisualizationTableConfig `field:"optional" json:"tableConfig" yaml:"tableConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#threshold_coloring_enabled ChronicleDashboardChart#threshold_coloring_enabled}.
	ThresholdColoringEnabled interface{} `field:"optional" json:"thresholdColoringEnabled" yaml:"thresholdColoringEnabled"`
	// tooltip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#tooltip ChronicleDashboardChart#tooltip}
	Tooltip *ChronicleDashboardChartDashboardChartVisualizationTooltip `field:"optional" json:"tooltip" yaml:"tooltip"`
	// visual_maps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#visual_maps ChronicleDashboardChart#visual_maps}
	VisualMaps interface{} `field:"optional" json:"visualMaps" yaml:"visualMaps"`
	// x_axes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#x_axes ChronicleDashboardChart#x_axes}
	XAxes interface{} `field:"optional" json:"xAxes" yaml:"xAxes"`
	// y_axes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#y_axes ChronicleDashboardChart#y_axes}
	YAxes interface{} `field:"optional" json:"yAxes" yaml:"yAxes"`
}

