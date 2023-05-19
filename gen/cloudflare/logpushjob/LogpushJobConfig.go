package logpushjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogpushJobConfig struct {
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
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	//
	// Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/reference/logpush-api-configuration#destination). Available values: `access_requests`, `firewall_events`, `http_requests`, `spectrum_events`, `nel_reports`, `audit_logs`, `gateway_dns`, `gateway_http`, `gateway_network`, `dns_logs`, `network_analytics_logs`, `workers_trace_events`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#dataset LogpushJob#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	//
	// Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/reference/logpush-api-configuration#destination).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#destination_conf LogpushJob#destination_conf}
	DestinationConf *string `field:"required" json:"destinationConf" yaml:"destinationConf"`
	// The account identifier to target for the resource. Must provide only one of `account_id`, `zone_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#account_id LogpushJob#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Whether to enable the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#enabled LogpushJob#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Use filters to select the events to include and/or remove from your logs. For more information, refer to [Filters](https://developers.cloudflare.com/logs/reference/logpush-api-configuration/filters/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#filter LogpushJob#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// A higher frequency will result in logs being pushed on faster with smaller files.
	//
	// `low` frequency will push logs less often with larger files. Available values: `high`, `low`. Defaults to `high`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#frequency LogpushJob#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#id LogpushJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The kind of logpush job to create. Available values: `edge`, `instant-logs`, `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#kind LogpushJob#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Configuration string for the Logshare API. It specifies things like requested fields and timestamp formats. See [Logpush options documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#options).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#logpull_options LogpushJob#logpull_options}
	LogpullOptions *string `field:"optional" json:"logpullOptions" yaml:"logpullOptions"`
	// The name of the logpush job to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#name LogpushJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Ownership challenge token to prove destination ownership, required when destination is Amazon S3, Google Cloud Storage, Microsoft Azure or Sumo Logic.
	//
	// See [Developer documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#usage).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#ownership_challenge LogpushJob#ownership_challenge}
	OwnershipChallenge *string `field:"optional" json:"ownershipChallenge" yaml:"ownershipChallenge"`
	// The zone identifier to target for the resource. Must provide only one of `account_id`, `zone_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/logpush_job#zone_id LogpushJob#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

