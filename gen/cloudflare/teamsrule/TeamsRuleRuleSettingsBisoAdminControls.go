package teamsrule


type TeamsRuleRuleSettingsBisoAdminControls struct {
	// Disable copy-paste.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#disable_copy_paste TeamsRule#disable_copy_paste}
	DisableCopyPaste interface{} `field:"optional" json:"disableCopyPaste" yaml:"disableCopyPaste"`
	// Disable download.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#disable_download TeamsRule#disable_download}
	DisableDownload interface{} `field:"optional" json:"disableDownload" yaml:"disableDownload"`
	// Disable keyboard usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#disable_keyboard TeamsRule#disable_keyboard}
	DisableKeyboard interface{} `field:"optional" json:"disableKeyboard" yaml:"disableKeyboard"`
	// Disable printing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#disable_printing TeamsRule#disable_printing}
	DisablePrinting interface{} `field:"optional" json:"disablePrinting" yaml:"disablePrinting"`
	// Disable upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#disable_upload TeamsRule#disable_upload}
	DisableUpload interface{} `field:"optional" json:"disableUpload" yaml:"disableUpload"`
}

