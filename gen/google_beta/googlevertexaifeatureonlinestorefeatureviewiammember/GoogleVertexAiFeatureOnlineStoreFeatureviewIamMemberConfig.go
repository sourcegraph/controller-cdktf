package googlevertexaifeatureonlinestorefeatureviewiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVertexAiFeatureOnlineStoreFeatureviewIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_member#feature_online_store GoogleVertexAiFeatureOnlineStoreFeatureviewIamMember#feature_online_store}.
	FeatureOnlineStore *string `field:"required" json:"featureOnlineStore" yaml:"featureOnlineStore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_member#feature_view GoogleVertexAiFeatureOnlineStoreFeatureviewIamMember#feature_view}.
	FeatureView *string `field:"required" json:"featureView" yaml:"featureView"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_member#member GoogleVertexAiFeatureOnlineStoreFeatureviewIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_member#role GoogleVertexAiFeatureOnlineStoreFeatureviewIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_member#condition GoogleVertexAiFeatureOnlineStoreFeatureviewIamMember#condition}
	Condition *GoogleVertexAiFeatureOnlineStoreFeatureviewIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_member#id GoogleVertexAiFeatureOnlineStoreFeatureviewIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_member#project GoogleVertexAiFeatureOnlineStoreFeatureviewIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_member#region GoogleVertexAiFeatureOnlineStoreFeatureviewIamMember#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

