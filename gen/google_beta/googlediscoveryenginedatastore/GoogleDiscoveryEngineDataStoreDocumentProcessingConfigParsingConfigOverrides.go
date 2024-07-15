package googlediscoveryenginedatastore


type GoogleDiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_discovery_engine_data_store#file_type GoogleDiscoveryEngineDataStore#file_type}.
	FileType *string `field:"required" json:"fileType" yaml:"fileType"`
	// digital_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_discovery_engine_data_store#digital_parsing_config GoogleDiscoveryEngineDataStore#digital_parsing_config}
	DigitalParsingConfig *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesDigitalParsingConfig `field:"optional" json:"digitalParsingConfig" yaml:"digitalParsingConfig"`
	// ocr_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_discovery_engine_data_store#ocr_parsing_config GoogleDiscoveryEngineDataStore#ocr_parsing_config}
	OcrParsingConfig *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesOcrParsingConfig `field:"optional" json:"ocrParsingConfig" yaml:"ocrParsingConfig"`
}

