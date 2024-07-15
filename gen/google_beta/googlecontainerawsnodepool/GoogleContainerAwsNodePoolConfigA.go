package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolConfigA struct {
	// config_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#config_encryption GoogleContainerAwsNodePool#config_encryption}
	ConfigEncryption *GoogleContainerAwsNodePoolConfigConfigEncryption `field:"required" json:"configEncryption" yaml:"configEncryption"`
	// The name of the AWS IAM role assigned to nodes in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#iam_instance_profile GoogleContainerAwsNodePool#iam_instance_profile}
	IamInstanceProfile *string `field:"required" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// autoscaling_metrics_collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#autoscaling_metrics_collection GoogleContainerAwsNodePool#autoscaling_metrics_collection}
	AutoscalingMetricsCollection *GoogleContainerAwsNodePoolConfigAutoscalingMetricsCollection `field:"optional" json:"autoscalingMetricsCollection" yaml:"autoscalingMetricsCollection"`
	// The OS image type to use on node pool instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#image_type GoogleContainerAwsNodePool#image_type}
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// instance_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#instance_placement GoogleContainerAwsNodePool#instance_placement}
	InstancePlacement *GoogleContainerAwsNodePoolConfigInstancePlacement `field:"optional" json:"instancePlacement" yaml:"instancePlacement"`
	// Optional. The AWS instance type. When unspecified, it defaults to `m5.large`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#instance_type GoogleContainerAwsNodePool#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Optional.
	//
	// The initial labels assigned to nodes of this node pool. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#labels GoogleContainerAwsNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#proxy_config GoogleContainerAwsNodePool#proxy_config}
	ProxyConfig *GoogleContainerAwsNodePoolConfigProxyConfig `field:"optional" json:"proxyConfig" yaml:"proxyConfig"`
	// root_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#root_volume GoogleContainerAwsNodePool#root_volume}
	RootVolume *GoogleContainerAwsNodePoolConfigRootVolume `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// Optional.
	//
	// The IDs of additional security groups to add to nodes in this pool. The manager will automatically create security groups with minimum rules needed for a functioning cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#security_group_ids GoogleContainerAwsNodePool#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// spot_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#spot_config GoogleContainerAwsNodePool#spot_config}
	SpotConfig *GoogleContainerAwsNodePoolConfigSpotConfig `field:"optional" json:"spotConfig" yaml:"spotConfig"`
	// ssh_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#ssh_config GoogleContainerAwsNodePool#ssh_config}
	SshConfig *GoogleContainerAwsNodePoolConfigSshConfig `field:"optional" json:"sshConfig" yaml:"sshConfig"`
	// Optional.
	//
	// Key/value metadata to assign to each underlying AWS resource. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#tags GoogleContainerAwsNodePool#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#taints GoogleContainerAwsNodePool#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
}

