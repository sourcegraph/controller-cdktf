package groupowner

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupOwnerConfig struct {
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
	// The id of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_owner#group_id GroupOwner#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The user id of the group owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_owner#id_of_group_owner GroupOwner#id_of_group_owner}
	IdOfGroupOwner *string `field:"required" json:"idOfGroupOwner" yaml:"idOfGroupOwner"`
	// The entity type of the owner. Enum: "GROUP" "USER".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_owner#type GroupOwner#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

