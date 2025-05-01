package googlediscoveryenginedatastore


type GoogleDiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfigLayoutBasedChunkingConfig struct {
	// The token size limit for each chunk. Supported values: 100-500 (inclusive). Default value: 500.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_discovery_engine_data_store#chunk_size GoogleDiscoveryEngineDataStore#chunk_size}
	ChunkSize *float64 `field:"optional" json:"chunkSize" yaml:"chunkSize"`
	// Whether to include appending different levels of headings to chunks from the middle of the document to prevent context loss.
	//
	// Default value: False.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_discovery_engine_data_store#include_ancestor_headings GoogleDiscoveryEngineDataStore#include_ancestor_headings}
	IncludeAncestorHeadings interface{} `field:"optional" json:"includeAncestorHeadings" yaml:"includeAncestorHeadings"`
}

