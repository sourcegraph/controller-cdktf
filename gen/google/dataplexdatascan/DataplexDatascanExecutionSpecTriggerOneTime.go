package dataplexdatascan


type DataplexDatascanExecutionSpecTriggerOneTime struct {
	// Time to live for the DataScan and its results after the one-time run completes.
	//
	// Accepts a string with a unit suffix 's' (e.g., '7200s'). Default is 24 hours. Ranges between 0 and 31536000 seconds (1 year).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#ttl_after_scan_completion DataplexDatascan#ttl_after_scan_completion}
	TtlAfterScanCompletion *string `field:"optional" json:"ttlAfterScanCompletion" yaml:"ttlAfterScanCompletion"`
}

