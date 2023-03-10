package mskconnectcustomplugin


type MskconnectCustomPluginLocation struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_custom_plugin#s3 MskconnectCustomPlugin#s3}
	S3 *MskconnectCustomPluginLocationS3 `field:"required" json:"s3" yaml:"s3"`
}

