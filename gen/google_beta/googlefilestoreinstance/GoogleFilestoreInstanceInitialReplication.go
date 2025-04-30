package googlefilestoreinstance


type GoogleFilestoreInstanceInitialReplication struct {
	// replicas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_filestore_instance#replicas GoogleFilestoreInstance#replicas}
	Replicas interface{} `field:"optional" json:"replicas" yaml:"replicas"`
	// The replication role. Default value: "STANDBY" Possible values: ["ROLE_UNSPECIFIED", "ACTIVE", "STANDBY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_filestore_instance#role GoogleFilestoreInstance#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

