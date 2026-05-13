package googledataproccluster


type GoogleDataprocClusterClusterConfigSecurityConfig struct {
	// identity_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_cluster#identity_config GoogleDataprocCluster#identity_config}
	IdentityConfig *GoogleDataprocClusterClusterConfigSecurityConfigIdentityConfig `field:"optional" json:"identityConfig" yaml:"identityConfig"`
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_cluster#kerberos_config GoogleDataprocCluster#kerberos_config}
	KerberosConfig *GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfig `field:"optional" json:"kerberosConfig" yaml:"kerberosConfig"`
}

