package googlevectorsearchcollection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVectorSearchCollectionConfig struct {
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
	// ID of the Collection to create.
	//
	// The id must be 1-63 characters long, and comply with
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	// Specifically, it must be 1-63 characters long and match the regular
	// expression '[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#collection_id GoogleVectorSearchCollection#collection_id}
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#location GoogleVectorSearchCollection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// JSON Schema for data. Field names must contain only alphanumeric characters, underscores, and hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#data_schema GoogleVectorSearchCollection#data_schema}
	DataSchema *string `field:"optional" json:"dataSchema" yaml:"dataSchema"`
	// User-specified description of the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#description GoogleVectorSearchCollection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-specified display name of the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#display_name GoogleVectorSearchCollection#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#encryption_spec GoogleVectorSearchCollection#encryption_spec}
	EncryptionSpec *GoogleVectorSearchCollectionEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#id GoogleVectorSearchCollection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#labels GoogleVectorSearchCollection#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#project GoogleVectorSearchCollection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#timeouts GoogleVectorSearchCollection#timeouts}
	Timeouts *GoogleVectorSearchCollectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vector_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vector_search_collection#vector_schema GoogleVectorSearchCollection#vector_schema}
	VectorSchema interface{} `field:"optional" json:"vectorSchema" yaml:"vectorSchema"`
}

