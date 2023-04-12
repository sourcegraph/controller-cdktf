package adminrolecustomassignments

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdminRoleCustomAssignmentsConfig struct {
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
	// ID of the Custom Role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/admin_role_custom_assignments#custom_role_id AdminRoleCustomAssignments#custom_role_id}
	CustomRoleId *string `field:"required" json:"customRoleId" yaml:"customRoleId"`
	// ID of the target Resource Set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/admin_role_custom_assignments#resource_set_id AdminRoleCustomAssignments#resource_set_id}
	ResourceSetId *string `field:"required" json:"resourceSetId" yaml:"resourceSetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/admin_role_custom_assignments#id AdminRoleCustomAssignments#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The hrefs that point to User(s) and/or Group(s) that receive the Role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/admin_role_custom_assignments#members AdminRoleCustomAssignments#members}
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
}

