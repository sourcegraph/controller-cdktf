package googlecomputebackendservice


type GoogleComputeBackendServiceBackend struct {
	// The fully-qualified URL of an Instance Group or Network Endpoint Group resource.
	//
	// In case of instance group this defines the list
	// of instances that serve traffic. Member virtual machine
	// instances from each instance group must live in the same zone as
	// the instance group itself. No two backends in a backend service
	// are allowed to use same Instance Group resource.
	//
	// For Network Endpoint Groups this defines list of endpoints. All
	// endpoints of Network Endpoint Group must be hosted on instances
	// located in the same zone as the Network Endpoint Group.
	//
	// Backend services cannot mix Instance Group and
	// Network Endpoint Group backends.
	//
	// Note that you must specify an Instance Group or Network Endpoint
	// Group resource using the fully-qualified URL, rather than a
	// partial URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#group GoogleComputeBackendService#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// Specifies the balancing mode for this backend.
	//
	// For global HTTP(S) or TCP/SSL load balancing, the default is
	// UTILIZATION. Valid values are UTILIZATION, RATE (for HTTP(S)),
	// CUSTOM_METRICS (for HTTP(s)) and CONNECTION (for TCP/SSL).
	//
	// See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
	// for an explanation of load balancing modes. Default value: "UTILIZATION" Possible values: ["UTILIZATION", "RATE", "CONNECTION", "CUSTOM_METRICS", "IN_FLIGHT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#balancing_mode GoogleComputeBackendService#balancing_mode}
	BalancingMode *string `field:"optional" json:"balancingMode" yaml:"balancingMode"`
	// A multiplier applied to the group's maximum servicing capacity (based on UTILIZATION, RATE or CONNECTION).
	//
	// Default value is 1, which means the group will serve up to 100%
	// of its configured capacity (depending on balancingMode). A
	// setting of 0 means the group is completely drained, offering
	// 0% of its available Capacity. Valid range is [0.0,1.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#capacity_scaler GoogleComputeBackendService#capacity_scaler}
	CapacityScaler *float64 `field:"optional" json:"capacityScaler" yaml:"capacityScaler"`
	// custom_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#custom_metrics GoogleComputeBackendService#custom_metrics}
	CustomMetrics interface{} `field:"optional" json:"customMetrics" yaml:"customMetrics"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#description GoogleComputeBackendService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The max number of simultaneous connections for the group. Can be used with either CONNECTION or UTILIZATION balancing modes.
	//
	// For CONNECTION mode, either maxConnections or one
	// of maxConnectionsPerInstance or maxConnectionsPerEndpoint,
	// as appropriate for group type, must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_connections GoogleComputeBackendService#max_connections}
	MaxConnections *float64 `field:"optional" json:"maxConnections" yaml:"maxConnections"`
	// The max number of simultaneous connections that a single backend network endpoint can handle.
	//
	// This is used to calculate the
	// capacity of the group. Can be used in either CONNECTION or
	// UTILIZATION balancing modes.
	//
	// For CONNECTION mode, either
	// maxConnections or maxConnectionsPerEndpoint must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_connections_per_endpoint GoogleComputeBackendService#max_connections_per_endpoint}
	MaxConnectionsPerEndpoint *float64 `field:"optional" json:"maxConnectionsPerEndpoint" yaml:"maxConnectionsPerEndpoint"`
	// The max number of simultaneous connections that a single backend instance can handle.
	//
	// This is used to calculate the
	// capacity of the group. Can be used in either CONNECTION or
	// UTILIZATION balancing modes.
	//
	// For CONNECTION mode, either maxConnections or
	// maxConnectionsPerInstance must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_connections_per_instance GoogleComputeBackendService#max_connections_per_instance}
	MaxConnectionsPerInstance *float64 `field:"optional" json:"maxConnectionsPerInstance" yaml:"maxConnectionsPerInstance"`
	// Defines a maximum number of in-flight requests for the whole NEG or instance group.
	//
	// Not available if backend's balancingMode is RATE
	// or CONNECTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_in_flight_requests GoogleComputeBackendService#max_in_flight_requests}
	MaxInFlightRequests *float64 `field:"optional" json:"maxInFlightRequests" yaml:"maxInFlightRequests"`
	// Defines a maximum number of in-flight requests for a single endpoint. Not available if backend's balancingMode is RATE or CONNECTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_in_flight_requests_per_endpoint GoogleComputeBackendService#max_in_flight_requests_per_endpoint}
	MaxInFlightRequestsPerEndpoint *float64 `field:"optional" json:"maxInFlightRequestsPerEndpoint" yaml:"maxInFlightRequestsPerEndpoint"`
	// Defines a maximum number of in-flight requests for a single VM. Not available if backend's balancingMode is RATE or CONNECTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_in_flight_requests_per_instance GoogleComputeBackendService#max_in_flight_requests_per_instance}
	MaxInFlightRequestsPerInstance *float64 `field:"optional" json:"maxInFlightRequestsPerInstance" yaml:"maxInFlightRequestsPerInstance"`
	// The max requests per second (RPS) of the group.
	//
	// Can be used with either RATE or UTILIZATION balancing modes,
	// but required if RATE mode. For RATE mode, either maxRate or one
	// of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
	// group type, must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_rate GoogleComputeBackendService#max_rate}
	MaxRate *float64 `field:"optional" json:"maxRate" yaml:"maxRate"`
	// The max requests per second (RPS) that a single backend network endpoint can handle.
	//
	// This is used to calculate the capacity of
	// the group. Can be used in either balancing mode. For RATE mode,
	// either maxRate or maxRatePerEndpoint must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_rate_per_endpoint GoogleComputeBackendService#max_rate_per_endpoint}
	MaxRatePerEndpoint *float64 `field:"optional" json:"maxRatePerEndpoint" yaml:"maxRatePerEndpoint"`
	// The max requests per second (RPS) that a single backend instance can handle.
	//
	// This is used to calculate the capacity of
	// the group. Can be used in either balancing mode. For RATE mode,
	// either maxRate or maxRatePerInstance must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_rate_per_instance GoogleComputeBackendService#max_rate_per_instance}
	MaxRatePerInstance *float64 `field:"optional" json:"maxRatePerInstance" yaml:"maxRatePerInstance"`
	// Used when balancingMode is UTILIZATION. This ratio defines the CPU utilization target for the group. Valid range is [0.0, 1.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#max_utilization GoogleComputeBackendService#max_utilization}
	MaxUtilization *float64 `field:"optional" json:"maxUtilization" yaml:"maxUtilization"`
	// This field indicates whether this backend should be fully utilized before sending traffic to backends with default preference.
	//
	// This field cannot be set when loadBalancingScheme is set to 'EXTERNAL'. The possible values are:
	//   - PREFERRED: Backends with this preference level will be filled up to their capacity limits first,
	//     based on RTT.
	//   - DEFAULT: If preferred backends don't have enough capacity, backends in this layer would be used and
	//     traffic would be assigned based on the load balancing algorithm you use. This is the default Possible values: ["PREFERRED", "DEFAULT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#preference GoogleComputeBackendService#preference}
	Preference *string `field:"optional" json:"preference" yaml:"preference"`
	// This field specifies how long a connection should be kept alive for: - LONG: Most of the requests are expected to take more than multiple   seconds to finish.
	//
	// - SHORT: Most requests are expected to finish with a sub-second latency. Possible values: ["LONG", "SHORT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#traffic_duration GoogleComputeBackendService#traffic_duration}
	TrafficDuration *string `field:"optional" json:"trafficDuration" yaml:"trafficDuration"`
}

