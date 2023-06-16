package googlecloudidentitygroupmembership


type GoogleCloudIdentityGroupMembershipRoles struct {
	// The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER. Possible values: ["OWNER", "MANAGER", "MEMBER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloud_identity_group_membership#name GoogleCloudIdentityGroupMembership#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

