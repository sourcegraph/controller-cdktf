package googleoracledatabasecloudvmcluster


type GoogleOracleDatabaseCloudVmClusterPropertiesTimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_cloud_vm_cluster#id GoogleOracleDatabaseCloudVmCluster#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

