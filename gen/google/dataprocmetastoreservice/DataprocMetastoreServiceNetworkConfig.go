package dataprocmetastoreservice

type DataprocMetastoreServiceNetworkConfig struct {
	// consumers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataproc_metastore_service#consumers DataprocMetastoreService#consumers}
	Consumers any `field:"required" json:"consumers" yaml:"consumers"`
}
