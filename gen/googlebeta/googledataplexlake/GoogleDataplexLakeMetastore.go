package googledataplexlake


type GoogleDataplexLakeMetastore struct {
	// Optional. A relative reference to the Dataproc Metastore (https://cloud.google.com/dataproc-metastore/docs) service associated with the lake: `projects/{project_id}/locations/{location_id}/services/{service_id}`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_lake#service GoogleDataplexLake#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

