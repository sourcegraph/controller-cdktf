package googleedgecontainercluster


type GoogleEdgecontainerClusterControlPlaneEncryption struct {
	// The Cloud KMS CryptoKey e.g. projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey} to use for protecting control plane disks. If not specified, a Google-managed key will be used instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#kms_key GoogleEdgecontainerCluster#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

