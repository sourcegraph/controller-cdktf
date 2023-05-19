package notificationpolicy


type NotificationPolicyFilters struct {
	// State of the pool to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#enabled NotificationPolicy#enabled}
	Enabled *[]*string `field:"optional" json:"enabled" yaml:"enabled"`
	// Source configuration to alert on for pool or origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#event_source NotificationPolicy#event_source}
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// Stream event type to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#event_type NotificationPolicy#event_type}
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// Identifier health check. Required when using `filters.0.status`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#health_check_id NotificationPolicy#health_check_id}
	HealthCheckId *[]*string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Stream input id to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#input_id NotificationPolicy#input_id}
	InputId *[]*string `field:"optional" json:"inputId" yaml:"inputId"`
	// A numerical limit. Example: `100`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#limit NotificationPolicy#limit}
	Limit *[]*string `field:"optional" json:"limit" yaml:"limit"`
	// Health status to alert on for pool or origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#new_health NotificationPolicy#new_health}
	NewHealth *[]*string `field:"optional" json:"newHealth" yaml:"newHealth"`
	// Packets per second threshold for dos alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#packets_per_second NotificationPolicy#packets_per_second}
	PacketsPerSecond *[]*string `field:"optional" json:"packetsPerSecond" yaml:"packetsPerSecond"`
	// Load balancer pool identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#pool_id NotificationPolicy#pool_id}
	PoolId *[]*string `field:"optional" json:"poolId" yaml:"poolId"`
	// Product name. Available values: `worker_requests`, `worker_durable_objects_requests`, `worker_durable_objects_duration`, `worker_durable_objects_data_transfer`, `worker_durable_objects_stored_data`, `worker_durable_objects_storage_deletes`, `worker_durable_objects_storage_writes`, `worker_durable_objects_storage_reads`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#product NotificationPolicy#product}
	Product *[]*string `field:"optional" json:"product" yaml:"product"`
	// Protocol to alert on for dos.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#protocol NotificationPolicy#protocol}
	Protocol *[]*string `field:"optional" json:"protocol" yaml:"protocol"`
	// Requests per second threshold for dos alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#requests_per_second NotificationPolicy#requests_per_second}
	RequestsPerSecond *[]*string `field:"optional" json:"requestsPerSecond" yaml:"requestsPerSecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#services NotificationPolicy#services}.
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// A numerical limit. Example: `99.9`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#slo NotificationPolicy#slo}
	Slo *[]*string `field:"optional" json:"slo" yaml:"slo"`
	// Status to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#status NotificationPolicy#status}
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// Target host to alert on for dos.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#target_host NotificationPolicy#target_host}
	TargetHost *[]*string `field:"optional" json:"targetHost" yaml:"targetHost"`
	// Target domain to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#target_zone_name NotificationPolicy#target_zone_name}
	TargetZoneName *[]*string `field:"optional" json:"targetZoneName" yaml:"targetZoneName"`
	// A list of zone identifiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy#zones NotificationPolicy#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

