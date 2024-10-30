package logstream


type LogStreamSettings struct {
	// AWS account ID. Required only for 'aws_eventbridge' type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#account_id LogStream#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Edition of the Splunk Cloud instance. Could be one of: 'aws', 'aws_govcloud', 'gcp'. Required only for 'splunk_cloud_logstreaming' type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#edition LogStream#edition}
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// An alphanumeric name (no spaces) to identify this event source in AWS EventBridge. Required only for 'aws_eventbridge' type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#event_source_name LogStream#event_source_name}
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
	// The domain name for Splunk Cloud instance.
	//
	// Don't include http or https in the string. For example: 'acme.splunkcloud.com'. Required only for 'splunk_cloud_logstreaming' type
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#host LogStream#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The destination AWS region where event source is located. Required only for 'aws_eventbridge' type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#region LogStream#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The HEC token for your Splunk Cloud HTTP Event Collector. Required only for 'splunk_cloud_logstreaming' type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#token LogStream#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

