package resourcegrants


type ResourceGrantsGrant struct {
	// The role to grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/resource_grants#role ResourceGrants#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// OID of the subject. Must be a user or a group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/resource_grants#subject ResourceGrants#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

