package googlediscoveryenginedatastore


type GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfig struct {
	// If true, the LLM based annotation is added to the image during parsing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_store#enable_image_annotation GoogleDiscoveryEngineDataStore#enable_image_annotation}
	EnableImageAnnotation interface{} `field:"optional" json:"enableImageAnnotation" yaml:"enableImageAnnotation"`
	// If true, the LLM based annotation is added to the table during parsing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_store#enable_table_annotation GoogleDiscoveryEngineDataStore#enable_table_annotation}
	EnableTableAnnotation interface{} `field:"optional" json:"enableTableAnnotation" yaml:"enableTableAnnotation"`
	// List of HTML classes to exclude from the parsed content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_store#exclude_html_classes GoogleDiscoveryEngineDataStore#exclude_html_classes}
	ExcludeHtmlClasses *[]*string `field:"optional" json:"excludeHtmlClasses" yaml:"excludeHtmlClasses"`
	// List of HTML elements to exclude from the parsed content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_store#exclude_html_elements GoogleDiscoveryEngineDataStore#exclude_html_elements}
	ExcludeHtmlElements *[]*string `field:"optional" json:"excludeHtmlElements" yaml:"excludeHtmlElements"`
	// List of HTML ids to exclude from the parsed content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_store#exclude_html_ids GoogleDiscoveryEngineDataStore#exclude_html_ids}
	ExcludeHtmlIds *[]*string `field:"optional" json:"excludeHtmlIds" yaml:"excludeHtmlIds"`
	// Contains the required structure types to extract from the document. Supported values: 'shareholder-structure'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_store#structured_content_types GoogleDiscoveryEngineDataStore#structured_content_types}
	StructuredContentTypes *[]*string `field:"optional" json:"structuredContentTypes" yaml:"structuredContentTypes"`
}

