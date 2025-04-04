package env


type EnvEnvValueFromFieldRef struct {
	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/env#api_version Env#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Path of the field to select in the specified API version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/env#field_path Env#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
}

