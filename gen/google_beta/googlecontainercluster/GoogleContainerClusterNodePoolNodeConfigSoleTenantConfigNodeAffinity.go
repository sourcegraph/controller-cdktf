package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigSoleTenantConfigNodeAffinity struct {
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#key GoogleContainerCluster#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#operator GoogleContainerCluster#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#values GoogleContainerCluster#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

