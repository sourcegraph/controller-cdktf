package googleedgecontainercluster


type GoogleEdgecontainerClusterMaintenancePolicyMaintenanceExclusions struct {
	// A unique (per cluster) id for the window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#id GoogleEdgecontainerCluster#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#window GoogleEdgecontainerCluster#window}
	Window *GoogleEdgecontainerClusterMaintenancePolicyMaintenanceExclusionsWindow `field:"optional" json:"window" yaml:"window"`
}

