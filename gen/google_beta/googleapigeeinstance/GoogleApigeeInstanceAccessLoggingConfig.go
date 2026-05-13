package googleapigeeinstance


type GoogleApigeeInstanceAccessLoggingConfig struct {
	// Boolean flag that specifies whether the customer access log feature is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_instance#enabled GoogleApigeeInstance#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Ship the access log entries that match the statusCode defined in the filter.
	//
	// The statusCode is the only expected/supported filter field. (Ex: statusCode)
	// The filter will parse it to the Common Expression Language semantics for expression
	// evaluation to build the filter condition. (Ex: "filter": statusCode >= 200 && statusCode < 300 )
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_instance#filter GoogleApigeeInstance#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

