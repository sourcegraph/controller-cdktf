package googlecontainercluster


type GoogleContainerClusterNodeConfigShieldedInstanceConfig struct {
	// Defines whether the instance has integrity monitoring enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enable_integrity_monitoring GoogleContainerCluster#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Defines whether the instance has Secure Boot enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enable_secure_boot GoogleContainerCluster#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
}
