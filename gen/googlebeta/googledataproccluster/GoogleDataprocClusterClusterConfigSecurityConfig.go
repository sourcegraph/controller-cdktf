package googledataproccluster


type GoogleDataprocClusterClusterConfigSecurityConfig struct {
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster#kerberos_config GoogleDataprocCluster#kerberos_config}
	KerberosConfig *GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfig `field:"required" json:"kerberosConfig" yaml:"kerberosConfig"`
}

