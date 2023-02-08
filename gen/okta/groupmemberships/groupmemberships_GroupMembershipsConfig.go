package groupmemberships

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupMembershipsConfig struct {
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
	// ID of a Okta group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_memberships#group_id GroupMemberships#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The list of Okta user IDs which the group should have membership managed for.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_memberships#users GroupMemberships#users}
	Users *[]*string `field:"required" json:"users" yaml:"users"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_memberships#id GroupMemberships#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource concerns itself with all users added/deleted to the group; even those managed outside of the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_memberships#track_all_users GroupMemberships#track_all_users}
	TrackAllUsers interface{} `field:"optional" json:"trackAllUsers" yaml:"trackAllUsers"`
}

