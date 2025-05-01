package googlediscoveryenginedatastore


type GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig struct {
	// digital_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_discovery_engine_data_store#digital_parsing_config GoogleDiscoveryEngineDataStore#digital_parsing_config}
	DigitalParsingConfig *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfig `field:"optional" json:"digitalParsingConfig" yaml:"digitalParsingConfig"`
	// layout_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_discovery_engine_data_store#layout_parsing_config GoogleDiscoveryEngineDataStore#layout_parsing_config}
	LayoutParsingConfig *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfig `field:"optional" json:"layoutParsingConfig" yaml:"layoutParsingConfig"`
	// ocr_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_discovery_engine_data_store#ocr_parsing_config GoogleDiscoveryEngineDataStore#ocr_parsing_config}
	OcrParsingConfig *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfig `field:"optional" json:"ocrParsingConfig" yaml:"ocrParsingConfig"`
}

