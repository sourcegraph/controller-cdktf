package googleappenginestandardappversion


type GoogleAppEngineStandardAppVersionLibraries struct {
	// Name of the library. Example "django".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_app_engine_standard_app_version#name GoogleAppEngineStandardAppVersion#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Version of the library to select, or "latest".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_app_engine_standard_app_version#version GoogleAppEngineStandardAppVersion#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

