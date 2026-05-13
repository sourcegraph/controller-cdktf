package googlecontainercluster


type GoogleContainerClusterManagedOpentelemetryConfig struct {
	// The scope of the Managed OpenTelemetry pipeline. Available options include SCOPE_UNSPECIFIED, NONE, and COLLECTION_AND_INSTRUMENTATION_COMPONENTS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#scope GoogleContainerCluster#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

