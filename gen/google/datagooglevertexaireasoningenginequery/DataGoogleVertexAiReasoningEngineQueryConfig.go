package datagooglevertexaireasoningenginequery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleVertexAiReasoningEngineQueryConfig struct {
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
	// The id of the Vertex Agent Engine to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/data-sources/vertex_ai_reasoning_engine_query#reasoning_engine_id DataGoogleVertexAiReasoningEngineQuery#reasoning_engine_id}
	ReasoningEngineId *string `field:"required" json:"reasoningEngineId" yaml:"reasoningEngineId"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/data-sources/vertex_ai_reasoning_engine_query#region DataGoogleVertexAiReasoningEngineQuery#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Class method to be used for the query. It is optional and defaults to "query" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/data-sources/vertex_ai_reasoning_engine_query#class_method DataGoogleVertexAiReasoningEngineQuery#class_method}
	ClassMethod *string `field:"optional" json:"classMethod" yaml:"classMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/data-sources/vertex_ai_reasoning_engine_query#id DataGoogleVertexAiReasoningEngineQuery#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Input content provided by users in JSON object format. Examples include text query, function calling parameters, media bytes, etc..
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/data-sources/vertex_ai_reasoning_engine_query#input DataGoogleVertexAiReasoningEngineQuery#input}
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The project of the resource. If not provided, the provider default project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/data-sources/vertex_ai_reasoning_engine_query#project DataGoogleVertexAiReasoningEngineQuery#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

