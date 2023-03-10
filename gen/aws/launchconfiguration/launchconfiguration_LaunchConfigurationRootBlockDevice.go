package launchconfiguration


type LaunchConfigurationRootBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_configuration#delete_on_termination LaunchConfiguration#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_configuration#encrypted LaunchConfiguration#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_configuration#iops LaunchConfiguration#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_configuration#throughput LaunchConfiguration#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_configuration#volume_size LaunchConfiguration#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_configuration#volume_type LaunchConfiguration#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

