package directsplunkobservability

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectSplunkObservabilityConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#name DirectSplunkObservability#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the Nobl9 project the resource sits in, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#project DirectSplunkObservability#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// SplunkObservability Realm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#realm DirectSplunkObservability#realm}
	Realm *string `field:"required" json:"realm" yaml:"realm"`
	// Source of Metrics and/or Services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#source_of DirectSplunkObservability#source_of}
	SourceOf *[]*string `field:"required" json:"sourceOf" yaml:"sourceOf"`
	// [required] | Splunk API Access Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#access_token DirectSplunkObservability#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Optional description of the resource.
	//
	// Here, you can add details about who is responsible for the integration (team/owner) or the purpose of creating it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#description DirectSplunkObservability#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-friendly display name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#display_name DirectSplunkObservability#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#id DirectSplunkObservability#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// query_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#query_delay DirectSplunkObservability#query_delay}
	QueryDelay *DirectSplunkObservabilityQueryDelay `field:"optional" json:"queryDelay" yaml:"queryDelay"`
	// Release channel of the created datasource [stable/beta].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk_observability#release_channel DirectSplunkObservability#release_channel}
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
}

