package opsworksnodejsapplayer


type OpsworksNodejsAppLayerLoadBasedAutoScaling struct {
	// downscaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#downscaling OpsworksNodejsAppLayer#downscaling}
	Downscaling *OpsworksNodejsAppLayerLoadBasedAutoScalingDownscaling `field:"optional" json:"downscaling" yaml:"downscaling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#enable OpsworksNodejsAppLayer#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// upscaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#upscaling OpsworksNodejsAppLayer#upscaling}
	Upscaling *OpsworksNodejsAppLayerLoadBasedAutoScalingUpscaling `field:"optional" json:"upscaling" yaml:"upscaling"`
}

