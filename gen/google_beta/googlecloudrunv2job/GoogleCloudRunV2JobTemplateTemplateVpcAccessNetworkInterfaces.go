package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateVpcAccessNetworkInterfaces struct {
	// The VPC network that the Cloud Run resource will be able to send traffic to.
	//
	// At least one of network or subnetwork must be specified. If both
	// network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If network is not specified, it will be
	// looked up from the subnetwork.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job#network GoogleCloudRunV2Job#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The VPC subnetwork that the Cloud Run resource will get IPs from.
	//
	// At least one of network or subnetwork must be specified. If both
	// network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If subnetwork is not specified, the
	// subnetwork with the same name with the network will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job#subnetwork GoogleCloudRunV2Job#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// Network tags applied to this Cloud Run job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job#tags GoogleCloudRunV2Job#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

