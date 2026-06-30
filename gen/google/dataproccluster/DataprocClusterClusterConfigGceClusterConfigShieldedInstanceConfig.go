package dataproccluster

type DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig struct {
	// Defines whether instances have integrity monitoring enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataproc_cluster#enable_integrity_monitoring DataprocCluster#enable_integrity_monitoring}
	EnableIntegrityMonitoring any `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Defines whether instances have Secure Boot enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataproc_cluster#enable_secure_boot DataprocCluster#enable_secure_boot}
	EnableSecureBoot any `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Defines whether instances have the vTPM enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataproc_cluster#enable_vtpm DataprocCluster#enable_vtpm}
	EnableVtpm any `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}
