package opsworkscustomlayer

type OpsworksCustomLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/opsworks_custom_layer#enabled OpsworksCustomLayer#enabled}.
	Enabled any `field:"optional" json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/opsworks_custom_layer#log_streams OpsworksCustomLayer#log_streams}
	LogStreams any `field:"optional" json:"logStreams" yaml:"logStreams"`
}
