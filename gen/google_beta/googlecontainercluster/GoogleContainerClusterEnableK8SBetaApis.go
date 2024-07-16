package googlecontainercluster


type GoogleContainerClusterEnableK8SBetaApis struct {
	// Enabled Kubernetes Beta APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enabled_apis GoogleContainerCluster#enabled_apis}
	EnabledApis *[]*string `field:"required" json:"enabledApis" yaml:"enabledApis"`
}

