package googlememorystoreinstance


type GoogleMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstances struct {
	// The full resource path of the Nth instance in the format: projects/{project}/locations/{region}/instance/{instance-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#instance GoogleMemorystoreInstance#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
}

