package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab struct {
	// The relative resource name of a Secret Manager secret version, in the following form:.
	//
	// "projects/{projectNumber}/secrets/{secret_id}/versions/{version_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#cloud_secret GoogleDataprocMetastoreService#cloud_secret}
	CloudSecret *string `field:"required" json:"cloudSecret" yaml:"cloudSecret"`
}

