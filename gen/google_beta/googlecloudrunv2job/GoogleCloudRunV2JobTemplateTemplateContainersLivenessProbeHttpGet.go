package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbeHttpGet struct {
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_v2_job#http_headers GoogleCloudRunV2Job#http_headers}
	HttpHeaders interface{} `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Path to access on the HTTP server. Defaults to '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_v2_job#path GoogleCloudRunV2Job#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

