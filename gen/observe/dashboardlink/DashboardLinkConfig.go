package dashboardlink

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardLinkConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A description for the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dashboard_link#description DashboardLink#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// OID of dashboard to link from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dashboard_link#from_dashboard DashboardLink#from_dashboard}
	FromDashboard *string `field:"required" json:"fromDashboard" yaml:"fromDashboard"`
	// The label for the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dashboard_link#link_label DashboardLink#link_label}
	LinkLabel *string `field:"required" json:"linkLabel" yaml:"linkLabel"`
	// Link name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dashboard_link#name DashboardLink#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// OID of dashboard to link to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dashboard_link#to_dashboard DashboardLink#to_dashboard}
	ToDashboard *string `field:"required" json:"toDashboard" yaml:"toDashboard"`
	// Observe folder OID for this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dashboard_link#folder DashboardLink#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Name of card to link in originating dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dashboard_link#from_card DashboardLink#from_card}
	FromCard *string `field:"optional" json:"fromCard" yaml:"fromCard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dashboard_link#id DashboardLink#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dashboard_link#workspace DashboardLink#workspace}
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

