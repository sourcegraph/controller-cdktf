package googleiapsettings


type GoogleIapSettingsAccessSettingsGcipSettings struct {
	// Login page URI associated with the GCIP tenants.
	//
	// Typically, all resources within
	// the same project share the same login page, though it could be overridden at the
	// sub resource level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#login_page_uri GoogleIapSettings#login_page_uri}
	LoginPageUri *string `field:"optional" json:"loginPageUri" yaml:"loginPageUri"`
	// GCIP tenant ids that are linked to the IAP resource.
	//
	// tenantIds could be a string
	// beginning with a number character to indicate authenticating with GCIP tenant flow,
	// or in the format of _ to indicate authenticating with GCIP agent flow. If agent flow
	// is used, tenantIds should only contain one single element, while for tenant flow,
	// tenantIds can contain multiple elements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#tenant_ids GoogleIapSettings#tenant_ids}
	TenantIds *[]*string `field:"optional" json:"tenantIds" yaml:"tenantIds"`
}

