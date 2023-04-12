package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab struct {
	// The relative resource name of a Secret Manager secret version, in the following form:.
	//
	// "projects/{projectNumber}/secrets/{secret_id}/versions/{version_id}".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_service#cloud_secret GoogleDataprocMetastoreService#cloud_secret}
	CloudSecret *string `field:"required" json:"cloudSecret" yaml:"cloudSecret"`
}

