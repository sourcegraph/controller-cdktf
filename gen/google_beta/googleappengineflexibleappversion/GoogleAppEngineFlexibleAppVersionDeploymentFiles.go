package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionDeploymentFiles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#name GoogleAppEngineFlexibleAppVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Source URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#source_url GoogleAppEngineFlexibleAppVersion#source_url}
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// SHA1 checksum of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#sha1_sum GoogleAppEngineFlexibleAppVersion#sha1_sum}
	Sha1Sum *string `field:"optional" json:"sha1Sum" yaml:"sha1Sum"`
}

