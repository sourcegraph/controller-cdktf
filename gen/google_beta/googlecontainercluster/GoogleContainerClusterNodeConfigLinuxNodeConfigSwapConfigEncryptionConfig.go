package googlecontainercluster


type GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig struct {
	// If true, swap space will not be encrypted. Defaults to false (encrypted).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#disabled GoogleContainerCluster#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

