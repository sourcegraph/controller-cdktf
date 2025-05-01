package googlecomputeimage


type GoogleComputeImageShieldedInstanceInitialStatePk struct {
	// The raw content in the secure keys file.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_image#content GoogleComputeImage#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The file type of source file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_image#file_type GoogleComputeImage#file_type}
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

