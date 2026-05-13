package googlevectorsearchcollection


type GoogleVectorSearchCollectionEncryptionSpec struct {
	// Resource name of the Cloud KMS key used to protect the resource.
	//
	// The Cloud KMS key must be in the same region as the resource. It must have
	// the format
	// 'projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#crypto_key_name GoogleVectorSearchCollection#crypto_key_name}
	CryptoKeyName *string `field:"required" json:"cryptoKeyName" yaml:"cryptoKeyName"`
}

