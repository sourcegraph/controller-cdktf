package opsworksganglialayer


type OpsworksGangliaLayerLoadBasedAutoScaling struct {
	// downscaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#downscaling OpsworksGangliaLayer#downscaling}
	Downscaling *OpsworksGangliaLayerLoadBasedAutoScalingDownscaling `field:"optional" json:"downscaling" yaml:"downscaling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#enable OpsworksGangliaLayer#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// upscaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#upscaling OpsworksGangliaLayer#upscaling}
	Upscaling *OpsworksGangliaLayerLoadBasedAutoScalingUpscaling `field:"optional" json:"upscaling" yaml:"upscaling"`
}

