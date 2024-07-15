package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceNetworkConfig struct {
	// consumers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#consumers GoogleDataprocMetastoreService#consumers}
	Consumers interface{} `field:"required" json:"consumers" yaml:"consumers"`
	// Enables custom routes to be imported and exported for the Dataproc Metastore service's peered VPC network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#custom_routes_enabled GoogleDataprocMetastoreService#custom_routes_enabled}
	CustomRoutesEnabled interface{} `field:"optional" json:"customRoutesEnabled" yaml:"customRoutesEnabled"`
}

