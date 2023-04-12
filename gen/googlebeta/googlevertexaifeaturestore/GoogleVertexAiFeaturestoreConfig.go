package googlevertexaifeaturestore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVertexAiFeaturestoreConfig struct {
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
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#encryption_spec GoogleVertexAiFeaturestore#encryption_spec}
	EncryptionSpec *GoogleVertexAiFeaturestoreEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#force_destroy GoogleVertexAiFeaturestore#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#id GoogleVertexAiFeaturestore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs to assign to this Featurestore.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#labels GoogleVertexAiFeaturestore#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the Featurestore.
	//
	// This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#name GoogleVertexAiFeaturestore#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// online_serving_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#online_serving_config GoogleVertexAiFeaturestore#online_serving_config}
	OnlineServingConfig *GoogleVertexAiFeaturestoreOnlineServingConfig `field:"optional" json:"onlineServingConfig" yaml:"onlineServingConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#project GoogleVertexAiFeaturestore#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the dataset. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#region GoogleVertexAiFeaturestore#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#timeouts GoogleVertexAiFeaturestore#timeouts}
	Timeouts *GoogleVertexAiFeaturestoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

