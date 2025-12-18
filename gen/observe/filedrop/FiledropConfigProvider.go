package filedrop


type FiledropConfigProvider struct {
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/filedrop#aws Filedrop#aws}
	Aws *FiledropConfigProviderAws `field:"required" json:"aws" yaml:"aws"`
}

