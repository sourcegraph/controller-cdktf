package pagesproject


type PagesProjectDeploymentConfigs struct {
	// preview block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#preview PagesProject#preview}
	Preview *PagesProjectDeploymentConfigsPreview `field:"required" json:"preview" yaml:"preview"`
	// production block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#production PagesProject#production}
	Production *PagesProjectDeploymentConfigsProduction `field:"required" json:"production" yaml:"production"`
}

