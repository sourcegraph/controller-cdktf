package dataplexentrylink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexEntryLinkConfig struct {
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
	// The id of the entry group this entry link is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_entry_link#entry_group_id DataplexEntryLink#entry_group_id}
	EntryGroupId *string `field:"required" json:"entryGroupId" yaml:"entryGroupId"`
	// The id of the entry link to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_entry_link#entry_link_id DataplexEntryLink#entry_link_id}
	EntryLinkId *string `field:"required" json:"entryLinkId" yaml:"entryLinkId"`
	// Relative resource name of the Entry Link Type used to create this Entry Link. For example: projects/dataplex-types/locations/global/entryLinkTypes/definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_entry_link#entry_link_type DataplexEntryLink#entry_link_type}
	EntryLinkType *string `field:"required" json:"entryLinkType" yaml:"entryLinkType"`
	// entry_references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_entry_link#entry_references DataplexEntryLink#entry_references}
	EntryReferences interface{} `field:"required" json:"entryReferences" yaml:"entryReferences"`
	// The location for the entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_entry_link#location DataplexEntryLink#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// aspects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_entry_link#aspects DataplexEntryLink#aspects}
	Aspects interface{} `field:"optional" json:"aspects" yaml:"aspects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_entry_link#id DataplexEntryLink#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_entry_link#project DataplexEntryLink#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_entry_link#timeouts DataplexEntryLink#timeouts}
	Timeouts *DataplexEntryLinkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

