package googlebigqueryanalyticshubdataexchange


type GoogleBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig struct {
	// dcr_exchange_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_analytics_hub_data_exchange#dcr_exchange_config GoogleBigqueryAnalyticsHubDataExchange#dcr_exchange_config}
	DcrExchangeConfig *GoogleBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfig `field:"optional" json:"dcrExchangeConfig" yaml:"dcrExchangeConfig"`
	// default_exchange_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_analytics_hub_data_exchange#default_exchange_config GoogleBigqueryAnalyticsHubDataExchange#default_exchange_config}
	DefaultExchangeConfig *GoogleBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfig `field:"optional" json:"defaultExchangeConfig" yaml:"defaultExchangeConfig"`
}

