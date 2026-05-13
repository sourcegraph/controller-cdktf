package googleoracledatabasedbsystem


type GoogleOracleDatabaseDbSystemProperties struct {
	// The number of CPU cores to enable for the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#compute_count GoogleOracleDatabaseDbSystem#compute_count}
	ComputeCount *float64 `field:"required" json:"computeCount" yaml:"computeCount"`
	// The database edition of the DbSystem. Possible values: STANDARD_EDITION ENTERPRISE_EDITION ENTERPRISE_EDITION_HIGH_PERFORMANCE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#database_edition GoogleOracleDatabaseDbSystem#database_edition}
	DatabaseEdition *string `field:"required" json:"databaseEdition" yaml:"databaseEdition"`
	// The initial data storage size in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#initial_data_storage_size_gb GoogleOracleDatabaseDbSystem#initial_data_storage_size_gb}
	InitialDataStorageSizeGb *float64 `field:"required" json:"initialDataStorageSizeGb" yaml:"initialDataStorageSizeGb"`
	// The license model of the DbSystem. Possible values: LICENSE_INCLUDED BRING_YOUR_OWN_LICENSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#license_model GoogleOracleDatabaseDbSystem#license_model}
	LicenseModel *string `field:"required" json:"licenseModel" yaml:"licenseModel"`
	// Shape of DB System.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#shape GoogleOracleDatabaseDbSystem#shape}
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// SSH public keys to be stored with the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#ssh_public_keys GoogleOracleDatabaseDbSystem#ssh_public_keys}
	SshPublicKeys *[]*string `field:"required" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// The compute model of the DbSystem. Possible values: ECPU OCPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#compute_model GoogleOracleDatabaseDbSystem#compute_model}
	ComputeModel *string `field:"optional" json:"computeModel" yaml:"computeModel"`
	// data_collection_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#data_collection_options GoogleOracleDatabaseDbSystem#data_collection_options}
	DataCollectionOptions *GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptions `field:"optional" json:"dataCollectionOptions" yaml:"dataCollectionOptions"`
	// The data storage size in GB that is currently available to DbSystems.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#data_storage_size_gb GoogleOracleDatabaseDbSystem#data_storage_size_gb}
	DataStorageSizeGb *float64 `field:"optional" json:"dataStorageSizeGb" yaml:"dataStorageSizeGb"`
	// db_home block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#db_home GoogleOracleDatabaseDbSystem#db_home}
	DbHome *GoogleOracleDatabaseDbSystemPropertiesDbHome `field:"optional" json:"dbHome" yaml:"dbHome"`
	// db_system_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#db_system_options GoogleOracleDatabaseDbSystem#db_system_options}
	DbSystemOptions *GoogleOracleDatabaseDbSystemPropertiesDbSystemOptions `field:"optional" json:"dbSystemOptions" yaml:"dbSystemOptions"`
	// The host domain name of the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#domain GoogleOracleDatabaseDbSystem#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Prefix for DB System host names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#hostname_prefix GoogleOracleDatabaseDbSystem#hostname_prefix}
	HostnamePrefix *string `field:"optional" json:"hostnamePrefix" yaml:"hostnamePrefix"`
	// The memory size in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#memory_size_gb GoogleOracleDatabaseDbSystem#memory_size_gb}
	MemorySizeGb *float64 `field:"optional" json:"memorySizeGb" yaml:"memorySizeGb"`
	// The number of nodes in the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#node_count GoogleOracleDatabaseDbSystem#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// The private IP address of the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#private_ip GoogleOracleDatabaseDbSystem#private_ip}
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// The reco/redo storage size in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#reco_storage_size_gb GoogleOracleDatabaseDbSystem#reco_storage_size_gb}
	RecoStorageSizeGb *float64 `field:"optional" json:"recoStorageSizeGb" yaml:"recoStorageSizeGb"`
	// time_zone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#time_zone GoogleOracleDatabaseDbSystem#time_zone}
	TimeZone *GoogleOracleDatabaseDbSystemPropertiesTimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
}

