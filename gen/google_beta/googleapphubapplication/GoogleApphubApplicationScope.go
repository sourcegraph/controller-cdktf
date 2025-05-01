package googleapphubapplication


type GoogleApphubApplicationScope struct {
	// Required. Scope Type.   Possible values: REGIONAL GLOBAL Possible values: ["REGIONAL", "GLOBAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apphub_application#type GoogleApphubApplication#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

