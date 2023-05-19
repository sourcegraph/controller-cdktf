package teamsaccount


type TeamsAccountBlockPage struct {
	// Hex code of block page background color.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#background_color TeamsAccount#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Indicator of enablement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#enabled TeamsAccount#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Block page footer text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#footer_text TeamsAccount#footer_text}
	FooterText *string `field:"optional" json:"footerText" yaml:"footerText"`
	// Block page header text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#header_text TeamsAccount#header_text}
	HeaderText *string `field:"optional" json:"headerText" yaml:"headerText"`
	// URL of block page logo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#logo_path TeamsAccount#logo_path}
	LogoPath *string `field:"optional" json:"logoPath" yaml:"logoPath"`
	// Admin email for users to contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#mailto_address TeamsAccount#mailto_address}
	MailtoAddress *string `field:"optional" json:"mailtoAddress" yaml:"mailtoAddress"`
	// Subject line for emails created from block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#mailto_subject TeamsAccount#mailto_subject}
	MailtoSubject *string `field:"optional" json:"mailtoSubject" yaml:"mailtoSubject"`
	// Name of block page configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#name TeamsAccount#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

