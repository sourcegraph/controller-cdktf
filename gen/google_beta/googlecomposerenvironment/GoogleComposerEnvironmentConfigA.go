package googlecomposerenvironment


type GoogleComposerEnvironmentConfigA struct {
	// database_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#database_config GoogleComposerEnvironment#database_config}
	DatabaseConfig *GoogleComposerEnvironmentConfigDatabaseConfig `field:"optional" json:"databaseConfig" yaml:"databaseConfig"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#encryption_config GoogleComposerEnvironment#encryption_config}
	EncryptionConfig *GoogleComposerEnvironmentConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The size of the Cloud Composer environment.
	//
	// This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#environment_size GoogleComposerEnvironment#environment_size}
	EnvironmentSize *string `field:"optional" json:"environmentSize" yaml:"environmentSize"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#maintenance_window GoogleComposerEnvironment#maintenance_window}
	MaintenanceWindow *GoogleComposerEnvironmentConfigMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// master_authorized_networks_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#master_authorized_networks_config GoogleComposerEnvironment#master_authorized_networks_config}
	MasterAuthorizedNetworksConfig *GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfig `field:"optional" json:"masterAuthorizedNetworksConfig" yaml:"masterAuthorizedNetworksConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#node_config GoogleComposerEnvironment#node_config}
	NodeConfig *GoogleComposerEnvironmentConfigNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// The number of nodes in the Kubernetes Engine cluster that will be used to run this environment.
	//
	// This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#node_count GoogleComposerEnvironment#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// private_environment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#private_environment_config GoogleComposerEnvironment#private_environment_config}
	PrivateEnvironmentConfig *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig `field:"optional" json:"privateEnvironmentConfig" yaml:"privateEnvironmentConfig"`
	// recovery_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#recovery_config GoogleComposerEnvironment#recovery_config}
	RecoveryConfig *GoogleComposerEnvironmentConfigRecoveryConfig `field:"optional" json:"recoveryConfig" yaml:"recoveryConfig"`
	// software_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#software_config GoogleComposerEnvironment#software_config}
	SoftwareConfig *GoogleComposerEnvironmentConfigSoftwareConfig `field:"optional" json:"softwareConfig" yaml:"softwareConfig"`
	// web_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#web_server_config GoogleComposerEnvironment#web_server_config}
	WebServerConfig *GoogleComposerEnvironmentConfigWebServerConfig `field:"optional" json:"webServerConfig" yaml:"webServerConfig"`
	// web_server_network_access_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#web_server_network_access_control GoogleComposerEnvironment#web_server_network_access_control}
	WebServerNetworkAccessControl *GoogleComposerEnvironmentConfigWebServerNetworkAccessControl `field:"optional" json:"webServerNetworkAccessControl" yaml:"webServerNetworkAccessControl"`
	// workloads_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_composer_environment#workloads_config GoogleComposerEnvironment#workloads_config}
	WorkloadsConfig *GoogleComposerEnvironmentConfigWorkloadsConfig `field:"optional" json:"workloadsConfig" yaml:"workloadsConfig"`
}

