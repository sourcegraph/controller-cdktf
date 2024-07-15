package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#name GoogleCloudRunV2Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloud_sql_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#cloud_sql_instance GoogleCloudRunV2Service#cloud_sql_instance}
	CloudSqlInstance *GoogleCloudRunV2ServiceTemplateVolumesCloudSqlInstance `field:"optional" json:"cloudSqlInstance" yaml:"cloudSqlInstance"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#empty_dir GoogleCloudRunV2Service#empty_dir}
	EmptyDir *GoogleCloudRunV2ServiceTemplateVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#gcs GoogleCloudRunV2Service#gcs}
	Gcs *GoogleCloudRunV2ServiceTemplateVolumesGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#nfs GoogleCloudRunV2Service#nfs}
	Nfs *GoogleCloudRunV2ServiceTemplateVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#secret GoogleCloudRunV2Service#secret}
	Secret *GoogleCloudRunV2ServiceTemplateVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

