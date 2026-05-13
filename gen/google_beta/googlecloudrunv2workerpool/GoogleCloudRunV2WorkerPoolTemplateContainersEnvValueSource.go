package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateContainersEnvValueSource struct {
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#secret_key_ref GoogleCloudRunV2WorkerPool#secret_key_ref}
	SecretKeyRef *GoogleCloudRunV2WorkerPoolTemplateContainersEnvValueSourceSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

