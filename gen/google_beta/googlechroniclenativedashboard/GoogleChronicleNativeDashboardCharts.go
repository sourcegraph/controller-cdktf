package googlechroniclenativedashboard


type GoogleChronicleNativeDashboardCharts struct {
	// chart_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_native_dashboard#chart_layout GoogleChronicleNativeDashboard#chart_layout}
	ChartLayout *GoogleChronicleNativeDashboardChartsChartLayout `field:"optional" json:"chartLayout" yaml:"chartLayout"`
	// The resource name of the associated DashboardChart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_native_dashboard#dashboard_chart GoogleChronicleNativeDashboard#dashboard_chart}
	DashboardChart *string `field:"optional" json:"dashboardChart" yaml:"dashboardChart"`
	// List of dashboard filter IDs applied to this chart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_native_dashboard#filters_ids GoogleChronicleNativeDashboard#filters_ids}
	FiltersIds *[]*string `field:"optional" json:"filtersIds" yaml:"filtersIds"`
}

