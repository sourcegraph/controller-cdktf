package cesexample


type CesExampleMessagesChunksImage struct {
	// Raw bytes of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#data CesExample#data}
	Data *string `field:"required" json:"data" yaml:"data"`
	// The IANA standard MIME type of the source data. Supported image types includes: * image/png * image/jpeg * image/webp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#mime_type CesExample#mime_type}
	MimeType *string `field:"required" json:"mimeType" yaml:"mimeType"`
}

