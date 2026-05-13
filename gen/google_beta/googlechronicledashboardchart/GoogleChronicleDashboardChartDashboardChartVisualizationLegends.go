package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationLegends struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#bottom GoogleChronicleDashboardChart#bottom}.
	Bottom *float64 `field:"optional" json:"bottom" yaml:"bottom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#id GoogleChronicleDashboardChart#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#left GoogleChronicleDashboardChart#left}.
	Left *float64 `field:"optional" json:"left" yaml:"left"`
	// Possible values: ["AUTO", "LEFT", "RIGHT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#legend_align GoogleChronicleDashboardChart#legend_align}
	LegendAlign *string `field:"optional" json:"legendAlign" yaml:"legendAlign"`
	// Possible values: ["VERTICAL", "HORIZONTAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#legend_orient GoogleChronicleDashboardChart#legend_orient}
	LegendOrient *string `field:"optional" json:"legendOrient" yaml:"legendOrient"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#padding GoogleChronicleDashboardChart#padding}.
	Padding *[]*float64 `field:"optional" json:"padding" yaml:"padding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#right GoogleChronicleDashboardChart#right}.
	Right *float64 `field:"optional" json:"right" yaml:"right"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#show GoogleChronicleDashboardChart#show}.
	Show interface{} `field:"optional" json:"show" yaml:"show"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#top GoogleChronicleDashboardChart#top}.
	Top *float64 `field:"optional" json:"top" yaml:"top"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#z GoogleChronicleDashboardChart#z}.
	Z *float64 `field:"optional" json:"z" yaml:"z"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#z_level GoogleChronicleDashboardChart#z_level}.
	ZLevel *float64 `field:"optional" json:"zLevel" yaml:"zLevel"`
}

