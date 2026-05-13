package vectorsearchcollection


type VectorSearchCollectionEncryptionSpec struct {
	// Resource name of the Cloud KMS key used to protect the resource.
	//
	// The Cloud KMS key must be in the same region as the resource. It must have
	// the format
	// 'projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vector_search_collection#crypto_key_name VectorSearchCollection#crypto_key_name}
	CryptoKeyName *string `field:"required" json:"cryptoKeyName" yaml:"cryptoKeyName"`
}

