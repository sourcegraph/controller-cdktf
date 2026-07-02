package chronicledashboardchart

type ChronicleDashboardChartDashboardChartVisualizationTooltip struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#show ChronicleDashboardChart#show}.
	Show any `field:"optional" json:"show" yaml:"show"`
	// Possible values: ["TOOLTIP_TRIGGER_UNSPECIFIED", "TOOLTIP_TRIGGER_NONE", "TOOLTIP_TRIGGER_ITEM", "TOOLTIP_TRIGGER_AXIS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#tooltip_trigger ChronicleDashboardChart#tooltip_trigger}
	TooltipTrigger *string `field:"optional" json:"tooltipTrigger" yaml:"tooltipTrigger"`
}
