package networkmanagersite


type NetworkmanagerSiteTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_site#create NetworkmanagerSite#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_site#delete NetworkmanagerSite#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_site#update NetworkmanagerSite#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

