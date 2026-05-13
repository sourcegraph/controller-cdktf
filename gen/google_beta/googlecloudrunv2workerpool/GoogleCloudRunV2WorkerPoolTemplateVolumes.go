package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#name GoogleCloudRunV2WorkerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloud_sql_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#cloud_sql_instance GoogleCloudRunV2WorkerPool#cloud_sql_instance}
	CloudSqlInstance *GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance `field:"optional" json:"cloudSqlInstance" yaml:"cloudSqlInstance"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#empty_dir GoogleCloudRunV2WorkerPool#empty_dir}
	EmptyDir *GoogleCloudRunV2WorkerPoolTemplateVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#gcs GoogleCloudRunV2WorkerPool#gcs}
	Gcs *GoogleCloudRunV2WorkerPoolTemplateVolumesGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#nfs GoogleCloudRunV2WorkerPool#nfs}
	Nfs *GoogleCloudRunV2WorkerPoolTemplateVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#secret GoogleCloudRunV2WorkerPool#secret}
	Secret *GoogleCloudRunV2WorkerPoolTemplateVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

