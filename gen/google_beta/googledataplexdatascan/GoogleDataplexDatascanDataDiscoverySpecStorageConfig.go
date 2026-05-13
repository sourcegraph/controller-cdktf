package googledataplexdatascan


type GoogleDataplexDatascanDataDiscoverySpecStorageConfig struct {
	// csv_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#csv_options GoogleDataplexDatascan#csv_options}
	CsvOptions *GoogleDataplexDatascanDataDiscoverySpecStorageConfigCsvOptions `field:"optional" json:"csvOptions" yaml:"csvOptions"`
	// Defines the data to exclude during discovery.
	//
	// Provide a list of patterns that identify the data to exclude. For Cloud Storage bucket assets, these patterns are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these patterns are interpreted as patterns to match table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#exclude_patterns GoogleDataplexDatascan#exclude_patterns}
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Defines the data to include during discovery when only a subset of the data should be considered.
	//
	// Provide a list of patterns that identify the data to include. For Cloud Storage bucket assets, these patterns are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these patterns are interpreted as patterns to match table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#include_patterns GoogleDataplexDatascan#include_patterns}
	IncludePatterns *[]*string `field:"optional" json:"includePatterns" yaml:"includePatterns"`
	// json_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#json_options GoogleDataplexDatascan#json_options}
	JsonOptions *GoogleDataplexDatascanDataDiscoverySpecStorageConfigJsonOptions `field:"optional" json:"jsonOptions" yaml:"jsonOptions"`
}

