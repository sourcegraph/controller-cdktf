package networkserviceswasmplugin


type NetworkServicesWasmPluginVersions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#version_name NetworkServicesWasmPlugin#version_name}.
	VersionName *string `field:"required" json:"versionName" yaml:"versionName"`
	// Optional. A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#description NetworkServicesWasmPlugin#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// URI of the container image containing the plugin, stored in the Artifact Registry. When a new WasmPluginVersion resource is created, the digest of the container image is saved in the imageDigest field.
	// When downloading an image, the digest value is used instead of an image tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#image_uri NetworkServicesWasmPlugin#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// Optional. Set of labels associated with the WasmPlugin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#labels NetworkServicesWasmPlugin#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// A base64-encoded string containing the configuration for the plugin.
	//
	// The configuration is provided to the plugin at runtime through the ON_CONFIGURE callback.
	// When a new WasmPluginVersion resource is created, the digest of the contents is saved in the pluginConfigDigest field.
	// Conflics with pluginConfigUri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#plugin_config_data NetworkServicesWasmPlugin#plugin_config_data}
	PluginConfigData *string `field:"optional" json:"pluginConfigData" yaml:"pluginConfigData"`
	// URI of the plugin configuration stored in the Artifact Registry.
	//
	// The configuration is provided to the plugin at runtime through the ON_CONFIGURE callback.
	// The container image must contain only a single file with the name plugin.config.
	// When a new WasmPluginVersion resource is created, the digest of the container image is saved in the pluginConfigDigest field.
	// Conflics with pluginConfigData.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_wasm_plugin#plugin_config_uri NetworkServicesWasmPlugin#plugin_config_uri}
	PluginConfigUri *string `field:"optional" json:"pluginConfigUri" yaml:"pluginConfigUri"`
}

