package googlevectorsearchcollection


type GoogleVectorSearchCollectionVectorSchemaDenseVectorVertexEmbeddingConfig struct {
	// Required: ID of the embedding model to use. See https://cloud.google.com/vertex-ai/generative-ai/docs/learn/models#embeddings-models for the list of supported models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#model_id GoogleVectorSearchCollection#model_id}
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Possible values: RETRIEVAL_QUERY RETRIEVAL_DOCUMENT SEMANTIC_SIMILARITY CLASSIFICATION CLUSTERING QUESTION_ANSWERING FACT_VERIFICATION CODE_RETRIEVAL_QUERY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#task_type GoogleVectorSearchCollection#task_type}
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Required: Text template for the input to the model.
	//
	// The template must
	// contain one or more references to fields in the DataObject, e.g.:
	// "Movie Title: {title} ---- Movie Plot: {plot}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#text_template GoogleVectorSearchCollection#text_template}
	TextTemplate *string `field:"required" json:"textTemplate" yaml:"textTemplate"`
}

