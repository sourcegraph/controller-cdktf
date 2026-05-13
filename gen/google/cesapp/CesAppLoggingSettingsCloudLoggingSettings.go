package cesapp


type CesAppLoggingSettingsCloudLoggingSettings struct {
	// Whether to enable Cloud Logging for the sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#enable_cloud_logging CesApp#enable_cloud_logging}
	EnableCloudLogging interface{} `field:"optional" json:"enableCloudLogging" yaml:"enableCloudLogging"`
}

