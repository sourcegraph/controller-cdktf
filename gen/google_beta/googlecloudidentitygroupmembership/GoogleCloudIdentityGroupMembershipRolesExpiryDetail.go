package googlecloudidentitygroupmembership


type GoogleCloudIdentityGroupMembershipRolesExpiryDetail struct {
	// The time at which the MembershipRole will expire.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits.
	//
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_identity_group_membership#expire_time GoogleCloudIdentityGroupMembership#expire_time}
	ExpireTime *string `field:"required" json:"expireTime" yaml:"expireTime"`
}

