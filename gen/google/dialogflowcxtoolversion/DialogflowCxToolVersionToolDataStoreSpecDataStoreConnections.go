package dialogflowcxtoolversion


type DialogflowCxToolVersionToolDataStoreSpecDataStoreConnections struct {
	// The full name of the referenced data store. Formats: projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore} projects/{project}/locations/{location}/dataStores/{dataStore}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#data_store DialogflowCxToolVersion#data_store}
	DataStore *string `field:"optional" json:"dataStore" yaml:"dataStore"`
	// The type of the connected data store. See [DataStoreType](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/DataStoreConnection#datastoretype) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#data_store_type DialogflowCxToolVersion#data_store_type}
	DataStoreType *string `field:"optional" json:"dataStoreType" yaml:"dataStoreType"`
	// The document processing mode for the data store connection.
	//
	// Should only be set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is considered as DOCUMENTS, as this is the legacy mode.
	// See [DocumentProcessingMode](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/DataStoreConnection#documentprocessingmode) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#document_processing_mode DialogflowCxToolVersion#document_processing_mode}
	DocumentProcessingMode *string `field:"optional" json:"documentProcessingMode" yaml:"documentProcessingMode"`
}

