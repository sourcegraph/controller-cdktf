package googlebigquerydataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigquerydataset/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_bigquery_dataset google_bigquery_dataset}.
type GoogleBigqueryDataset interface {
	cdktf.TerraformResource
	Access() GoogleBigqueryDatasetAccessList
	AccessInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *float64
	DatasetId() *string
	SetDatasetId(val *string)
	DatasetIdInput() *string
	DefaultCollation() *string
	SetDefaultCollation(val *string)
	DefaultCollationInput() *string
	DefaultEncryptionConfiguration() GoogleBigqueryDatasetDefaultEncryptionConfigurationOutputReference
	DefaultEncryptionConfigurationInput() *GoogleBigqueryDatasetDefaultEncryptionConfiguration
	DefaultPartitionExpirationMs() *float64
	SetDefaultPartitionExpirationMs(val *float64)
	DefaultPartitionExpirationMsInput() *float64
	DefaultTableExpirationMs() *float64
	SetDefaultTableExpirationMs(val *float64)
	DefaultTableExpirationMsInput() *float64
	DeleteContentsOnDestroy() interface{}
	SetDeleteContentsOnDestroy(val interface{})
	DeleteContentsOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	Etag() *string
	ExternalCatalogDatasetOptions() GoogleBigqueryDatasetExternalCatalogDatasetOptionsOutputReference
	ExternalCatalogDatasetOptionsInput() *GoogleBigqueryDatasetExternalCatalogDatasetOptions
	ExternalDatasetReference() GoogleBigqueryDatasetExternalDatasetReferenceOutputReference
	ExternalDatasetReferenceInput() *GoogleBigqueryDatasetExternalDatasetReference
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	FriendlyName() *string
	SetFriendlyName(val *string)
	FriendlyNameInput() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsCaseInsensitive() interface{}
	SetIsCaseInsensitive(val interface{})
	IsCaseInsensitiveInput() interface{}
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LastModifiedTime() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxTimeTravelHours() *string
	SetMaxTimeTravelHours(val *string)
	MaxTimeTravelHoursInput() *string
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
	// Experimental.
	RawOverrides() interface{}
	ResourceTags() *map[string]*string
	SetResourceTags(val *map[string]*string)
	ResourceTagsInput() *map[string]*string
	SelfLink() *string
	StorageBillingModel() *string
	SetStorageBillingModel(val *string)
	StorageBillingModelInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBigqueryDatasetTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAccess(value interface{})
	PutDefaultEncryptionConfiguration(value *GoogleBigqueryDatasetDefaultEncryptionConfiguration)
	PutExternalCatalogDatasetOptions(value *GoogleBigqueryDatasetExternalCatalogDatasetOptions)
	PutExternalDatasetReference(value *GoogleBigqueryDatasetExternalDatasetReference)
	PutTimeouts(value *GoogleBigqueryDatasetTimeouts)
	ResetAccess()
	ResetDefaultCollation()
	ResetDefaultEncryptionConfiguration()
	ResetDefaultPartitionExpirationMs()
	ResetDefaultTableExpirationMs()
	ResetDeleteContentsOnDestroy()
	ResetDescription()
	ResetExternalCatalogDatasetOptions()
	ResetExternalDatasetReference()
	ResetFriendlyName()
	ResetId()
	ResetIsCaseInsensitive()
	ResetLabels()
	ResetLocation()
	ResetMaxTimeTravelHours()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetResourceTags()
	ResetStorageBillingModel()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleBigqueryDataset
type jsiiProxy_GoogleBigqueryDataset struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBigqueryDataset) Access() GoogleBigqueryDatasetAccessList {
	var returns GoogleBigqueryDatasetAccessList
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) AccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DefaultCollation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DefaultCollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCollationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DefaultEncryptionConfiguration() GoogleBigqueryDatasetDefaultEncryptionConfigurationOutputReference {
	var returns GoogleBigqueryDatasetDefaultEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"defaultEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DefaultEncryptionConfigurationInput() *GoogleBigqueryDatasetDefaultEncryptionConfiguration {
	var returns *GoogleBigqueryDatasetDefaultEncryptionConfiguration
	_jsii_.Get(
		j,
		"defaultEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DefaultPartitionExpirationMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPartitionExpirationMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DefaultPartitionExpirationMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPartitionExpirationMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DefaultTableExpirationMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTableExpirationMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DefaultTableExpirationMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTableExpirationMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DeleteContentsOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteContentsOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DeleteContentsOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteContentsOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) ExternalCatalogDatasetOptions() GoogleBigqueryDatasetExternalCatalogDatasetOptionsOutputReference {
	var returns GoogleBigqueryDatasetExternalCatalogDatasetOptionsOutputReference
	_jsii_.Get(
		j,
		"externalCatalogDatasetOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) ExternalCatalogDatasetOptionsInput() *GoogleBigqueryDatasetExternalCatalogDatasetOptions {
	var returns *GoogleBigqueryDatasetExternalCatalogDatasetOptions
	_jsii_.Get(
		j,
		"externalCatalogDatasetOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) ExternalDatasetReference() GoogleBigqueryDatasetExternalDatasetReferenceOutputReference {
	var returns GoogleBigqueryDatasetExternalDatasetReferenceOutputReference
	_jsii_.Get(
		j,
		"externalDatasetReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) ExternalDatasetReferenceInput() *GoogleBigqueryDatasetExternalDatasetReference {
	var returns *GoogleBigqueryDatasetExternalDatasetReference
	_jsii_.Get(
		j,
		"externalDatasetReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) FriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) FriendlyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) IsCaseInsensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) IsCaseInsensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCaseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) LastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) MaxTimeTravelHours() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeTravelHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) MaxTimeTravelHoursInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeTravelHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) ResourceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) ResourceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) StorageBillingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBillingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) StorageBillingModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBillingModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) Timeouts() GoogleBigqueryDatasetTimeoutsOutputReference {
	var returns GoogleBigqueryDatasetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDataset) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_bigquery_dataset google_bigquery_dataset} Resource.
func NewGoogleBigqueryDataset(scope constructs.Construct, id *string, config *GoogleBigqueryDatasetConfig) GoogleBigqueryDataset {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryDatasetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryDataset{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryDataset.GoogleBigqueryDataset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_bigquery_dataset google_bigquery_dataset} Resource.
func NewGoogleBigqueryDataset_Override(g GoogleBigqueryDataset, scope constructs.Construct, id *string, config *GoogleBigqueryDatasetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryDataset.GoogleBigqueryDataset",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetDefaultCollation(val *string) {
	if err := j.validateSetDefaultCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultCollation",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetDefaultPartitionExpirationMs(val *float64) {
	if err := j.validateSetDefaultPartitionExpirationMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPartitionExpirationMs",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetDefaultTableExpirationMs(val *float64) {
	if err := j.validateSetDefaultTableExpirationMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTableExpirationMs",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetDeleteContentsOnDestroy(val interface{}) {
	if err := j.validateSetDeleteContentsOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteContentsOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetFriendlyName(val *string) {
	if err := j.validateSetFriendlyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetIsCaseInsensitive(val interface{}) {
	if err := j.validateSetIsCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCaseInsensitive",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetMaxTimeTravelHours(val *string) {
	if err := j.validateSetMaxTimeTravelHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTimeTravelHours",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetResourceTags(val *map[string]*string) {
	if err := j.validateSetResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDataset)SetStorageBillingModel(val *string) {
	if err := j.validateSetStorageBillingModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageBillingModel",
		val,
	)
}

// Generates CDKTF code for importing a GoogleBigqueryDataset resource upon running "cdktf plan <stack-name>".
func GoogleBigqueryDataset_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleBigqueryDataset_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryDataset.GoogleBigqueryDataset",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func GoogleBigqueryDataset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryDataset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryDataset.GoogleBigqueryDataset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryDataset_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryDataset_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryDataset.GoogleBigqueryDataset",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryDataset_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryDataset_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryDataset.GoogleBigqueryDataset",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBigqueryDataset_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleBigqueryDataset.GoogleBigqueryDataset",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBigqueryDataset) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryDataset) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryDataset) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryDataset) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryDataset) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryDataset) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryDataset) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryDataset) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryDataset) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryDataset) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDataset) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryDataset) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) PutAccess(value interface{}) {
	if err := g.validatePutAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccess",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) PutDefaultEncryptionConfiguration(value *GoogleBigqueryDatasetDefaultEncryptionConfiguration) {
	if err := g.validatePutDefaultEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) PutExternalCatalogDatasetOptions(value *GoogleBigqueryDatasetExternalCatalogDatasetOptions) {
	if err := g.validatePutExternalCatalogDatasetOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExternalCatalogDatasetOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) PutExternalDatasetReference(value *GoogleBigqueryDatasetExternalDatasetReference) {
	if err := g.validatePutExternalDatasetReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExternalDatasetReference",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) PutTimeouts(value *GoogleBigqueryDatasetTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetDefaultCollation() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultCollation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetDefaultEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultEncryptionConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetDefaultPartitionExpirationMs() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultPartitionExpirationMs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetDefaultTableExpirationMs() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultTableExpirationMs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetDeleteContentsOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteContentsOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetExternalCatalogDatasetOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalCatalogDatasetOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetExternalDatasetReference() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalDatasetReference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetFriendlyName() {
	_jsii_.InvokeVoid(
		g,
		"resetFriendlyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetIsCaseInsensitive() {
	_jsii_.InvokeVoid(
		g,
		"resetIsCaseInsensitive",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetMaxTimeTravelHours() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTimeTravelHours",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetResourceTags() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetStorageBillingModel() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageBillingModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDataset) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDataset) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDataset) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDataset) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDataset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDataset) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

