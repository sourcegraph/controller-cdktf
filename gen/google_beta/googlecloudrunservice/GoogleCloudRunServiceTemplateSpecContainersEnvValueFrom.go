package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersEnvValueFrom struct {
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#secret_key_ref GoogleCloudRunService#secret_key_ref}
	SecretKeyRef *GoogleCloudRunServiceTemplateSpecContainersEnvValueFromSecretKeyRef `field:"required" json:"secretKeyRef" yaml:"secretKeyRef"`
}

