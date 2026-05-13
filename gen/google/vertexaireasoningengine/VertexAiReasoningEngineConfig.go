package vertexaireasoningengine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiReasoningEngineConfig struct {
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
	// The display name of the ReasoningEngine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#display_name VertexAiReasoningEngine#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Optional.
	//
	// The deletion policy for the reasoning engine. Setting this to FORCE allows the reasoning engine to be deleted regardless of child undeleted resources. Possible values: ["FORCE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#deletion_policy VertexAiReasoningEngine#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The description of the ReasoningEngine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#description VertexAiReasoningEngine#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#encryption_spec VertexAiReasoningEngine#encryption_spec}
	EncryptionSpec *VertexAiReasoningEngineEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#id VertexAiReasoningEngine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels associated with this ReasoningEngine. You can use these to organize and group your ReasoningEngines.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#labels VertexAiReasoningEngine#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#project VertexAiReasoningEngine#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the reasoning engine. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#region VertexAiReasoningEngine#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#spec VertexAiReasoningEngine#spec}
	Spec *VertexAiReasoningEngineSpec `field:"optional" json:"spec" yaml:"spec"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#timeouts VertexAiReasoningEngine#timeouts}
	Timeouts *VertexAiReasoningEngineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

