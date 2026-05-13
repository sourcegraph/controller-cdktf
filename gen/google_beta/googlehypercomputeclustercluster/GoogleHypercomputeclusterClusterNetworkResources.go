package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterNetworkResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#id GoogleHypercomputeclusterCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#config GoogleHypercomputeclusterCluster#config}
	Config *GoogleHypercomputeclusterClusterNetworkResourcesConfig `field:"optional" json:"config" yaml:"config"`
}

