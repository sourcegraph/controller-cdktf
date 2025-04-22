package googleoracledatabasecloudvmcluster


type GoogleOracleDatabaseCloudVmClusterProperties struct {
	// Number of enabled CPU cores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#cpu_core_count GoogleOracleDatabaseCloudVmCluster#cpu_core_count}
	CpuCoreCount *float64 `field:"required" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// License type of VM Cluster.   Possible values:  LICENSE_TYPE_UNSPECIFIED LICENSE_INCLUDED BRING_YOUR_OWN_LICENSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#license_type GoogleOracleDatabaseCloudVmCluster#license_type}
	LicenseType *string `field:"required" json:"licenseType" yaml:"licenseType"`
	// OCI Cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#cluster_name GoogleOracleDatabaseCloudVmCluster#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The data disk group size to be allocated in TBs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#data_storage_size_tb GoogleOracleDatabaseCloudVmCluster#data_storage_size_tb}
	DataStorageSizeTb *float64 `field:"optional" json:"dataStorageSizeTb" yaml:"dataStorageSizeTb"`
	// Local storage per VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#db_node_storage_size_gb GoogleOracleDatabaseCloudVmCluster#db_node_storage_size_gb}
	DbNodeStorageSizeGb *float64 `field:"optional" json:"dbNodeStorageSizeGb" yaml:"dbNodeStorageSizeGb"`
	// OCID of database servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#db_server_ocids GoogleOracleDatabaseCloudVmCluster#db_server_ocids}
	DbServerOcids *[]*string `field:"optional" json:"dbServerOcids" yaml:"dbServerOcids"`
	// diagnostics_data_collection_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#diagnostics_data_collection_options GoogleOracleDatabaseCloudVmCluster#diagnostics_data_collection_options}
	DiagnosticsDataCollectionOptions *GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions `field:"optional" json:"diagnosticsDataCollectionOptions" yaml:"diagnosticsDataCollectionOptions"`
	// The type of redundancy.   Possible values:  DISK_REDUNDANCY_UNSPECIFIED HIGH NORMAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#disk_redundancy GoogleOracleDatabaseCloudVmCluster#disk_redundancy}
	DiskRedundancy *string `field:"optional" json:"diskRedundancy" yaml:"diskRedundancy"`
	// Grid Infrastructure Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#gi_version GoogleOracleDatabaseCloudVmCluster#gi_version}
	GiVersion *string `field:"optional" json:"giVersion" yaml:"giVersion"`
	// Prefix for VM cluster host names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#hostname_prefix GoogleOracleDatabaseCloudVmCluster#hostname_prefix}
	HostnamePrefix *string `field:"optional" json:"hostnamePrefix" yaml:"hostnamePrefix"`
	// Use local backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#local_backup_enabled GoogleOracleDatabaseCloudVmCluster#local_backup_enabled}
	LocalBackupEnabled interface{} `field:"optional" json:"localBackupEnabled" yaml:"localBackupEnabled"`
	// Memory allocated in GBs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#memory_size_gb GoogleOracleDatabaseCloudVmCluster#memory_size_gb}
	MemorySizeGb *float64 `field:"optional" json:"memorySizeGb" yaml:"memorySizeGb"`
	// Number of database servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#node_count GoogleOracleDatabaseCloudVmCluster#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// OCPU count per VM. Minimum is 0.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#ocpu_count GoogleOracleDatabaseCloudVmCluster#ocpu_count}
	OcpuCount *float64 `field:"optional" json:"ocpuCount" yaml:"ocpuCount"`
	// Use exadata sparse snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#sparse_diskgroup_enabled GoogleOracleDatabaseCloudVmCluster#sparse_diskgroup_enabled}
	SparseDiskgroupEnabled interface{} `field:"optional" json:"sparseDiskgroupEnabled" yaml:"sparseDiskgroupEnabled"`
	// SSH public keys to be stored with cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#ssh_public_keys GoogleOracleDatabaseCloudVmCluster#ssh_public_keys}
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// time_zone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#time_zone GoogleOracleDatabaseCloudVmCluster#time_zone}
	TimeZone *GoogleOracleDatabaseCloudVmClusterPropertiesTimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
}

