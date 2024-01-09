package directdatadog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectDatadogConfig struct {
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
	// Unique name of the resource, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#name DirectDatadog#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the Nobl9 project the resource sits in, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#project DirectDatadog#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// `com` or `eu`, Datadog SaaS instance, which corresponds to one of Datadog's two locations (https://www.datadoghq.com/ in the U.S. or https://datadoghq.eu/ in the European Union).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#site DirectDatadog#site}
	Site *string `field:"required" json:"site" yaml:"site"`
	// Source of Metrics and/or Services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#source_of DirectDatadog#source_of}
	SourceOf *[]*string `field:"required" json:"sourceOf" yaml:"sourceOf"`
	// [required] | Datadog API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#api_key DirectDatadog#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// [required] | Datadog Application Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#application_key DirectDatadog#application_key}
	ApplicationKey *string `field:"optional" json:"applicationKey" yaml:"applicationKey"`
	// Optional description of the resource.
	//
	// Here, you can add details about who is responsible for the integration (team/owner) or the purpose of creating it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#description DirectDatadog#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-friendly display name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#display_name DirectDatadog#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// historical_data_retrieval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#historical_data_retrieval DirectDatadog#historical_data_retrieval}
	HistoricalDataRetrieval *DirectDatadogHistoricalDataRetrieval `field:"optional" json:"historicalDataRetrieval" yaml:"historicalDataRetrieval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#id DirectDatadog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// [Logs documentation](https://docs.nobl9.com/Features/SLO_troubleshooting/event-logs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#log_collection_enabled DirectDatadog#log_collection_enabled}
	LogCollectionEnabled interface{} `field:"optional" json:"logCollectionEnabled" yaml:"logCollectionEnabled"`
	// query_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#query_delay DirectDatadog#query_delay}
	QueryDelay *DirectDatadogQueryDelay `field:"optional" json:"queryDelay" yaml:"queryDelay"`
	// Release channel of the created datasource [stable/beta].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#release_channel DirectDatadog#release_channel}
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
}

