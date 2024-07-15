package googlevertexaifeaturestoreentitytype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVertexAiFeaturestoreEntitytypeConfig struct {
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
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#featurestore GoogleVertexAiFeaturestoreEntitytype#featurestore}
	Featurestore *string `field:"required" json:"featurestore" yaml:"featurestore"`
	// Optional. Description of the EntityType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#description GoogleVertexAiFeaturestoreEntitytype#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#id GoogleVertexAiFeaturestoreEntitytype#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs to assign to this EntityType.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#labels GoogleVertexAiFeaturestoreEntitytype#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// monitoring_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#monitoring_config GoogleVertexAiFeaturestoreEntitytype#monitoring_config}
	MonitoringConfig *GoogleVertexAiFeaturestoreEntitytypeMonitoringConfig `field:"optional" json:"monitoringConfig" yaml:"monitoringConfig"`
	// The name of the EntityType.
	//
	// This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#name GoogleVertexAiFeaturestoreEntitytype#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Config for data retention policy in offline storage.
	//
	// TTL in days for feature values that will be stored in offline storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#offline_storage_ttl_days GoogleVertexAiFeaturestoreEntitytype#offline_storage_ttl_days}
	OfflineStorageTtlDays *float64 `field:"optional" json:"offlineStorageTtlDays" yaml:"offlineStorageTtlDays"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#timeouts GoogleVertexAiFeaturestoreEntitytype#timeouts}
	Timeouts *GoogleVertexAiFeaturestoreEntitytypeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

