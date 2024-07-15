package googlestoragebucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleStorageBucketConfig struct {
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
	// The Google Cloud Storage location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#location GoogleStorageBucket#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#name GoogleStorageBucket#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// autoclass block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#autoclass GoogleStorageBucket#autoclass}
	Autoclass *GoogleStorageBucketAutoclass `field:"optional" json:"autoclass" yaml:"autoclass"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#cors GoogleStorageBucket#cors}
	Cors interface{} `field:"optional" json:"cors" yaml:"cors"`
	// custom_placement_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#custom_placement_config GoogleStorageBucket#custom_placement_config}
	CustomPlacementConfig *GoogleStorageBucketCustomPlacementConfig `field:"optional" json:"customPlacementConfig" yaml:"customPlacementConfig"`
	// Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#default_event_based_hold GoogleStorageBucket#default_event_based_hold}
	DefaultEventBasedHold interface{} `field:"optional" json:"defaultEventBasedHold" yaml:"defaultEventBasedHold"`
	// Enables each object in the bucket to have its own retention policy, which prevents deletion until stored for a specific length of time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#enable_object_retention GoogleStorageBucket#enable_object_retention}
	EnableObjectRetention interface{} `field:"optional" json:"enableObjectRetention" yaml:"enableObjectRetention"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#encryption GoogleStorageBucket#encryption}
	Encryption *GoogleStorageBucketEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// When deleting a bucket, this boolean option will delete all contained objects.
	//
	// If you try to delete a bucket that contains objects, Terraform will fail that run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#force_destroy GoogleStorageBucket#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#id GoogleStorageBucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs to assign to the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#labels GoogleStorageBucket#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// lifecycle_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#lifecycle_rule GoogleStorageBucket#lifecycle_rule}
	LifecycleRule interface{} `field:"optional" json:"lifecycleRule" yaml:"lifecycleRule"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#logging GoogleStorageBucket#logging}
	Logging *GoogleStorageBucketLogging `field:"optional" json:"logging" yaml:"logging"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#project GoogleStorageBucket#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Prevents public access to a bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#public_access_prevention GoogleStorageBucket#public_access_prevention}
	PublicAccessPrevention *string `field:"optional" json:"publicAccessPrevention" yaml:"publicAccessPrevention"`
	// Enables Requester Pays on a storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#requester_pays GoogleStorageBucket#requester_pays}
	RequesterPays interface{} `field:"optional" json:"requesterPays" yaml:"requesterPays"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#retention_policy GoogleStorageBucket#retention_policy}
	RetentionPolicy *GoogleStorageBucketRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// Specifies the RPO setting of bucket.
	//
	// If set 'ASYNC_TURBO', The Turbo Replication will be enabled for the dual-region bucket. Value 'DEFAULT' will set RPO setting to default. Turbo Replication is only for buckets in dual-regions.See the docs for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#rpo GoogleStorageBucket#rpo}
	Rpo *string `field:"optional" json:"rpo" yaml:"rpo"`
	// soft_delete_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#soft_delete_policy GoogleStorageBucket#soft_delete_policy}
	SoftDeletePolicy *GoogleStorageBucketSoftDeletePolicy `field:"optional" json:"softDeletePolicy" yaml:"softDeletePolicy"`
	// The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#storage_class GoogleStorageBucket#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#timeouts GoogleStorageBucket#timeouts}
	Timeouts *GoogleStorageBucketTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Enables uniform bucket-level access on a bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#uniform_bucket_level_access GoogleStorageBucket#uniform_bucket_level_access}
	UniformBucketLevelAccess interface{} `field:"optional" json:"uniformBucketLevelAccess" yaml:"uniformBucketLevelAccess"`
	// versioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#versioning GoogleStorageBucket#versioning}
	Versioning *GoogleStorageBucketVersioning `field:"optional" json:"versioning" yaml:"versioning"`
	// website block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#website GoogleStorageBucket#website}
	Website *GoogleStorageBucketWebsite `field:"optional" json:"website" yaml:"website"`
}

