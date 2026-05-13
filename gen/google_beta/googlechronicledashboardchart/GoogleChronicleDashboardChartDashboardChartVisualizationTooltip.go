package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationTooltip struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#show GoogleChronicleDashboardChart#show}.
	Show interface{} `field:"optional" json:"show" yaml:"show"`
	// Possible values: ["TOOLTIP_TRIGGER_UNSPECIFIED", "TOOLTIP_TRIGGER_NONE", "TOOLTIP_TRIGGER_ITEM", "TOOLTIP_TRIGGER_AXIS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#tooltip_trigger GoogleChronicleDashboardChart#tooltip_trigger}
	TooltipTrigger *string `field:"optional" json:"tooltipTrigger" yaml:"tooltipTrigger"`
}

