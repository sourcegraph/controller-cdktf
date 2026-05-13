package googlechroniclefeed


type GoogleChronicleFeedDetailsHttpSettings struct {
	// Possible values: SOURCE_DELETION_NEVER SOURCE_DELETION_ON_SUCCESS SOURCE_DELETION_ON_SUCCESS_FILES_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#source_deletion_option GoogleChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"optional" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
	// Possible values: FILES FOLDERS FOLDERS_RECURSIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#source_type GoogleChronicleFeed#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// HTTP URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#uri GoogleChronicleFeed#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

