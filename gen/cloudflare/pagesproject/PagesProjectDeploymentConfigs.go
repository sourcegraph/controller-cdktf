package pagesproject


type PagesProjectDeploymentConfigs struct {
	// preview block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/pages_project#preview PagesProject#preview}
	Preview *PagesProjectDeploymentConfigsPreview `field:"optional" json:"preview" yaml:"preview"`
	// production block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/pages_project#production PagesProject#production}
	Production *PagesProjectDeploymentConfigsProduction `field:"optional" json:"production" yaml:"production"`
}

