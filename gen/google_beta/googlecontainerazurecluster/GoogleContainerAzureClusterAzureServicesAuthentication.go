package googlecontainerazurecluster


type GoogleContainerAzureClusterAzureServicesAuthentication struct {
	// The Azure Active Directory Application ID for Authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_cluster#application_id GoogleContainerAzureCluster#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Azure Active Directory Tenant ID for Authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_cluster#tenant_id GoogleContainerAzureCluster#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

