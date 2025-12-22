package rbacgroupmember


type RbacGroupMemberMember struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_group_member#group RbacGroupMember#group}.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_group_member#user RbacGroupMember#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

