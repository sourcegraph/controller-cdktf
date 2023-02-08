package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceLogConfig struct {
	// Specifies whether to enable logging for traffic served by this service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_services_edge_cache_service#enable GoogleNetworkServicesEdgeCacheService#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Configures the sampling rate of requests, where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported. The default value is 1.0, and the value of the field must be in [0, 1].
	//
	// This field can only be specified if logging is enabled for this service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_services_edge_cache_service#sample_rate GoogleNetworkServicesEdgeCacheService#sample_rate}
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

