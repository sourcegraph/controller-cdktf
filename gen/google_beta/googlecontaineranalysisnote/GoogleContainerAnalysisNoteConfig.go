package googlecontaineranalysisnote

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleContainerAnalysisNoteConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// attestation_authority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#attestation_authority GoogleContainerAnalysisNote#attestation_authority}
	AttestationAuthority *GoogleContainerAnalysisNoteAttestationAuthority `field:"required" json:"attestationAuthority" yaml:"attestationAuthority"`
	// The name of the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#name GoogleContainerAnalysisNote#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Time of expiration for this note. Leave empty if note does not expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#expiration_time GoogleContainerAnalysisNote#expiration_time}
	ExpirationTime *string `field:"optional" json:"expirationTime" yaml:"expirationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#id GoogleContainerAnalysisNote#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A detailed description of the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#long_description GoogleContainerAnalysisNote#long_description}
	LongDescription *string `field:"optional" json:"longDescription" yaml:"longDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#project GoogleContainerAnalysisNote#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Names of other notes related to this note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#related_note_names GoogleContainerAnalysisNote#related_note_names}
	RelatedNoteNames *[]*string `field:"optional" json:"relatedNoteNames" yaml:"relatedNoteNames"`
	// related_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#related_url GoogleContainerAnalysisNote#related_url}
	RelatedUrl interface{} `field:"optional" json:"relatedUrl" yaml:"relatedUrl"`
	// A one sentence description of the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#short_description GoogleContainerAnalysisNote#short_description}
	ShortDescription *string `field:"optional" json:"shortDescription" yaml:"shortDescription"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#timeouts GoogleContainerAnalysisNote#timeouts}
	Timeouts *GoogleContainerAnalysisNoteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

