package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceConnectionTrackingPolicy struct {
	// Specifies connection persistence when backends are unhealthy.
	//
	// If set to 'DEFAULT_FOR_PROTOCOL', the existing connections persist on
	// unhealthy backends only for connection-oriented protocols (TCP and SCTP)
	// and only if the Tracking Mode is PER_CONNECTION (default tracking mode)
	// or the Session Affinity is configured for 5-tuple. They do not persist
	// for UDP.
	//
	// If set to 'NEVER_PERSIST', after a backend becomes unhealthy, the existing
	// connections on the unhealthy backend are never persisted on the unhealthy
	// backend. They are always diverted to newly selected healthy backends
	// (unless all backends are unhealthy).
	//
	// If set to 'ALWAYS_PERSIST', existing connections always persist on
	// unhealthy backends regardless of protocol and session affinity. It is
	// generally not recommended to use this mode overriding the default. Default value: "DEFAULT_FOR_PROTOCOL" Possible values: ["DEFAULT_FOR_PROTOCOL", "NEVER_PERSIST", "ALWAYS_PERSIST"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#connection_persistence_on_unhealthy_backends GoogleComputeRegionBackendService#connection_persistence_on_unhealthy_backends}
	ConnectionPersistenceOnUnhealthyBackends *string `field:"optional" json:"connectionPersistenceOnUnhealthyBackends" yaml:"connectionPersistenceOnUnhealthyBackends"`
	// Enable Strong Session Affinity for Network Load Balancing. This option is not available publicly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#enable_strong_affinity GoogleComputeRegionBackendService#enable_strong_affinity}
	EnableStrongAffinity interface{} `field:"optional" json:"enableStrongAffinity" yaml:"enableStrongAffinity"`
	// Specifies how long to keep a Connection Tracking entry while there is no matching traffic (in seconds).
	//
	// For L4 ILB the minimum(default) is 10 minutes and maximum is 16 hours.
	//
	// For NLB the minimum(default) is 60 seconds and the maximum is 16 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#idle_timeout_sec GoogleComputeRegionBackendService#idle_timeout_sec}
	IdleTimeoutSec *float64 `field:"optional" json:"idleTimeoutSec" yaml:"idleTimeoutSec"`
	// Specifies the key used for connection tracking.
	//
	// There are two options:
	// 'PER_CONNECTION': The Connection Tracking is performed as per the
	// Connection Key (default Hash Method) for the specific protocol.
	//
	// 'PER_SESSION': The Connection Tracking is performed as per the
	// configured Session Affinity. It matches the configured Session Affinity. Default value: "PER_CONNECTION" Possible values: ["PER_CONNECTION", "PER_SESSION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#tracking_mode GoogleComputeRegionBackendService#tracking_mode}
	TrackingMode *string `field:"optional" json:"trackingMode" yaml:"trackingMode"`
}

