package vectorsearchcollection


type VectorSearchCollectionVectorSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vector_search_collection#field_name VectorSearchCollection#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// dense_vector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vector_search_collection#dense_vector VectorSearchCollection#dense_vector}
	DenseVector *VectorSearchCollectionVectorSchemaDenseVector `field:"optional" json:"denseVector" yaml:"denseVector"`
	// sparse_vector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vector_search_collection#sparse_vector VectorSearchCollection#sparse_vector}
	SparseVector *VectorSearchCollectionVectorSchemaSparseVector `field:"optional" json:"sparseVector" yaml:"sparseVector"`
}

