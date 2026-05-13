package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig struct {
	// base_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#base_value GoogleChronicleDashboardChart#base_value}
	BaseValue *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValue `field:"optional" json:"baseValue" yaml:"baseValue"`
	// limit_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#limit_value GoogleChronicleDashboardChart#limit_value}
	LimitValue *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValue `field:"optional" json:"limitValue" yaml:"limitValue"`
	// threshold_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#threshold_values GoogleChronicleDashboardChart#threshold_values}
	ThresholdValues interface{} `field:"optional" json:"thresholdValues" yaml:"thresholdValues"`
}

