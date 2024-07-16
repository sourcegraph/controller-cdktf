package googleappenginestandardappversion


type GoogleAppEngineStandardAppVersionDeployment struct {
	// files block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#files GoogleAppEngineStandardAppVersion#files}
	Files interface{} `field:"optional" json:"files" yaml:"files"`
	// zip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#zip GoogleAppEngineStandardAppVersion#zip}
	Zip *GoogleAppEngineStandardAppVersionDeploymentZip `field:"optional" json:"zip" yaml:"zip"`
}

