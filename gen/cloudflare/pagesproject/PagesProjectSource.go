package pagesproject


type PagesProjectSource struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#config PagesProject#config}
	Config *PagesProjectSourceConfig `field:"optional" json:"config" yaml:"config"`
	// Project host type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#type PagesProject#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

