package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig struct {
	// Required. The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#service GoogleCloudbuildv2Connection#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

