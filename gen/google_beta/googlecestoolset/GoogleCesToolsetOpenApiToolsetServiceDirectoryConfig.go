package googlecestoolset


type GoogleCesToolsetOpenApiToolsetServiceDirectoryConfig struct {
	// The name of [Service Directory](https://cloud.google.com/service-directory) service. Format: 'projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}'. Location of the service directory must be the same as the location of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#service GoogleCesToolset#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

