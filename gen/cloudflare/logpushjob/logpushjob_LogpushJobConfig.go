package logpushjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogpushJobConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#dataset LogpushJob#dataset}.
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#destination_conf LogpushJob#destination_conf}.
	DestinationConf *string `field:"required" json:"destinationConf" yaml:"destinationConf"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#account_id LogpushJob#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#enabled LogpushJob#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#frequency LogpushJob#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#id LogpushJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#logpull_options LogpushJob#logpull_options}.
	LogpullOptions *string `field:"optional" json:"logpullOptions" yaml:"logpullOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#name LogpushJob#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#ownership_challenge LogpushJob#ownership_challenge}.
	OwnershipChallenge *string `field:"optional" json:"ownershipChallenge" yaml:"ownershipChallenge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/logpush_job#zone_id LogpushJob#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

