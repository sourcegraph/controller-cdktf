package filedrop


type FiledropConfigA struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/filedrop#provider Filedrop#provider}
	Provider *FiledropConfigProvider `field:"required" json:"provider" yaml:"provider"`
}

