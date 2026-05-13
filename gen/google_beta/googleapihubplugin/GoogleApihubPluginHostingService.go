package googleapihubplugin


type GoogleApihubPluginHostingService struct {
	// The URI of the service implemented by the plugin developer, used to invoke the plugin's functionality.
	//
	// This information is only required for
	// user defined plugins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#service_uri GoogleApihubPlugin#service_uri}
	ServiceUri *string `field:"optional" json:"serviceUri" yaml:"serviceUri"`
}

