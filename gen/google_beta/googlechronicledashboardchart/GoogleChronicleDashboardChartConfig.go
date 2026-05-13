package googlechronicledashboardchart

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleChronicleDashboardChartConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// dashboard_chart block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#dashboard_chart GoogleChronicleDashboardChart#dashboard_chart}
	DashboardChart *GoogleChronicleDashboardChartDashboardChart `field:"required" json:"dashboardChart" yaml:"dashboardChart"`
	// The ID of the Chronicle instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#instance GoogleChronicleDashboardChart#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the Chronicle instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#location GoogleChronicleDashboardChart#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// chart_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#chart_layout GoogleChronicleDashboardChart#chart_layout}
	ChartLayout *GoogleChronicleDashboardChartChartLayout `field:"optional" json:"chartLayout" yaml:"chartLayout"`
	// dashboard_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#dashboard_query GoogleChronicleDashboardChart#dashboard_query}
	DashboardQuery *GoogleChronicleDashboardChartDashboardQuery `field:"optional" json:"dashboardQuery" yaml:"dashboardQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#id GoogleChronicleDashboardChart#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The parent NativeDashboard resource name, formatted as projects/{project}/locations/{location}/instances/{instance}/nativeDashboards/{dashboard_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#native_dashboard GoogleChronicleDashboardChart#native_dashboard}
	NativeDashboard *string `field:"optional" json:"nativeDashboard" yaml:"nativeDashboard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#project GoogleChronicleDashboardChart#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#timeouts GoogleChronicleDashboardChart#timeouts}
	Timeouts *GoogleChronicleDashboardChartTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

