package googlediscoveryenginesearchengine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDiscoveryEngineSearchEngineConfig struct {
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
	// The collection ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#collection_id GoogleDiscoveryEngineSearchEngine#collection_id}
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// The data stores associated with this engine.
	//
	// For SOLUTION_TYPE_SEARCH type of engines, they can only associate with at most one data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#data_store_ids GoogleDiscoveryEngineSearchEngine#data_store_ids}
	DataStoreIds *[]*string `field:"required" json:"dataStoreIds" yaml:"dataStoreIds"`
	// Required. The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#display_name GoogleDiscoveryEngineSearchEngine#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Unique ID to use for Search Engine App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#engine_id GoogleDiscoveryEngineSearchEngine#engine_id}
	EngineId *string `field:"required" json:"engineId" yaml:"engineId"`
	// Location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#location GoogleDiscoveryEngineSearchEngine#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// search_engine_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#search_engine_config GoogleDiscoveryEngineSearchEngine#search_engine_config}
	SearchEngineConfig *GoogleDiscoveryEngineSearchEngineSearchEngineConfig `field:"required" json:"searchEngineConfig" yaml:"searchEngineConfig"`
	// This is the application type this engine resource represents. The supported values: 'APP_TYPE_UNSPECIFIED', 'APP_TYPE_INTRANET'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#app_type GoogleDiscoveryEngineSearchEngine#app_type}
	AppType *string `field:"optional" json:"appType" yaml:"appType"`
	// common_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#common_config GoogleDiscoveryEngineSearchEngine#common_config}
	CommonConfig *GoogleDiscoveryEngineSearchEngineCommonConfig `field:"optional" json:"commonConfig" yaml:"commonConfig"`
	// Whether to disable analytics for searches performed on this engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#disable_analytics GoogleDiscoveryEngineSearchEngine#disable_analytics}
	DisableAnalytics interface{} `field:"optional" json:"disableAnalytics" yaml:"disableAnalytics"`
	// A map of the feature config for the engine to opt in or opt out of features.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#features GoogleDiscoveryEngineSearchEngine#features}
	Features *map[string]*string `field:"optional" json:"features" yaml:"features"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#id GoogleDiscoveryEngineSearchEngine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The industry vertical that the engine registers.
	//
	// The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to GENERIC. Vertical on Engine has to match vertical of the DataStore liniked to the engine. Default value: "GENERIC" Possible values: ["GENERIC", "MEDIA", "HEALTHCARE_FHIR"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#industry_vertical GoogleDiscoveryEngineSearchEngine#industry_vertical}
	IndustryVertical *string `field:"optional" json:"industryVertical" yaml:"industryVertical"`
	// The KMS key to be used to protect this Engine at creation time.
	//
	// Must be set for requests that need to comply with CMEK Org Policy
	// protections.
	//
	// If this field is set and processed successfully, the Engine will be
	// protected by the KMS key, as indicated in the cmek_config field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#kms_key_name GoogleDiscoveryEngineSearchEngine#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// knowledge_graph_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#knowledge_graph_config GoogleDiscoveryEngineSearchEngine#knowledge_graph_config}
	KnowledgeGraphConfig *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfig `field:"optional" json:"knowledgeGraphConfig" yaml:"knowledgeGraphConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#project GoogleDiscoveryEngineSearchEngine#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#timeouts GoogleDiscoveryEngineSearchEngine#timeouts}
	Timeouts *GoogleDiscoveryEngineSearchEngineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

