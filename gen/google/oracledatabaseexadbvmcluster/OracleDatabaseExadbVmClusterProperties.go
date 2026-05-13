package oracledatabaseexadbvmcluster


type OracleDatabaseExadbVmClusterProperties struct {
	// The number of ECPUs enabled per node for an exadata vm cluster on exascale infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#enabled_ecpu_count_per_node OracleDatabaseExadbVmCluster#enabled_ecpu_count_per_node}
	EnabledEcpuCountPerNode *float64 `field:"required" json:"enabledEcpuCountPerNode" yaml:"enabledEcpuCountPerNode"`
	// The name of ExascaleDbStorageVault associated with the ExadbVmCluster.
	//
	// It can refer to an existing ExascaleDbStorageVault. Or a new one can be
	// created during the ExadbVmCluster creation (requires
	// storage_vault_properties to be set).
	// Format:
	// projects/{project}/locations/{location}/exascaleDbStorageVaults/{exascale_db_storage_vault}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#exascale_db_storage_vault OracleDatabaseExadbVmCluster#exascale_db_storage_vault}
	ExascaleDbStorageVault *string `field:"required" json:"exascaleDbStorageVault" yaml:"exascaleDbStorageVault"`
	// Grid Infrastructure Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#grid_image_id OracleDatabaseExadbVmCluster#grid_image_id}
	GridImageId *string `field:"required" json:"gridImageId" yaml:"gridImageId"`
	// Prefix for VM cluster host names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#hostname_prefix OracleDatabaseExadbVmCluster#hostname_prefix}
	HostnamePrefix *string `field:"required" json:"hostnamePrefix" yaml:"hostnamePrefix"`
	// The number of nodes/VMs in the ExadbVmCluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#node_count OracleDatabaseExadbVmCluster#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// The shape attribute of the VM cluster.
	//
	// The type of Exascale storage used
	// for Exadata VM cluster. The default is SMART_STORAGE which supports Oracle
	// Database 23ai and later
	// Possible values:
	// SMART_STORAGE
	// BLOCK_STORAGE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#shape_attribute OracleDatabaseExadbVmCluster#shape_attribute}
	ShapeAttribute *string `field:"required" json:"shapeAttribute" yaml:"shapeAttribute"`
	// The SSH public keys for the ExadbVmCluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#ssh_public_keys OracleDatabaseExadbVmCluster#ssh_public_keys}
	SshPublicKeys *[]*string `field:"required" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// vm_file_system_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#vm_file_system_storage OracleDatabaseExadbVmCluster#vm_file_system_storage}
	VmFileSystemStorage *OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage `field:"required" json:"vmFileSystemStorage" yaml:"vmFileSystemStorage"`
	// The number of additional ECPUs per node for an Exadata VM cluster on exascale infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#additional_ecpu_count_per_node OracleDatabaseExadbVmCluster#additional_ecpu_count_per_node}
	AdditionalEcpuCountPerNode *float64 `field:"optional" json:"additionalEcpuCountPerNode" yaml:"additionalEcpuCountPerNode"`
	// The cluster name for Exascale vm cluster.
	//
	// The cluster name must begin with
	// an alphabetic character and may contain hyphens(-) but can not contain
	// underscores(_). It should be not more than 11 characters and is not case
	// sensitive.
	// OCI Cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#cluster_name OracleDatabaseExadbVmCluster#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// data_collection_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#data_collection_options OracleDatabaseExadbVmCluster#data_collection_options}
	DataCollectionOptions *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions `field:"optional" json:"dataCollectionOptions" yaml:"dataCollectionOptions"`
	// The license type of the ExadbVmCluster. Possible values: LICENSE_INCLUDED BRING_YOUR_OWN_LICENSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#license_model OracleDatabaseExadbVmCluster#license_model}
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// SCAN listener port - TCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#scan_listener_port_tcp OracleDatabaseExadbVmCluster#scan_listener_port_tcp}
	ScanListenerPortTcp *float64 `field:"optional" json:"scanListenerPortTcp" yaml:"scanListenerPortTcp"`
	// time_zone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#time_zone OracleDatabaseExadbVmCluster#time_zone}
	TimeZone *OracleDatabaseExadbVmClusterPropertiesTimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
}

