package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateVpcAccess struct {
	// VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}, where {project} can be project id or number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#connector GoogleCloudRunV2WorkerPool#connector}
	Connector *string `field:"optional" json:"connector" yaml:"connector"`
	// Traffic VPC egress settings. Possible values: ["ALL_TRAFFIC", "PRIVATE_RANGES_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#egress GoogleCloudRunV2WorkerPool#egress}
	Egress *string `field:"optional" json:"egress" yaml:"egress"`
	// network_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#network_interfaces GoogleCloudRunV2WorkerPool#network_interfaces}
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
}

