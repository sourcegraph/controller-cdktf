package observabilitytracescope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityTraceScopeConfig struct {
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
	// GCP region the TraceScope is stored in. Only 'global' is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#location ObservabilityTraceScope#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Names of the projects that are included in this trace scope.
	//
	// *  'projects/[PROJECT_ID]'
	//
	// A trace scope can include a maximum of 20 projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#resource_names ObservabilityTraceScope#resource_names}
	ResourceNames *[]*string `field:"required" json:"resourceNames" yaml:"resourceNames"`
	// A client-assigned identifier for the trace scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#trace_scope_id ObservabilityTraceScope#trace_scope_id}
	TraceScopeId *string `field:"required" json:"traceScopeId" yaml:"traceScopeId"`
	// Describes this trace scope.
	//
	// The maximum length of the description is 8000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#description ObservabilityTraceScope#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#id ObservabilityTraceScope#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#project ObservabilityTraceScope#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#timeouts ObservabilityTraceScope#timeouts}
	Timeouts *ObservabilityTraceScopeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

