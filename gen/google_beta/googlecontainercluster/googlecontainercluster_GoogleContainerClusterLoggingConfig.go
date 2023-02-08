package googlecontainercluster


type GoogleContainerClusterLoggingConfig struct {
	// GKE components exposing logs. Valid values include SYSTEM_COMPONENTS and WORKLOADS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enable_components GoogleContainerCluster#enable_components}
	EnableComponents *[]*string `field:"required" json:"enableComponents" yaml:"enableComponents"`
}

