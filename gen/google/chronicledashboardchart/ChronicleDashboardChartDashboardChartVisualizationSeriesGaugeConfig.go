package chronicledashboardchart


type ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig struct {
	// base_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#base_value ChronicleDashboardChart#base_value}
	BaseValue *ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValue `field:"optional" json:"baseValue" yaml:"baseValue"`
	// limit_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#limit_value ChronicleDashboardChart#limit_value}
	LimitValue *ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValue `field:"optional" json:"limitValue" yaml:"limitValue"`
	// threshold_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#threshold_values ChronicleDashboardChart#threshold_values}
	ThresholdValues interface{} `field:"optional" json:"thresholdValues" yaml:"thresholdValues"`
}

