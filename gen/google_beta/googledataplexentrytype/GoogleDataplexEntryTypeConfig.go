package googledataplexentrytype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataplexEntryTypeConfig struct {
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
	// Description of the EntryType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#description GoogleDataplexEntryType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User friendly display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#display_name GoogleDataplexEntryType#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The entry type id of the entry type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#entry_type_id GoogleDataplexEntryType#entry_type_id}
	EntryTypeId *string `field:"optional" json:"entryTypeId" yaml:"entryTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#id GoogleDataplexEntryType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the EntryType.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#labels GoogleDataplexEntryType#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location where entry type will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#location GoogleDataplexEntryType#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The platform that Entries of this type belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#platform GoogleDataplexEntryType#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#project GoogleDataplexEntryType#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// required_aspects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#required_aspects GoogleDataplexEntryType#required_aspects}
	RequiredAspects interface{} `field:"optional" json:"requiredAspects" yaml:"requiredAspects"`
	// The system that Entries of this type belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#system GoogleDataplexEntryType#system}
	SystemAttribute *string `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#timeouts GoogleDataplexEntryType#timeouts}
	Timeouts *GoogleDataplexEntryTypeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Indicates the class this Entry Type belongs to, for example, TABLE, DATABASE, MODEL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_entry_type#type_aliases GoogleDataplexEntryType#type_aliases}
	TypeAliases *[]*string `field:"optional" json:"typeAliases" yaml:"typeAliases"`
}

