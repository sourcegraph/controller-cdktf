package googlealloydbinstance


type GoogleAlloydbInstanceQueryInsightsConfig struct {
	// Number of query execution plans captured by Insights per minute for all queries combined.
	//
	// The default value is 5. Any integer between 0 and 20 is considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#query_plans_per_minute GoogleAlloydbInstance#query_plans_per_minute}
	QueryPlansPerMinute *float64 `field:"optional" json:"queryPlansPerMinute" yaml:"queryPlansPerMinute"`
	// Query string length. The default value is 1024. Any integer between 256 and 4500 is considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#query_string_length GoogleAlloydbInstance#query_string_length}
	QueryStringLength *float64 `field:"optional" json:"queryStringLength" yaml:"queryStringLength"`
	// Record application tags for an instance. This flag is turned "on" by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#record_application_tags GoogleAlloydbInstance#record_application_tags}
	RecordApplicationTags interface{} `field:"optional" json:"recordApplicationTags" yaml:"recordApplicationTags"`
	// Record client address for an instance. Client address is PII information. This flag is turned "on" by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#record_client_address GoogleAlloydbInstance#record_client_address}
	RecordClientAddress interface{} `field:"optional" json:"recordClientAddress" yaml:"recordClientAddress"`
}

