package googlecloudrundomainmapping


type GoogleCloudRunDomainMappingMetadata struct {
	// In Cloud Run the namespace must be equal to either the project ID or project number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_domain_mapping#namespace GoogleCloudRunDomainMapping#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Annotations is a key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// More
	// info: http://kubernetes.io/docs/user-guide/annotations
	//
	// *Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_domain_mapping#annotations GoogleCloudRunDomainMapping#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers
	// and routes.
	// More info: http://kubernetes.io/docs/user-guide/labels
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_domain_mapping#labels GoogleCloudRunDomainMapping#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

