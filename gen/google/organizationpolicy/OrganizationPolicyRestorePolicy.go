package organizationpolicy

type OrganizationPolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/organization_policy#default OrganizationPolicy#default}
	Default any `field:"required" json:"default" yaml:"default"`
}
