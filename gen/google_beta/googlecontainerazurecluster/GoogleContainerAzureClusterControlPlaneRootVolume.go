package googlecontainerazurecluster


type GoogleContainerAzureClusterControlPlaneRootVolume struct {
	// Optional.
	//
	// The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_cluster#size_gib GoogleContainerAzureCluster#size_gib}
	SizeGib *float64 `field:"optional" json:"sizeGib" yaml:"sizeGib"`
}

