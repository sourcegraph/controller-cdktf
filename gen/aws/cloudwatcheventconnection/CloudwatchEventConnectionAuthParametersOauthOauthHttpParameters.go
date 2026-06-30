package cloudwatcheventconnection

type CloudwatchEventConnectionAuthParametersOauthOauthHttpParameters struct {
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/cloudwatch_event_connection#body CloudwatchEventConnection#body}
	Body any `field:"optional" json:"body" yaml:"body"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/cloudwatch_event_connection#header CloudwatchEventConnection#header}
	Header any `field:"optional" json:"header" yaml:"header"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/cloudwatch_event_connection#query_string CloudwatchEventConnection#query_string}
	QueryString any `field:"optional" json:"queryString" yaml:"queryString"`
}
