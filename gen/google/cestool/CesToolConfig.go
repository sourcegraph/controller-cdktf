package cestool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CesToolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#app CesTool#app}
	App *string `field:"required" json:"app" yaml:"app"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#location CesTool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID to use for the tool, which will become the final component of the tool's resource name.
	//
	// If not provided, a unique ID will be
	// automatically assigned for the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#tool_id CesTool#tool_id}
	ToolId *string `field:"required" json:"toolId" yaml:"toolId"`
	// client_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#client_function CesTool#client_function}
	ClientFunction *CesToolClientFunction `field:"optional" json:"clientFunction" yaml:"clientFunction"`
	// data_store_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#data_store_tool CesTool#data_store_tool}
	DataStoreTool *CesToolDataStoreTool `field:"optional" json:"dataStoreTool" yaml:"dataStoreTool"`
	// Possible values: SYNCHRONOUS ASYNCHRONOUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#execution_type CesTool#execution_type}
	ExecutionType *string `field:"optional" json:"executionType" yaml:"executionType"`
	// google_search_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#google_search_tool CesTool#google_search_tool}
	GoogleSearchTool *CesToolGoogleSearchTool `field:"optional" json:"googleSearchTool" yaml:"googleSearchTool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#id CesTool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#project CesTool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// python_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#python_function CesTool#python_function}
	PythonFunction *CesToolPythonFunction `field:"optional" json:"pythonFunction" yaml:"pythonFunction"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#timeouts CesTool#timeouts}
	Timeouts *CesToolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

