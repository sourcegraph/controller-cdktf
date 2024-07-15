package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceCircuitBreakers struct {
	// connect_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#connect_timeout GoogleComputeRegionBackendService#connect_timeout}
	ConnectTimeout *GoogleComputeRegionBackendServiceCircuitBreakersConnectTimeout `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// The maximum number of connections to the backend cluster. Defaults to 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#max_connections GoogleComputeRegionBackendService#max_connections}
	MaxConnections *float64 `field:"optional" json:"maxConnections" yaml:"maxConnections"`
	// The maximum number of pending requests to the backend cluster. Defaults to 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#max_pending_requests GoogleComputeRegionBackendService#max_pending_requests}
	MaxPendingRequests *float64 `field:"optional" json:"maxPendingRequests" yaml:"maxPendingRequests"`
	// The maximum number of parallel requests to the backend cluster. Defaults to 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#max_requests GoogleComputeRegionBackendService#max_requests}
	MaxRequests *float64 `field:"optional" json:"maxRequests" yaml:"maxRequests"`
	// Maximum requests for a single backend connection.
	//
	// This parameter
	// is respected by both the HTTP/1.1 and HTTP/2 implementations. If
	// not specified, there is no limit. Setting this parameter to 1
	// will effectively disable keep alive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#max_requests_per_connection GoogleComputeRegionBackendService#max_requests_per_connection}
	MaxRequestsPerConnection *float64 `field:"optional" json:"maxRequestsPerConnection" yaml:"maxRequestsPerConnection"`
	// The maximum number of parallel retries to the backend cluster. Defaults to 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#max_retries GoogleComputeRegionBackendService#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
}

