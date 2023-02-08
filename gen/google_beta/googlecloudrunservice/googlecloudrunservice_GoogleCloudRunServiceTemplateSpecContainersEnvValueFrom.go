package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersEnvValueFrom struct {
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#secret_key_ref GoogleCloudRunService#secret_key_ref}
	SecretKeyRef *GoogleCloudRunServiceTemplateSpecContainersEnvValueFromSecretKeyRef `field:"required" json:"secretKeyRef" yaml:"secretKeyRef"`
}

