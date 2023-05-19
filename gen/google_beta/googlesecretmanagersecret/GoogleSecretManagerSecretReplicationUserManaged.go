package googlesecretmanagersecret


type GoogleSecretManagerSecretReplicationUserManaged struct {
	// replicas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_secret_manager_secret#replicas GoogleSecretManagerSecret#replicas}
	Replicas interface{} `field:"required" json:"replicas" yaml:"replicas"`
}

