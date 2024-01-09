package directazuremonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectAzureMonitorConfig struct {
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
	// [required] | Azure Application (client) ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#client_id DirectAzureMonitor#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Unique name of the resource, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#name DirectAzureMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the Nobl9 project the resource sits in, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#project DirectAzureMonitor#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Source of Metrics and/or Services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#source_of DirectAzureMonitor#source_of}
	SourceOf *[]*string `field:"required" json:"sourceOf" yaml:"sourceOf"`
	// [required] | Azure Tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#tenant_id DirectAzureMonitor#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// [required] | Azure Application (client) Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#client_secret DirectAzureMonitor#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Optional description of the resource.
	//
	// Here, you can add details about who is responsible for the integration (team/owner) or the purpose of creating it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#description DirectAzureMonitor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-friendly display name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#display_name DirectAzureMonitor#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// historical_data_retrieval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#historical_data_retrieval DirectAzureMonitor#historical_data_retrieval}
	HistoricalDataRetrieval *DirectAzureMonitorHistoricalDataRetrieval `field:"optional" json:"historicalDataRetrieval" yaml:"historicalDataRetrieval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#id DirectAzureMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// [Logs documentation](https://docs.nobl9.com/Features/SLO_troubleshooting/event-logs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#log_collection_enabled DirectAzureMonitor#log_collection_enabled}
	LogCollectionEnabled interface{} `field:"optional" json:"logCollectionEnabled" yaml:"logCollectionEnabled"`
	// query_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#query_delay DirectAzureMonitor#query_delay}
	QueryDelay *DirectAzureMonitorQueryDelay `field:"optional" json:"queryDelay" yaml:"queryDelay"`
	// Release channel of the created datasource [stable/beta].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_azure_monitor#release_channel DirectAzureMonitor#release_channel}
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
}

