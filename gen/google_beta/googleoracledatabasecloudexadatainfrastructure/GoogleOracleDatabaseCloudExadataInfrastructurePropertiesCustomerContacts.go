package googleoracledatabasecloudexadatainfrastructure


type GoogleOracleDatabaseCloudExadataInfrastructurePropertiesCustomerContacts struct {
	// The email address used by Oracle to send notifications regarding databases and infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure#email GoogleOracleDatabaseCloudExadataInfrastructure#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

