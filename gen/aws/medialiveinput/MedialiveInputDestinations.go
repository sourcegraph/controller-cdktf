package medialiveinput


type MedialiveInputDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_input#stream_name MedialiveInput#stream_name}.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
}

