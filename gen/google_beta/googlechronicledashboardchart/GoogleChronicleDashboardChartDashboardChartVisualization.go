package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualization struct {
	// button block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#button GoogleChronicleDashboardChart#button}
	Button *GoogleChronicleDashboardChartDashboardChartVisualizationButton `field:"optional" json:"button" yaml:"button"`
	// column_defs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#column_defs GoogleChronicleDashboardChart#column_defs}
	ColumnDefs interface{} `field:"optional" json:"columnDefs" yaml:"columnDefs"`
	// google_maps_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#google_maps_config GoogleChronicleDashboardChart#google_maps_config}
	GoogleMapsConfig *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig `field:"optional" json:"googleMapsConfig" yaml:"googleMapsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#grouping_type GoogleChronicleDashboardChart#grouping_type}.
	GroupingType *string `field:"optional" json:"groupingType" yaml:"groupingType"`
	// legends block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#legends GoogleChronicleDashboardChart#legends}
	Legends interface{} `field:"optional" json:"legends" yaml:"legends"`
	// markdown block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#markdown GoogleChronicleDashboardChart#markdown}
	Markdown *GoogleChronicleDashboardChartDashboardChartVisualizationMarkdown `field:"optional" json:"markdown" yaml:"markdown"`
	// series block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#series GoogleChronicleDashboardChart#series}
	Series interface{} `field:"optional" json:"series" yaml:"series"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#series_column GoogleChronicleDashboardChart#series_column}.
	SeriesColumn *[]*string `field:"optional" json:"seriesColumn" yaml:"seriesColumn"`
	// table_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#table_config GoogleChronicleDashboardChart#table_config}
	TableConfig *GoogleChronicleDashboardChartDashboardChartVisualizationTableConfig `field:"optional" json:"tableConfig" yaml:"tableConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#threshold_coloring_enabled GoogleChronicleDashboardChart#threshold_coloring_enabled}.
	ThresholdColoringEnabled interface{} `field:"optional" json:"thresholdColoringEnabled" yaml:"thresholdColoringEnabled"`
	// tooltip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#tooltip GoogleChronicleDashboardChart#tooltip}
	Tooltip *GoogleChronicleDashboardChartDashboardChartVisualizationTooltip `field:"optional" json:"tooltip" yaml:"tooltip"`
	// visual_maps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#visual_maps GoogleChronicleDashboardChart#visual_maps}
	VisualMaps interface{} `field:"optional" json:"visualMaps" yaml:"visualMaps"`
	// x_axes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#x_axes GoogleChronicleDashboardChart#x_axes}
	XAxes interface{} `field:"optional" json:"xAxes" yaml:"xAxes"`
	// y_axes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#y_axes GoogleChronicleDashboardChart#y_axes}
	YAxes interface{} `field:"optional" json:"yAxes" yaml:"yAxes"`
}

