package googledataprocmetastorefederation


type GoogleDataprocMetastoreFederationBackendMetastores struct {
	// The type of the backend metastore. Possible values: ["METASTORE_TYPE_UNSPECIFIED", "DATAPROC_METASTORE"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_federation#metastore_type GoogleDataprocMetastoreFederation#metastore_type}
	MetastoreType *string `field:"required" json:"metastoreType" yaml:"metastoreType"`
	// The relative resource name of the metastore that is being federated.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_federation#name GoogleDataprocMetastoreFederation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_federation#rank GoogleDataprocMetastoreFederation#rank}.
	Rank *string `field:"required" json:"rank" yaml:"rank"`
}

