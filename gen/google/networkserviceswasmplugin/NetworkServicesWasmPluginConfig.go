package networkserviceswasmplugin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesWasmPluginConfig struct {
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
	// The ID of the WasmPluginVersion resource that is the currently serving one.
	//
	// The version referred to must be a child of this WasmPlugin resource and should be listed in the "versions" field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#main_version_id NetworkServicesWasmPlugin#main_version_id}
	MainVersionId *string `field:"required" json:"mainVersionId" yaml:"mainVersionId"`
	// Identifier. Name of the WasmPlugin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#name NetworkServicesWasmPlugin#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#versions NetworkServicesWasmPlugin#versions}
	Versions interface{} `field:"required" json:"versions" yaml:"versions"`
	// Optional. A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#description NetworkServicesWasmPlugin#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#id NetworkServicesWasmPlugin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Set of labels associated with the WasmPlugin resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#labels NetworkServicesWasmPlugin#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of the traffic extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#location NetworkServicesWasmPlugin#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#log_config NetworkServicesWasmPlugin#log_config}
	LogConfig *NetworkServicesWasmPluginLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#project NetworkServicesWasmPlugin#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#timeouts NetworkServicesWasmPlugin#timeouts}
	Timeouts *NetworkServicesWasmPluginTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

