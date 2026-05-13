package googlevectorsearchcollection


type GoogleVectorSearchCollectionVectorSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#field_name GoogleVectorSearchCollection#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// dense_vector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#dense_vector GoogleVectorSearchCollection#dense_vector}
	DenseVector *GoogleVectorSearchCollectionVectorSchemaDenseVector `field:"optional" json:"denseVector" yaml:"denseVector"`
	// sparse_vector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#sparse_vector GoogleVectorSearchCollection#sparse_vector}
	SparseVector *GoogleVectorSearchCollectionVectorSchemaSparseVector `field:"optional" json:"sparseVector" yaml:"sparseVector"`
}

