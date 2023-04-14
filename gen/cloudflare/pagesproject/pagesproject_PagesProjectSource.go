package pagesproject


type PagesProjectSource struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/pages_project#config PagesProject#config}
	Config *PagesProjectSourceConfig `field:"optional" json:"config" yaml:"config"`
	// Project host type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/pages_project#type PagesProject#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

