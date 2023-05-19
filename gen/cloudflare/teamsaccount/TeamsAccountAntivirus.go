package teamsaccount


type TeamsAccountAntivirus struct {
	// Scan on file download.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#enabled_download_phase TeamsAccount#enabled_download_phase}
	EnabledDownloadPhase interface{} `field:"required" json:"enabledDownloadPhase" yaml:"enabledDownloadPhase"`
	// Scan on file upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#enabled_upload_phase TeamsAccount#enabled_upload_phase}
	EnabledUploadPhase interface{} `field:"required" json:"enabledUploadPhase" yaml:"enabledUploadPhase"`
	// Block requests for files that cannot be scanned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#fail_closed TeamsAccount#fail_closed}
	FailClosed interface{} `field:"required" json:"failClosed" yaml:"failClosed"`
}

