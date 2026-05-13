package dataplexdatascan


type DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger struct {
	// The score range is in [0,100].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#score_threshold DataplexDatascan#score_threshold}
	ScoreThreshold *float64 `field:"optional" json:"scoreThreshold" yaml:"scoreThreshold"`
}

