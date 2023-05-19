package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceNetworkConfig struct {
	// consumers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataproc_metastore_service#consumers GoogleDataprocMetastoreService#consumers}
	Consumers interface{} `field:"required" json:"consumers" yaml:"consumers"`
}

