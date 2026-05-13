package googlecloudidentitypolicy


type GoogleCloudIdentityPolicyPolicyQuery struct {
	// The OrgUnit the query applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#org_unit GoogleCloudIdentityPolicy#org_unit}
	OrgUnit *string `field:"required" json:"orgUnit" yaml:"orgUnit"`
	// The group that the query applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#group GoogleCloudIdentityPolicy#group}
	Group *string `field:"optional" json:"group" yaml:"group"`
	// The CEL query that defines which entities the Policy applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#query GoogleCloudIdentityPolicy#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

