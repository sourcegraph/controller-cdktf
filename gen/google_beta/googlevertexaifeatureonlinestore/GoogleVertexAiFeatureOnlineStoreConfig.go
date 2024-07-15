package googlevertexaifeatureonlinestore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVertexAiFeatureOnlineStoreConfig struct {
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
	// The resource name of the Feature Online Store.
	//
	// This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#name GoogleVertexAiFeatureOnlineStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// bigtable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#bigtable GoogleVertexAiFeatureOnlineStore#bigtable}
	Bigtable *GoogleVertexAiFeatureOnlineStoreBigtable `field:"optional" json:"bigtable" yaml:"bigtable"`
	// dedicated_serving_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#dedicated_serving_endpoint GoogleVertexAiFeatureOnlineStore#dedicated_serving_endpoint}
	DedicatedServingEndpoint *GoogleVertexAiFeatureOnlineStoreDedicatedServingEndpoint `field:"optional" json:"dedicatedServingEndpoint" yaml:"dedicatedServingEndpoint"`
	// embedding_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#embedding_management GoogleVertexAiFeatureOnlineStore#embedding_management}
	EmbeddingManagement *GoogleVertexAiFeatureOnlineStoreEmbeddingManagement `field:"optional" json:"embeddingManagement" yaml:"embeddingManagement"`
	// If set to true, any FeatureViews and Features for this FeatureOnlineStore will also be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#force_destroy GoogleVertexAiFeatureOnlineStore#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#id GoogleVertexAiFeatureOnlineStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels with user-defined metadata to organize your feature online stores.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#labels GoogleVertexAiFeatureOnlineStore#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// optimized block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#optimized GoogleVertexAiFeatureOnlineStore#optimized}
	Optimized *GoogleVertexAiFeatureOnlineStoreOptimized `field:"optional" json:"optimized" yaml:"optimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#project GoogleVertexAiFeatureOnlineStore#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of feature online store. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#region GoogleVertexAiFeatureOnlineStore#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#timeouts GoogleVertexAiFeatureOnlineStore#timeouts}
	Timeouts *GoogleVertexAiFeatureOnlineStoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

