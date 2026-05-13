package googlediscoveryengineassistant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDiscoveryEngineAssistantConfig struct {
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
	// The unique id of the assistant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#assistant_id GoogleDiscoveryEngineAssistant#assistant_id}
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// The unique id of the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#collection_id GoogleDiscoveryEngineAssistant#collection_id}
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// The assistant display name.
	//
	// It must be a UTF-8 encoded string with a length limit of 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#display_name GoogleDiscoveryEngineAssistant#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The unique id of the engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#engine_id GoogleDiscoveryEngineAssistant#engine_id}
	EngineId *string `field:"required" json:"engineId" yaml:"engineId"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#location GoogleDiscoveryEngineAssistant#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// customer_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#customer_policy GoogleDiscoveryEngineAssistant#customer_policy}
	CustomerPolicy *GoogleDiscoveryEngineAssistantCustomerPolicy `field:"optional" json:"customerPolicy" yaml:"customerPolicy"`
	// Description for additional information. Expected to be shown on the configuration UI, not to the users of the assistant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#description GoogleDiscoveryEngineAssistant#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// generation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#generation_config GoogleDiscoveryEngineAssistant#generation_config}
	GenerationConfig *GoogleDiscoveryEngineAssistantGenerationConfig `field:"optional" json:"generationConfig" yaml:"generationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#id GoogleDiscoveryEngineAssistant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#project GoogleDiscoveryEngineAssistant#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#timeouts GoogleDiscoveryEngineAssistant#timeouts}
	Timeouts *GoogleDiscoveryEngineAssistantTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of web grounding to use. The supported values: 'WEB_GROUNDING_TYPE_DISABLED', 'WEB_GROUNDING_TYPE_GOOGLE_SEARCH', 'WEB_GROUNDING_TYPE_ENTERPRISE_WEB_SEARCH'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#web_grounding_type GoogleDiscoveryEngineAssistant#web_grounding_type}
	WebGroundingType *string `field:"optional" json:"webGroundingType" yaml:"webGroundingType"`
}

