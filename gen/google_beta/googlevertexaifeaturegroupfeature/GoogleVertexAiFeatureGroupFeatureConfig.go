package googlevertexaifeaturegroupfeature

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVertexAiFeatureGroupFeatureConfig struct {
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
	// The name of the Feature Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group_feature#feature_group GoogleVertexAiFeatureGroupFeature#feature_group}
	FeatureGroup *string `field:"required" json:"featureGroup" yaml:"featureGroup"`
	// The resource name of the Feature Group Feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group_feature#name GoogleVertexAiFeatureGroupFeature#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The region for the resource. It should be the same as the feature group's region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group_feature#region GoogleVertexAiFeatureGroupFeature#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The description of the FeatureGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group_feature#description GoogleVertexAiFeatureGroupFeature#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group_feature#id GoogleVertexAiFeatureGroupFeature#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels with user-defined metadata to organize your FeatureGroup.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group_feature#labels GoogleVertexAiFeatureGroupFeature#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group_feature#project GoogleVertexAiFeatureGroupFeature#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group_feature#timeouts GoogleVertexAiFeatureGroupFeature#timeouts}
	Timeouts *GoogleVertexAiFeatureGroupFeatureTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The name of the BigQuery Table/View column hosting data for this version.
	//
	// If no value is provided, will use featureId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group_feature#version_column_name GoogleVertexAiFeatureGroupFeature#version_column_name}
	VersionColumnName *string `field:"optional" json:"versionColumnName" yaml:"versionColumnName"`
}

