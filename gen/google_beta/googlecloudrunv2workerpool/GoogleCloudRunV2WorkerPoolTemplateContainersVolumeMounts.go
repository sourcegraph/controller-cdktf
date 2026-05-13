package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateContainersVolumeMounts struct {
	// Path within the container at which the volume should be mounted.
	//
	// Must not contain ':'. For Cloud SQL volumes, it can be left empty, or must otherwise be /cloudsql. All instances defined in the Volume will be available as /cloudsql/[instance]. For more information on Cloud SQL volumes, visit https://cloud.google.com/sql/docs/mysql/connect-run
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#mount_path GoogleCloudRunV2WorkerPool#mount_path}
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// This must match the Name of a Volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#name GoogleCloudRunV2WorkerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Path within the volume from which the container's volume should be mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#sub_path GoogleCloudRunV2WorkerPool#sub_path}
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}

