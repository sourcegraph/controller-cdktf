package googlefilestoreinstance


type GoogleFilestoreInstanceInitialReplicationReplicas struct {
	// The peer instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_filestore_instance#peer_instance GoogleFilestoreInstance#peer_instance}
	PeerInstance *string `field:"required" json:"peerInstance" yaml:"peerInstance"`
}

