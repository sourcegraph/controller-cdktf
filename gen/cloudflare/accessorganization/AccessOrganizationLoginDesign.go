package accessorganization


type AccessOrganizationLoginDesign struct {
	// The background color on the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_organization#background_color AccessOrganization#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The text at the bottom of the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_organization#footer_text AccessOrganization#footer_text}
	FooterText *string `field:"optional" json:"footerText" yaml:"footerText"`
	// The text at the top of the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_organization#header_text AccessOrganization#header_text}
	HeaderText *string `field:"optional" json:"headerText" yaml:"headerText"`
	// The URL of the logo on the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_organization#logo_path AccessOrganization#logo_path}
	LogoPath *string `field:"optional" json:"logoPath" yaml:"logoPath"`
	// The text color on the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_organization#text_color AccessOrganization#text_color}
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

