package googlevectorsearchcollection


type GoogleVectorSearchCollectionVectorSchemaDenseVector struct {
	// Dimensionality of the vector field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#dimensions GoogleVectorSearchCollection#dimensions}
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// vertex_embedding_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#vertex_embedding_config GoogleVectorSearchCollection#vertex_embedding_config}
	VertexEmbeddingConfig *GoogleVectorSearchCollectionVectorSchemaDenseVectorVertexEmbeddingConfig `field:"optional" json:"vertexEmbeddingConfig" yaml:"vertexEmbeddingConfig"`
}

