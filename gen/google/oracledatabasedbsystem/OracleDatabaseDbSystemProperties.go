package oracledatabasedbsystem


type OracleDatabaseDbSystemProperties struct {
	// The number of CPU cores to enable for the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#compute_count OracleDatabaseDbSystem#compute_count}
	ComputeCount *float64 `field:"required" json:"computeCount" yaml:"computeCount"`
	// The database edition of the DbSystem. Possible values: STANDARD_EDITION ENTERPRISE_EDITION ENTERPRISE_EDITION_HIGH_PERFORMANCE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#database_edition OracleDatabaseDbSystem#database_edition}
	DatabaseEdition *string `field:"required" json:"databaseEdition" yaml:"databaseEdition"`
	// The initial data storage size in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#initial_data_storage_size_gb OracleDatabaseDbSystem#initial_data_storage_size_gb}
	InitialDataStorageSizeGb *float64 `field:"required" json:"initialDataStorageSizeGb" yaml:"initialDataStorageSizeGb"`
	// The license model of the DbSystem. Possible values: LICENSE_INCLUDED BRING_YOUR_OWN_LICENSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#license_model OracleDatabaseDbSystem#license_model}
	LicenseModel *string `field:"required" json:"licenseModel" yaml:"licenseModel"`
	// Shape of DB System.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#shape OracleDatabaseDbSystem#shape}
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// SSH public keys to be stored with the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#ssh_public_keys OracleDatabaseDbSystem#ssh_public_keys}
	SshPublicKeys *[]*string `field:"required" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// The compute model of the DbSystem. Possible values: ECPU OCPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#compute_model OracleDatabaseDbSystem#compute_model}
	ComputeModel *string `field:"optional" json:"computeModel" yaml:"computeModel"`
	// data_collection_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#data_collection_options OracleDatabaseDbSystem#data_collection_options}
	DataCollectionOptions *OracleDatabaseDbSystemPropertiesDataCollectionOptions `field:"optional" json:"dataCollectionOptions" yaml:"dataCollectionOptions"`
	// The data storage size in GB that is currently available to DbSystems.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#data_storage_size_gb OracleDatabaseDbSystem#data_storage_size_gb}
	DataStorageSizeGb *float64 `field:"optional" json:"dataStorageSizeGb" yaml:"dataStorageSizeGb"`
	// db_home block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#db_home OracleDatabaseDbSystem#db_home}
	DbHome *OracleDatabaseDbSystemPropertiesDbHome `field:"optional" json:"dbHome" yaml:"dbHome"`
	// db_system_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#db_system_options OracleDatabaseDbSystem#db_system_options}
	DbSystemOptions *OracleDatabaseDbSystemPropertiesDbSystemOptions `field:"optional" json:"dbSystemOptions" yaml:"dbSystemOptions"`
	// The host domain name of the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#domain OracleDatabaseDbSystem#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Prefix for DB System host names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#hostname_prefix OracleDatabaseDbSystem#hostname_prefix}
	HostnamePrefix *string `field:"optional" json:"hostnamePrefix" yaml:"hostnamePrefix"`
	// The memory size in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#memory_size_gb OracleDatabaseDbSystem#memory_size_gb}
	MemorySizeGb *float64 `field:"optional" json:"memorySizeGb" yaml:"memorySizeGb"`
	// The number of nodes in the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#node_count OracleDatabaseDbSystem#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// The private IP address of the DbSystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#private_ip OracleDatabaseDbSystem#private_ip}
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// The reco/redo storage size in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#reco_storage_size_gb OracleDatabaseDbSystem#reco_storage_size_gb}
	RecoStorageSizeGb *float64 `field:"optional" json:"recoStorageSizeGb" yaml:"recoStorageSizeGb"`
	// time_zone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#time_zone OracleDatabaseDbSystem#time_zone}
	TimeZone *OracleDatabaseDbSystemPropertiesTimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
}

