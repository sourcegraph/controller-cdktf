package teamsaccount


type TeamsAccountAntivirus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#enabled_download_phase TeamsAccount#enabled_download_phase}.
	EnabledDownloadPhase interface{} `field:"required" json:"enabledDownloadPhase" yaml:"enabledDownloadPhase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#enabled_upload_phase TeamsAccount#enabled_upload_phase}.
	EnabledUploadPhase interface{} `field:"required" json:"enabledUploadPhase" yaml:"enabledUploadPhase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#fail_closed TeamsAccount#fail_closed}.
	FailClosed interface{} `field:"required" json:"failClosed" yaml:"failClosed"`
}

