package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig struct {
	// Required. The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#service GoogleDeveloperConnectConnection#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

