package bigquerydataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/bigquerydataset/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_dataset google_bigquery_dataset}.
type BigqueryDataset interface {
	cdktf.TerraformResource
	Access() BigqueryDatasetAccessList
	AccessInput() any
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreationTime() *float64
	DatasetId() *string
	SetDatasetId(val *string)
	DatasetIdInput() *string
	DefaultCollation() *string
	SetDefaultCollation(val *string)
	DefaultCollationInput() *string
	DefaultEncryptionConfiguration() BigqueryDatasetDefaultEncryptionConfigurationOutputReference
	DefaultEncryptionConfigurationInput() *BigqueryDatasetDefaultEncryptionConfiguration
	DefaultPartitionExpirationMs() *float64
	SetDefaultPartitionExpirationMs(val *float64)
	DefaultPartitionExpirationMsInput() *float64
	DefaultTableExpirationMs() *float64
	SetDefaultTableExpirationMs(val *float64)
	DefaultTableExpirationMsInput() *float64
	DeleteContentsOnDestroy() any
	SetDeleteContentsOnDestroy(val any)
	DeleteContentsOnDestroyInput() any
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	Etag() *string
	ExternalCatalogDatasetOptions() BigqueryDatasetExternalCatalogDatasetOptionsOutputReference
	ExternalCatalogDatasetOptionsInput() *BigqueryDatasetExternalCatalogDatasetOptions
	ExternalDatasetReference() BigqueryDatasetExternalDatasetReferenceOutputReference
	ExternalDatasetReferenceInput() *BigqueryDatasetExternalDatasetReference
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
	IsCaseInsensitive() any
	SetIsCaseInsensitive(val any)
	IsCaseInsensitiveInput() any
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
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
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
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BigqueryDatasetTimeoutsOutputReference
	TimeoutsInput() any
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAccess(value any)
	PutDefaultEncryptionConfiguration(value *BigqueryDatasetDefaultEncryptionConfiguration)
	PutExternalCatalogDatasetOptions(value *BigqueryDatasetExternalCatalogDatasetOptions)
	PutExternalDatasetReference(value *BigqueryDatasetExternalDatasetReference)
	PutTimeouts(value *BigqueryDatasetTimeouts)
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
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for BigqueryDataset
type jsiiProxy_BigqueryDataset struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BigqueryDataset) Access() BigqueryDatasetAccessList {
	var returns BigqueryDatasetAccessList
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) AccessInput() any {
	var returns any
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultCollation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultCollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCollationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultEncryptionConfiguration() BigqueryDatasetDefaultEncryptionConfigurationOutputReference {
	var returns BigqueryDatasetDefaultEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"defaultEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultEncryptionConfigurationInput() *BigqueryDatasetDefaultEncryptionConfiguration {
	var returns *BigqueryDatasetDefaultEncryptionConfiguration
	_jsii_.Get(
		j,
		"defaultEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultPartitionExpirationMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPartitionExpirationMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultPartitionExpirationMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPartitionExpirationMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultTableExpirationMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTableExpirationMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultTableExpirationMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTableExpirationMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DeleteContentsOnDestroy() any {
	var returns any
	_jsii_.Get(
		j,
		"deleteContentsOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DeleteContentsOnDestroyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deleteContentsOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ExternalCatalogDatasetOptions() BigqueryDatasetExternalCatalogDatasetOptionsOutputReference {
	var returns BigqueryDatasetExternalCatalogDatasetOptionsOutputReference
	_jsii_.Get(
		j,
		"externalCatalogDatasetOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ExternalCatalogDatasetOptionsInput() *BigqueryDatasetExternalCatalogDatasetOptions {
	var returns *BigqueryDatasetExternalCatalogDatasetOptions
	_jsii_.Get(
		j,
		"externalCatalogDatasetOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ExternalDatasetReference() BigqueryDatasetExternalDatasetReferenceOutputReference {
	var returns BigqueryDatasetExternalDatasetReferenceOutputReference
	_jsii_.Get(
		j,
		"externalDatasetReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ExternalDatasetReferenceInput() *BigqueryDatasetExternalDatasetReference {
	var returns *BigqueryDatasetExternalDatasetReference
	_jsii_.Get(
		j,
		"externalDatasetReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) FriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) FriendlyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) IsCaseInsensitive() any {
	var returns any
	_jsii_.Get(
		j,
		"isCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) IsCaseInsensitiveInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isCaseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) LastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) MaxTimeTravelHours() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeTravelHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) MaxTimeTravelHoursInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeTravelHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ResourceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ResourceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) StorageBillingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBillingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) StorageBillingModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBillingModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Timeouts() BigqueryDatasetTimeoutsOutputReference {
	var returns BigqueryDatasetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_dataset google_bigquery_dataset} Resource.
func NewBigqueryDataset(scope constructs.Construct, id *string, config *BigqueryDatasetConfig) BigqueryDataset {
	_init_.Initialize()

	if err := validateNewBigqueryDatasetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryDataset{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDataset.BigqueryDataset",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_dataset google_bigquery_dataset} Resource.
func NewBigqueryDataset_Override(b BigqueryDataset, scope constructs.Construct, id *string, config *BigqueryDatasetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDataset.BigqueryDataset",
		[]any{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetDefaultCollation(val *string) {
	if err := j.validateSetDefaultCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultCollation",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetDefaultPartitionExpirationMs(val *float64) {
	if err := j.validateSetDefaultPartitionExpirationMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPartitionExpirationMs",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetDefaultTableExpirationMs(val *float64) {
	if err := j.validateSetDefaultTableExpirationMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTableExpirationMs",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetDeleteContentsOnDestroy(val any) {
	if err := j.validateSetDeleteContentsOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteContentsOnDestroy",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetFriendlyName(val *string) {
	if err := j.validateSetFriendlyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetIsCaseInsensitive(val any) {
	if err := j.validateSetIsCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCaseInsensitive",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetMaxTimeTravelHours(val *string) {
	if err := j.validateSetMaxTimeTravelHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTimeTravelHours",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetResourceTags(val *map[string]*string) {
	if err := j.validateSetResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset) SetStorageBillingModel(val *string) {
	if err := j.validateSetStorageBillingModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageBillingModel",
		val,
	)
}

// Generates CDKTF code for importing a BigqueryDataset resource upon running "cdktf plan <stack-name>".
func BigqueryDataset_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBigqueryDataset_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDataset.BigqueryDataset",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func BigqueryDataset_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateBigqueryDataset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDataset.BigqueryDataset",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDataset_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateBigqueryDataset_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDataset.BigqueryDataset",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDataset_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateBigqueryDataset_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDataset.BigqueryDataset",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func BigqueryDataset_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.bigqueryDataset.BigqueryDataset",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BigqueryDataset) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (b *jsiiProxy_BigqueryDataset) AddOverride(path *string, value any) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]any{path, value},
	)
}

func (b *jsiiProxy_BigqueryDataset) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]any{id, provider},
	)
}

func (b *jsiiProxy_BigqueryDataset) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]any{id},
	)
}

func (b *jsiiProxy_BigqueryDataset) MoveTo(moveTarget *string, index any) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (b *jsiiProxy_BigqueryDataset) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]any{id},
	)
}

func (b *jsiiProxy_BigqueryDataset) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutAccess(value any) {
	if err := b.validatePutAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAccess",
		[]any{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutDefaultEncryptionConfiguration(value *BigqueryDatasetDefaultEncryptionConfiguration) {
	if err := b.validatePutDefaultEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDefaultEncryptionConfiguration",
		[]any{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutExternalCatalogDatasetOptions(value *BigqueryDatasetExternalCatalogDatasetOptions) {
	if err := b.validatePutExternalCatalogDatasetOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putExternalCatalogDatasetOptions",
		[]any{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutExternalDatasetReference(value *BigqueryDatasetExternalDatasetReference) {
	if err := b.validatePutExternalDatasetReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putExternalDatasetReference",
		[]any{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutTimeouts(value *BigqueryDatasetTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]any{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetAccess() {
	_jsii_.InvokeVoid(
		b,
		"resetAccess",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDefaultCollation() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultCollation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDefaultEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultEncryptionConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDefaultPartitionExpirationMs() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultPartitionExpirationMs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDefaultTableExpirationMs() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultTableExpirationMs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDeleteContentsOnDestroy() {
	_jsii_.InvokeVoid(
		b,
		"resetDeleteContentsOnDestroy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetExternalCatalogDatasetOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetExternalCatalogDatasetOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetExternalDatasetReference() {
	_jsii_.InvokeVoid(
		b,
		"resetExternalDatasetReference",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetFriendlyName() {
	_jsii_.InvokeVoid(
		b,
		"resetFriendlyName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetIsCaseInsensitive() {
	_jsii_.InvokeVoid(
		b,
		"resetIsCaseInsensitive",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetLabels() {
	_jsii_.InvokeVoid(
		b,
		"resetLabels",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetLocation() {
	_jsii_.InvokeVoid(
		b,
		"resetLocation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetMaxTimeTravelHours() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxTimeTravelHours",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetProject() {
	_jsii_.InvokeVoid(
		b,
		"resetProject",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetResourceTags() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetStorageBillingModel() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageBillingModel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
