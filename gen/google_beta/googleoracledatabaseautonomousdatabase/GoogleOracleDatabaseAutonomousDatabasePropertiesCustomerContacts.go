package googleoracledatabaseautonomousdatabase


type GoogleOracleDatabaseAutonomousDatabasePropertiesCustomerContacts struct {
	// The email address used by Oracle to send notifications regarding databases and infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_oracle_database_autonomous_database#email GoogleOracleDatabaseAutonomousDatabase#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

