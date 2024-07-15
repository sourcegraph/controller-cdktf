package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig struct {
	// keytab block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#keytab GoogleDataprocMetastoreService#keytab}
	Keytab *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab `field:"required" json:"keytab" yaml:"keytab"`
	// A Cloud Storage URI that specifies the path to a krb5.conf file. It is of the form gs://{bucket_name}/path/to/krb5.conf, although the file does not need to be named krb5.conf explicitly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#krb5_config_gcs_uri GoogleDataprocMetastoreService#krb5_config_gcs_uri}
	Krb5ConfigGcsUri *string `field:"required" json:"krb5ConfigGcsUri" yaml:"krb5ConfigGcsUri"`
	// A Kerberos principal that exists in the both the keytab the KDC to authenticate as.
	//
	// A typical principal is of the form "primary/instance@REALM", but there is no exact format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#principal GoogleDataprocMetastoreService#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

