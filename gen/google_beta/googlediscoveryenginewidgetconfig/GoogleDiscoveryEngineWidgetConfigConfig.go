package googlediscoveryenginewidgetconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDiscoveryEngineWidgetConfigConfig struct {
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
	// The engine ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#engine_id GoogleDiscoveryEngineWidgetConfig#engine_id}
	EngineId *string `field:"required" json:"engineId" yaml:"engineId"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#location GoogleDiscoveryEngineWidgetConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// access_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#access_settings GoogleDiscoveryEngineWidgetConfig#access_settings}
	AccessSettings *GoogleDiscoveryEngineWidgetConfigAccessSettings `field:"optional" json:"accessSettings" yaml:"accessSettings"`
	// The collection ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#collection_id GoogleDiscoveryEngineWidgetConfig#collection_id}
	CollectionId *string `field:"optional" json:"collectionId" yaml:"collectionId"`
	// homepage_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#homepage_setting GoogleDiscoveryEngineWidgetConfig#homepage_setting}
	HomepageSetting *GoogleDiscoveryEngineWidgetConfigHomepageSetting `field:"optional" json:"homepageSetting" yaml:"homepageSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#id GoogleDiscoveryEngineWidgetConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#project GoogleDiscoveryEngineWidgetConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#timeouts GoogleDiscoveryEngineWidgetConfig#timeouts}
	Timeouts *GoogleDiscoveryEngineWidgetConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// ui_branding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#ui_branding GoogleDiscoveryEngineWidgetConfig#ui_branding}
	UiBranding *GoogleDiscoveryEngineWidgetConfigUiBranding `field:"optional" json:"uiBranding" yaml:"uiBranding"`
	// ui_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#ui_settings GoogleDiscoveryEngineWidgetConfig#ui_settings}
	UiSettings *GoogleDiscoveryEngineWidgetConfigUiSettings `field:"optional" json:"uiSettings" yaml:"uiSettings"`
	// The unique ID to use for the WidgetConfig. Currently only accepts "default_search_widget_config".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#widget_config_id GoogleDiscoveryEngineWidgetConfig#widget_config_id}
	WidgetConfigId *string `field:"optional" json:"widgetConfigId" yaml:"widgetConfigId"`
}

