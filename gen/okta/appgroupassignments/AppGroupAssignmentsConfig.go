package appgroupassignments

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppGroupAssignmentsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignments#app_id AppGroupAssignments#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignments#group AppGroupAssignments#group}
	Group interface{} `field:"required" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignments#id AppGroupAssignments#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

