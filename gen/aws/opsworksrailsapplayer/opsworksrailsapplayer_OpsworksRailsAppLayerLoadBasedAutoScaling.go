package opsworksrailsapplayer


type OpsworksRailsAppLayerLoadBasedAutoScaling struct {
	// downscaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#downscaling OpsworksRailsAppLayer#downscaling}
	Downscaling *OpsworksRailsAppLayerLoadBasedAutoScalingDownscaling `field:"optional" json:"downscaling" yaml:"downscaling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#enable OpsworksRailsAppLayer#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// upscaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#upscaling OpsworksRailsAppLayer#upscaling}
	Upscaling *OpsworksRailsAppLayerLoadBasedAutoScalingUpscaling `field:"optional" json:"upscaling" yaml:"upscaling"`
}

