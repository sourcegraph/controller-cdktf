package googlestoragebucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestoragebucket/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_storage_bucket google_storage_bucket}.
type GoogleStorageBucket interface {
	cdktf.TerraformResource
	Autoclass() GoogleStorageBucketAutoclassOutputReference
	AutoclassInput() *GoogleStorageBucketAutoclass
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Cors() GoogleStorageBucketCorsList
	CorsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomPlacementConfig() GoogleStorageBucketCustomPlacementConfigOutputReference
	CustomPlacementConfigInput() *GoogleStorageBucketCustomPlacementConfig
	DefaultEventBasedHold() interface{}
	SetDefaultEventBasedHold(val interface{})
	DefaultEventBasedHoldInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Encryption() GoogleStorageBucketEncryptionOutputReference
	EncryptionInput() *GoogleStorageBucketEncryption
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleRule() GoogleStorageBucketLifecycleRuleList
	LifecycleRuleInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Logging() GoogleStorageBucketLoggingOutputReference
	LoggingInput() *GoogleStorageBucketLogging
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicAccessPrevention() *string
	SetPublicAccessPrevention(val *string)
	PublicAccessPreventionInput() *string
	// Experimental.
	RawOverrides() interface{}
	RequesterPays() interface{}
	SetRequesterPays(val interface{})
	RequesterPaysInput() interface{}
	RetentionPolicy() GoogleStorageBucketRetentionPolicyOutputReference
	RetentionPolicyInput() *GoogleStorageBucketRetentionPolicy
	SelfLink() *string
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleStorageBucketTimeoutsOutputReference
	TimeoutsInput() interface{}
	UniformBucketLevelAccess() interface{}
	SetUniformBucketLevelAccess(val interface{})
	UniformBucketLevelAccessInput() interface{}
	Url() *string
	Versioning() GoogleStorageBucketVersioningOutputReference
	VersioningInput() *GoogleStorageBucketVersioning
	Website() GoogleStorageBucketWebsiteOutputReference
	WebsiteInput() *GoogleStorageBucketWebsite
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAutoclass(value *GoogleStorageBucketAutoclass)
	PutCors(value interface{})
	PutCustomPlacementConfig(value *GoogleStorageBucketCustomPlacementConfig)
	PutEncryption(value *GoogleStorageBucketEncryption)
	PutLifecycleRule(value interface{})
	PutLogging(value *GoogleStorageBucketLogging)
	PutRetentionPolicy(value *GoogleStorageBucketRetentionPolicy)
	PutTimeouts(value *GoogleStorageBucketTimeouts)
	PutVersioning(value *GoogleStorageBucketVersioning)
	PutWebsite(value *GoogleStorageBucketWebsite)
	ResetAutoclass()
	ResetCors()
	ResetCustomPlacementConfig()
	ResetDefaultEventBasedHold()
	ResetEncryption()
	ResetForceDestroy()
	ResetId()
	ResetLabels()
	ResetLifecycleRule()
	ResetLogging()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPublicAccessPrevention()
	ResetRequesterPays()
	ResetRetentionPolicy()
	ResetStorageClass()
	ResetTimeouts()
	ResetUniformBucketLevelAccess()
	ResetVersioning()
	ResetWebsite()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleStorageBucket
type jsiiProxy_GoogleStorageBucket struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleStorageBucket) Autoclass() GoogleStorageBucketAutoclassOutputReference {
	var returns GoogleStorageBucketAutoclassOutputReference
	_jsii_.Get(
		j,
		"autoclass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) AutoclassInput() *GoogleStorageBucketAutoclass {
	var returns *GoogleStorageBucketAutoclass
	_jsii_.Get(
		j,
		"autoclassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Cors() GoogleStorageBucketCorsList {
	var returns GoogleStorageBucketCorsList
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) CorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) CustomPlacementConfig() GoogleStorageBucketCustomPlacementConfigOutputReference {
	var returns GoogleStorageBucketCustomPlacementConfigOutputReference
	_jsii_.Get(
		j,
		"customPlacementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) CustomPlacementConfigInput() *GoogleStorageBucketCustomPlacementConfig {
	var returns *GoogleStorageBucketCustomPlacementConfig
	_jsii_.Get(
		j,
		"customPlacementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) DefaultEventBasedHold() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultEventBasedHold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) DefaultEventBasedHoldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultEventBasedHoldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Encryption() GoogleStorageBucketEncryptionOutputReference {
	var returns GoogleStorageBucketEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) EncryptionInput() *GoogleStorageBucketEncryption {
	var returns *GoogleStorageBucketEncryption
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) LifecycleRule() GoogleStorageBucketLifecycleRuleList {
	var returns GoogleStorageBucketLifecycleRuleList
	_jsii_.Get(
		j,
		"lifecycleRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) LifecycleRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Logging() GoogleStorageBucketLoggingOutputReference {
	var returns GoogleStorageBucketLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) LoggingInput() *GoogleStorageBucketLogging {
	var returns *GoogleStorageBucketLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) PublicAccessPrevention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicAccessPrevention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) PublicAccessPreventionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicAccessPreventionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) RequesterPays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) RequesterPaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) RetentionPolicy() GoogleStorageBucketRetentionPolicyOutputReference {
	var returns GoogleStorageBucketRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"retentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) RetentionPolicyInput() *GoogleStorageBucketRetentionPolicy {
	var returns *GoogleStorageBucketRetentionPolicy
	_jsii_.Get(
		j,
		"retentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Timeouts() GoogleStorageBucketTimeoutsOutputReference {
	var returns GoogleStorageBucketTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) UniformBucketLevelAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniformBucketLevelAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) UniformBucketLevelAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniformBucketLevelAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Versioning() GoogleStorageBucketVersioningOutputReference {
	var returns GoogleStorageBucketVersioningOutputReference
	_jsii_.Get(
		j,
		"versioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) VersioningInput() *GoogleStorageBucketVersioning {
	var returns *GoogleStorageBucketVersioning
	_jsii_.Get(
		j,
		"versioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) Website() GoogleStorageBucketWebsiteOutputReference {
	var returns GoogleStorageBucketWebsiteOutputReference
	_jsii_.Get(
		j,
		"website",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucket) WebsiteInput() *GoogleStorageBucketWebsite {
	var returns *GoogleStorageBucketWebsite
	_jsii_.Get(
		j,
		"websiteInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_storage_bucket google_storage_bucket} Resource.
func NewGoogleStorageBucket(scope constructs.Construct, id *string, config *GoogleStorageBucketConfig) GoogleStorageBucket {
	_init_.Initialize()

	if err := validateNewGoogleStorageBucketParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageBucket{}

	_jsii_.Create(
		"google-beta.googleStorageBucket.GoogleStorageBucket",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_storage_bucket google_storage_bucket} Resource.
func NewGoogleStorageBucket_Override(g GoogleStorageBucket, scope constructs.Construct, id *string, config *GoogleStorageBucketConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleStorageBucket.GoogleStorageBucket",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetDefaultEventBasedHold(val interface{}) {
	if err := j.validateSetDefaultEventBasedHoldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultEventBasedHold",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetPublicAccessPrevention(val *string) {
	if err := j.validateSetPublicAccessPreventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccessPrevention",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetRequesterPays(val interface{}) {
	if err := j.validateSetRequesterPaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPays",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucket)SetUniformBucketLevelAccess(val interface{}) {
	if err := j.validateSetUniformBucketLevelAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniformBucketLevelAccess",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GoogleStorageBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleStorageBucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleStorageBucket.GoogleStorageBucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleStorageBucket_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleStorageBucket_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleStorageBucket.GoogleStorageBucket",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleStorageBucket_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleStorageBucket_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleStorageBucket.GoogleStorageBucket",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleStorageBucket_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleStorageBucket.GoogleStorageBucket",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutAutoclass(value *GoogleStorageBucketAutoclass) {
	if err := g.validatePutAutoclassParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoclass",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutCors(value interface{}) {
	if err := g.validatePutCorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCors",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutCustomPlacementConfig(value *GoogleStorageBucketCustomPlacementConfig) {
	if err := g.validatePutCustomPlacementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomPlacementConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutEncryption(value *GoogleStorageBucketEncryption) {
	if err := g.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutLifecycleRule(value interface{}) {
	if err := g.validatePutLifecycleRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLifecycleRule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutLogging(value *GoogleStorageBucketLogging) {
	if err := g.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogging",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutRetentionPolicy(value *GoogleStorageBucketRetentionPolicy) {
	if err := g.validatePutRetentionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetentionPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutTimeouts(value *GoogleStorageBucketTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutVersioning(value *GoogleStorageBucketVersioning) {
	if err := g.validatePutVersioningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVersioning",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) PutWebsite(value *GoogleStorageBucketWebsite) {
	if err := g.validatePutWebsiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebsite",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetAutoclass() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoclass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetCors() {
	_jsii_.InvokeVoid(
		g,
		"resetCors",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetCustomPlacementConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomPlacementConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetDefaultEventBasedHold() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultEventBasedHold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetEncryption() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetLifecycleRule() {
	_jsii_.InvokeVoid(
		g,
		"resetLifecycleRule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetPublicAccessPrevention() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicAccessPrevention",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetRequesterPays() {
	_jsii_.InvokeVoid(
		g,
		"resetRequesterPays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetRetentionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetentionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetStorageClass() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetUniformBucketLevelAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetUniformBucketLevelAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetVersioning() {
	_jsii_.InvokeVoid(
		g,
		"resetVersioning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) ResetWebsite() {
	_jsii_.InvokeVoid(
		g,
		"resetWebsite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucket) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucket) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

