package googlediscoveryengineservingconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDiscoveryEngineServingConfigConfig struct {
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
	// The ID of the engine associated with the serving config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#engine_id GoogleDiscoveryEngineServingConfig#engine_id}
	EngineId *string `field:"required" json:"engineId" yaml:"engineId"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#location GoogleDiscoveryEngineServingConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource IDs of the boost controls to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#boost_control_ids GoogleDiscoveryEngineServingConfig#boost_control_ids}
	BoostControlIds *[]*string `field:"optional" json:"boostControlIds" yaml:"boostControlIds"`
	// The collection ID. Currently only accepts "default_collection".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#collection_id GoogleDiscoveryEngineServingConfig#collection_id}
	CollectionId *string `field:"optional" json:"collectionId" yaml:"collectionId"`
	// The resource IDs of the filter controls to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#filter_control_ids GoogleDiscoveryEngineServingConfig#filter_control_ids}
	FilterControlIds *[]*string `field:"optional" json:"filterControlIds" yaml:"filterControlIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#id GoogleDiscoveryEngineServingConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#project GoogleDiscoveryEngineServingConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The resource IDs of the promote controls to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#promote_control_ids GoogleDiscoveryEngineServingConfig#promote_control_ids}
	PromoteControlIds *[]*string `field:"optional" json:"promoteControlIds" yaml:"promoteControlIds"`
	// The resource IDs of the redirect controls to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#redirect_control_ids GoogleDiscoveryEngineServingConfig#redirect_control_ids}
	RedirectControlIds *[]*string `field:"optional" json:"redirectControlIds" yaml:"redirectControlIds"`
	// 'The unique ID of the serving config. Currently only accepts "default_search".'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#serving_config_id GoogleDiscoveryEngineServingConfig#serving_config_id}
	ServingConfigId *string `field:"optional" json:"servingConfigId" yaml:"servingConfigId"`
	// The resource IDs of the synonyms controls to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#synonyms_control_ids GoogleDiscoveryEngineServingConfig#synonyms_control_ids}
	SynonymsControlIds *[]*string `field:"optional" json:"synonymsControlIds" yaml:"synonymsControlIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_serving_config#timeouts GoogleDiscoveryEngineServingConfig#timeouts}
	Timeouts *GoogleDiscoveryEngineServingConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

