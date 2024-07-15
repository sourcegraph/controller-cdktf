package googlesccmanagementfoldersecurityhealthanalyticscustommodule


type GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector struct {
	// The resource types to run the detector on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_folder_security_health_analytics_custom_module#resource_types GoogleSccManagementFolderSecurityHealthAnalyticsCustomModule#resource_types}
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

