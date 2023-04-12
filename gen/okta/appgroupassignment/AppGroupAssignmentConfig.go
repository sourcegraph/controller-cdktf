package appgroupassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppGroupAssignmentConfig struct {
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
	// App to associate group with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#app_id AppGroupAssignment#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Group associated with the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#group_id AppGroupAssignment#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#id AppGroupAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#priority AppGroupAssignment#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#profile AppGroupAssignment#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Retain the group assignment on destroy.
	//
	// If set to true, the resource will be removed from state but not from the Okta app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#retain_assignment AppGroupAssignment#retain_assignment}
	RetainAssignment interface{} `field:"optional" json:"retainAssignment" yaml:"retainAssignment"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#timeouts AppGroupAssignment#timeouts}
	Timeouts *AppGroupAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

