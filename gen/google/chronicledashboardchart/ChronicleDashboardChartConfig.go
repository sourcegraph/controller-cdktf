package chronicledashboardchart

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChronicleDashboardChartConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#dashboard_chart ChronicleDashboardChart#dashboard_chart}
	DashboardChart *ChronicleDashboardChartDashboardChart `field:"required" json:"dashboardChart" yaml:"dashboardChart"`
	// The ID of the Chronicle instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#instance ChronicleDashboardChart#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the Chronicle instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#location ChronicleDashboardChart#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// chart_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#chart_layout ChronicleDashboardChart#chart_layout}
	ChartLayout *ChronicleDashboardChartChartLayout `field:"optional" json:"chartLayout" yaml:"chartLayout"`
	// dashboard_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#dashboard_query ChronicleDashboardChart#dashboard_query}
	DashboardQuery *ChronicleDashboardChartDashboardQuery `field:"optional" json:"dashboardQuery" yaml:"dashboardQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#id ChronicleDashboardChart#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The parent NativeDashboard resource name, formatted as projects/{project}/locations/{location}/instances/{instance}/nativeDashboards/{dashboard_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#native_dashboard ChronicleDashboardChart#native_dashboard}
	NativeDashboard *string `field:"optional" json:"nativeDashboard" yaml:"nativeDashboard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#project ChronicleDashboardChart#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#timeouts ChronicleDashboardChart#timeouts}
	Timeouts *ChronicleDashboardChartTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

