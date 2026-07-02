package chronicledashboardchart

type ChronicleDashboardChartDashboardChartVisualizationLegends struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#bottom ChronicleDashboardChart#bottom}.
	Bottom *float64 `field:"optional" json:"bottom" yaml:"bottom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#id ChronicleDashboardChart#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#left ChronicleDashboardChart#left}.
	Left *float64 `field:"optional" json:"left" yaml:"left"`
	// Possible values: ["AUTO", "LEFT", "RIGHT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#legend_align ChronicleDashboardChart#legend_align}
	LegendAlign *string `field:"optional" json:"legendAlign" yaml:"legendAlign"`
	// Possible values: ["VERTICAL", "HORIZONTAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#legend_orient ChronicleDashboardChart#legend_orient}
	LegendOrient *string `field:"optional" json:"legendOrient" yaml:"legendOrient"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#padding ChronicleDashboardChart#padding}.
	Padding *[]*float64 `field:"optional" json:"padding" yaml:"padding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#right ChronicleDashboardChart#right}.
	Right *float64 `field:"optional" json:"right" yaml:"right"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#show ChronicleDashboardChart#show}.
	Show any `field:"optional" json:"show" yaml:"show"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#top ChronicleDashboardChart#top}.
	Top *float64 `field:"optional" json:"top" yaml:"top"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#z ChronicleDashboardChart#z}.
	Z *float64 `field:"optional" json:"z" yaml:"z"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#z_level ChronicleDashboardChart#z_level}.
	ZLevel *float64 `field:"optional" json:"zLevel" yaml:"zLevel"`
}
