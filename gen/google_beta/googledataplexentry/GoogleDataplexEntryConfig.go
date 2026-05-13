package googledataplexentry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataplexEntryConfig struct {
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
	// The relative resource name of the entry type that was used to create this entry, in the format projects/{project_number}/locations/{locationId}/entryTypes/{entryTypeId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#entry_type GoogleDataplexEntry#entry_type}
	EntryType *string `field:"required" json:"entryType" yaml:"entryType"`
	// aspects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#aspects GoogleDataplexEntry#aspects}
	Aspects interface{} `field:"optional" json:"aspects" yaml:"aspects"`
	// The entry group id of the entry group the entry will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#entry_group_id GoogleDataplexEntry#entry_group_id}
	EntryGroupId *string `field:"optional" json:"entryGroupId" yaml:"entryGroupId"`
	// The entry id of the entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#entry_id GoogleDataplexEntry#entry_id}
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
	// entry_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#entry_source GoogleDataplexEntry#entry_source}
	EntrySource *GoogleDataplexEntryEntrySource `field:"optional" json:"entrySource" yaml:"entrySource"`
	// A name for the entry that can be referenced by an external system.
	//
	// For more information, see https://cloud.google.com/dataplex/docs/fully-qualified-names.
	// The maximum size of the field is 4000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#fully_qualified_name GoogleDataplexEntry#fully_qualified_name}
	FullyQualifiedName *string `field:"optional" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#id GoogleDataplexEntry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location where entry will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#location GoogleDataplexEntry#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The resource name of the parent entry, in the format projects/{project_number}/locations/{locationId}/entryGroups/{entryGroupId}/entries/{entryId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#parent_entry GoogleDataplexEntry#parent_entry}
	ParentEntry *string `field:"optional" json:"parentEntry" yaml:"parentEntry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#project GoogleDataplexEntry#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#timeouts GoogleDataplexEntry#timeouts}
	Timeouts *GoogleDataplexEntryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

