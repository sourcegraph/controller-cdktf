package googlecestoolset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCesToolsetConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#app GoogleCesToolset#app}
	App *string `field:"required" json:"app" yaml:"app"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#location GoogleCesToolset#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID to use for the toolset, which will become the final component of the toolset's resource name.
	//
	// If not provided, a unique ID will be
	// automatically assigned for the toolset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#toolset_id GoogleCesToolset#toolset_id}
	ToolsetId *string `field:"required" json:"toolsetId" yaml:"toolsetId"`
	// The description of the toolset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#description GoogleCesToolset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the toolset. Must be unique within the same app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#display_name GoogleCesToolset#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Possible values: SYNCHRONOUS ASYNCHRONOUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#execution_type GoogleCesToolset#execution_type}
	ExecutionType *string `field:"optional" json:"executionType" yaml:"executionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#id GoogleCesToolset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mcp_toolset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#mcp_toolset GoogleCesToolset#mcp_toolset}
	McpToolset *GoogleCesToolsetMcpToolset `field:"optional" json:"mcpToolset" yaml:"mcpToolset"`
	// open_api_toolset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#open_api_toolset GoogleCesToolset#open_api_toolset}
	OpenApiToolset *GoogleCesToolsetOpenApiToolset `field:"optional" json:"openApiToolset" yaml:"openApiToolset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#project GoogleCesToolset#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#timeouts GoogleCesToolset#timeouts}
	Timeouts *GoogleCesToolsetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

