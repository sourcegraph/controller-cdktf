package googlememorystoreinstance


type GoogleMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance struct {
	// The full resource path of the primary instance in the format: projects/{project}/locations/{region}/instances/{instance-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_memorystore_instance#instance GoogleMemorystoreInstance#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
}

