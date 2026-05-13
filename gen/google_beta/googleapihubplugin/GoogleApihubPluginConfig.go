package googleapihubplugin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApihubPluginConfig struct {
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
	// The display name of the plugin. Max length is 50 characters (Unicode code points).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#display_name GoogleApihubPlugin#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#location GoogleApihubPlugin#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID to use for the Plugin resource, which will become the final component of the Plugin's resource name.
	//
	// This field is optional.
	//
	// * If provided, the same will be used. The service will throw an error if
	// the specified id is already used by another Plugin resource in the API hub
	// instance.
	// * If not provided, a system generated id will be used.
	//
	// This value should be 4-63 characters, overall resource name which will be
	// of format
	// 'projects/{project}/locations/{location}/plugins/{plugin}',
	// its length is limited to 1000 characters and valid characters are
	// /a-z[0-9]-_/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#plugin_id GoogleApihubPlugin#plugin_id}
	PluginId *string `field:"required" json:"pluginId" yaml:"pluginId"`
	// actions_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#actions_config GoogleApihubPlugin#actions_config}
	ActionsConfig interface{} `field:"optional" json:"actionsConfig" yaml:"actionsConfig"`
	// config_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#config_template GoogleApihubPlugin#config_template}
	ConfigTemplate *GoogleApihubPluginConfigTemplate `field:"optional" json:"configTemplate" yaml:"configTemplate"`
	// The plugin description. Max length is 2000 characters (Unicode code points).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#description GoogleApihubPlugin#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// documentation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#documentation GoogleApihubPlugin#documentation}
	Documentation *GoogleApihubPluginDocumentation `field:"optional" json:"documentation" yaml:"documentation"`
	// hosting_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#hosting_service GoogleApihubPlugin#hosting_service}
	HostingService *GoogleApihubPluginHostingService `field:"optional" json:"hostingService" yaml:"hostingService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#id GoogleApihubPlugin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Possible values: PLUGIN_CATEGORY_UNSPECIFIED API_GATEWAY API_PRODUCER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#plugin_category GoogleApihubPlugin#plugin_category}
	PluginCategory *string `field:"optional" json:"pluginCategory" yaml:"pluginCategory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#project GoogleApihubPlugin#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#timeouts GoogleApihubPlugin#timeouts}
	Timeouts *GoogleApihubPluginTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

