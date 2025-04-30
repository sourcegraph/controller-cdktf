package googleaccesscontextmanagergcpuseraccessbinding


type GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettings struct {
	// active_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_access_context_manager_gcp_user_access_binding#active_settings GoogleAccessContextManagerGcpUserAccessBinding#active_settings}
	ActiveSettings *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings `field:"optional" json:"activeSettings" yaml:"activeSettings"`
	// dry_run_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_access_context_manager_gcp_user_access_binding#dry_run_settings GoogleAccessContextManagerGcpUserAccessBinding#dry_run_settings}
	DryRunSettings *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings `field:"optional" json:"dryRunSettings" yaml:"dryRunSettings"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_access_context_manager_gcp_user_access_binding#scope GoogleAccessContextManagerGcpUserAccessBinding#scope}
	Scope *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope `field:"optional" json:"scope" yaml:"scope"`
}

