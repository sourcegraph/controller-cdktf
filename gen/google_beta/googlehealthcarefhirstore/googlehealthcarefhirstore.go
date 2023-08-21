package googlehealthcarefhirstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehealthcarefhirstore/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_healthcare_fhir_store google_healthcare_fhir_store}.
type GoogleHealthcareFhirStore interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComplexDataTypeReferenceParsing() *string
	SetComplexDataTypeReferenceParsing(val *string)
	ComplexDataTypeReferenceParsingInput() *string
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
	Dataset() *string
	SetDataset(val *string)
	DatasetInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableReferentialIntegrity() interface{}
	SetDisableReferentialIntegrity(val interface{})
	DisableReferentialIntegrityInput() interface{}
	DisableResourceVersioning() interface{}
	SetDisableResourceVersioning(val interface{})
	DisableResourceVersioningInput() interface{}
	EnableHistoryImport() interface{}
	SetEnableHistoryImport(val interface{})
	EnableHistoryImportInput() interface{}
	EnableUpdateCreate() interface{}
	SetEnableUpdateCreate(val interface{})
	EnableUpdateCreateInput() interface{}
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationConfig() GoogleHealthcareFhirStoreNotificationConfigOutputReference
	NotificationConfigInput() *GoogleHealthcareFhirStoreNotificationConfig
	NotificationConfigs() GoogleHealthcareFhirStoreNotificationConfigsList
	NotificationConfigsInput() interface{}
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
	SelfLink() *string
	StreamConfigs() GoogleHealthcareFhirStoreStreamConfigsList
	StreamConfigsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleHealthcareFhirStoreTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutNotificationConfig(value *GoogleHealthcareFhirStoreNotificationConfig)
	PutNotificationConfigs(value interface{})
	PutStreamConfigs(value interface{})
	PutTimeouts(value *GoogleHealthcareFhirStoreTimeouts)
	ResetComplexDataTypeReferenceParsing()
	ResetDisableReferentialIntegrity()
	ResetDisableResourceVersioning()
	ResetEnableHistoryImport()
	ResetEnableUpdateCreate()
	ResetId()
	ResetLabels()
	ResetNotificationConfig()
	ResetNotificationConfigs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStreamConfigs()
	ResetTimeouts()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleHealthcareFhirStore
type jsiiProxy_GoogleHealthcareFhirStore struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) ComplexDataTypeReferenceParsing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexDataTypeReferenceParsing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) ComplexDataTypeReferenceParsingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexDataTypeReferenceParsingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Dataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) DatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) DisableReferentialIntegrity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableReferentialIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) DisableReferentialIntegrityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableReferentialIntegrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) DisableResourceVersioning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableResourceVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) DisableResourceVersioningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableResourceVersioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) EnableHistoryImport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHistoryImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) EnableHistoryImportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHistoryImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) EnableUpdateCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUpdateCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) EnableUpdateCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUpdateCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) NotificationConfig() GoogleHealthcareFhirStoreNotificationConfigOutputReference {
	var returns GoogleHealthcareFhirStoreNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) NotificationConfigInput() *GoogleHealthcareFhirStoreNotificationConfig {
	var returns *GoogleHealthcareFhirStoreNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) NotificationConfigs() GoogleHealthcareFhirStoreNotificationConfigsList {
	var returns GoogleHealthcareFhirStoreNotificationConfigsList
	_jsii_.Get(
		j,
		"notificationConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) NotificationConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) StreamConfigs() GoogleHealthcareFhirStoreStreamConfigsList {
	var returns GoogleHealthcareFhirStoreStreamConfigsList
	_jsii_.Get(
		j,
		"streamConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) StreamConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Timeouts() GoogleHealthcareFhirStoreTimeoutsOutputReference {
	var returns GoogleHealthcareFhirStoreTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStore) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_healthcare_fhir_store google_healthcare_fhir_store} Resource.
func NewGoogleHealthcareFhirStore(scope constructs.Construct, id *string, config *GoogleHealthcareFhirStoreConfig) GoogleHealthcareFhirStore {
	_init_.Initialize()

	if err := validateNewGoogleHealthcareFhirStoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHealthcareFhirStore{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStore",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_healthcare_fhir_store google_healthcare_fhir_store} Resource.
func NewGoogleHealthcareFhirStore_Override(g GoogleHealthcareFhirStore, scope constructs.Construct, id *string, config *GoogleHealthcareFhirStoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStore",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetComplexDataTypeReferenceParsing(val *string) {
	if err := j.validateSetComplexDataTypeReferenceParsingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexDataTypeReferenceParsing",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetDataset(val *string) {
	if err := j.validateSetDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataset",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetDisableReferentialIntegrity(val interface{}) {
	if err := j.validateSetDisableReferentialIntegrityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableReferentialIntegrity",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetDisableResourceVersioning(val interface{}) {
	if err := j.validateSetDisableResourceVersioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableResourceVersioning",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetEnableHistoryImport(val interface{}) {
	if err := j.validateSetEnableHistoryImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHistoryImport",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetEnableUpdateCreate(val interface{}) {
	if err := j.validateSetEnableUpdateCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUpdateCreate",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStore)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
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
func GoogleHealthcareFhirStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleHealthcareFhirStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleHealthcareFhirStore_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleHealthcareFhirStore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStore",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleHealthcareFhirStore_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleHealthcareFhirStore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStore",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleHealthcareFhirStore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcareFhirStore) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) PutNotificationConfig(value *GoogleHealthcareFhirStoreNotificationConfig) {
	if err := g.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) PutNotificationConfigs(value interface{}) {
	if err := g.validatePutNotificationConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNotificationConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) PutStreamConfigs(value interface{}) {
	if err := g.validatePutStreamConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStreamConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) PutTimeouts(value *GoogleHealthcareFhirStoreTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetComplexDataTypeReferenceParsing() {
	_jsii_.InvokeVoid(
		g,
		"resetComplexDataTypeReferenceParsing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetDisableReferentialIntegrity() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableReferentialIntegrity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetDisableResourceVersioning() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableResourceVersioning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetEnableHistoryImport() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableHistoryImport",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetEnableUpdateCreate() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableUpdateCreate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetNotificationConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetNotificationConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetStreamConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetStreamConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ResetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStore) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

