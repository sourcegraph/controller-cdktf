package googlecomposerenvironment


type GoogleComposerEnvironmentConfigDatabaseConfig struct {
	// Optional.
	//
	// Cloud SQL machine type used by Airflow database. It has to be one of: db-n1-standard-2, db-n1-standard-4, db-n1-standard-8 or db-n1-standard-16. If not specified, db-n1-standard-2 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#machine_type GoogleComposerEnvironment#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Optional. Cloud SQL database preferred zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#zone GoogleComposerEnvironment#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

