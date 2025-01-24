package googlealloydbinstance


type GoogleAlloydbInstanceObservabilityConfig struct {
	// Observability feature status for an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_alloydb_instance#enabled GoogleAlloydbInstance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Query string length. The default value is 10240. Any integer between 1024 and 100000 is considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_alloydb_instance#max_query_string_length GoogleAlloydbInstance#max_query_string_length}
	MaxQueryStringLength *float64 `field:"optional" json:"maxQueryStringLength" yaml:"maxQueryStringLength"`
	// Preserve comments in the query string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_alloydb_instance#preserve_comments GoogleAlloydbInstance#preserve_comments}
	PreserveComments interface{} `field:"optional" json:"preserveComments" yaml:"preserveComments"`
	// Number of query execution plans captured by Insights per minute for all queries combined.
	//
	// The default value is 5. Any integer between 0 and 200 is considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_alloydb_instance#query_plans_per_minute GoogleAlloydbInstance#query_plans_per_minute}
	QueryPlansPerMinute *float64 `field:"optional" json:"queryPlansPerMinute" yaml:"queryPlansPerMinute"`
	// Record application tags for an instance. This flag is turned "on" by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_alloydb_instance#record_application_tags GoogleAlloydbInstance#record_application_tags}
	RecordApplicationTags interface{} `field:"optional" json:"recordApplicationTags" yaml:"recordApplicationTags"`
	// Track actively running queries. If not set, default value is "off".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_alloydb_instance#track_active_queries GoogleAlloydbInstance#track_active_queries}
	TrackActiveQueries interface{} `field:"optional" json:"trackActiveQueries" yaml:"trackActiveQueries"`
	// Record wait events during query execution for an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_alloydb_instance#track_wait_events GoogleAlloydbInstance#track_wait_events}
	TrackWaitEvents interface{} `field:"optional" json:"trackWaitEvents" yaml:"trackWaitEvents"`
	// Record wait event types during query execution for an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_alloydb_instance#track_wait_event_types GoogleAlloydbInstance#track_wait_event_types}
	TrackWaitEventTypes interface{} `field:"optional" json:"trackWaitEventTypes" yaml:"trackWaitEventTypes"`
}

