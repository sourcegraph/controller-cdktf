package googleidentityplatformconfig


type GoogleIdentityPlatformConfigMultiTenant struct {
	// Whether this project can have tenants or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#allow_tenants GoogleIdentityPlatformConfig#allow_tenants}
	AllowTenants interface{} `field:"optional" json:"allowTenants" yaml:"allowTenants"`
	// The default cloud parent org or folder that the tenant project should be created under.
	//
	// The parent resource name should be in the format of "/", such as "folders/123" or "organizations/456".
	// If the value is not set, the tenant will be created under the same organization or folder as the agent project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#default_tenant_location GoogleIdentityPlatformConfig#default_tenant_location}
	DefaultTenantLocation *string `field:"optional" json:"defaultTenantLocation" yaml:"defaultTenantLocation"`
}

