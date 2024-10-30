package logstream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogStreamConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Unique name for the Log Stream object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#name LogStream#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Streaming provider used - 'aws_eventbridge' or 'splunk_cloud_logstreaming'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#type LogStream#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#settings LogStream#settings}
	Settings *LogStreamSettings `field:"optional" json:"settings" yaml:"settings"`
	// Stream status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/log_stream#status LogStream#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

