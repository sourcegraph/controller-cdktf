package googledatacatalogentry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatacatalogentry/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_data_catalog_entry google_data_catalog_entry}.
type GoogleDataCatalogEntry interface {
	cdktf.TerraformResource
	BigqueryDateShardedSpec() GoogleDataCatalogEntryBigqueryDateShardedSpecList
	BigqueryTableSpec() GoogleDataCatalogEntryBigqueryTableSpecList
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EntryGroup() *string
	SetEntryGroup(val *string)
	EntryGroupInput() *string
	EntryId() *string
	SetEntryId(val *string)
	EntryIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsFilesetSpec() GoogleDataCatalogEntryGcsFilesetSpecOutputReference
	GcsFilesetSpecInput() *GoogleDataCatalogEntryGcsFilesetSpec
	Id() *string
	SetId(val *string)
	IdInput() *string
	IntegratedSystem() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkedResource() *string
	SetLinkedResource(val *string)
	LinkedResourceInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
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
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDataCatalogEntryTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UserSpecifiedSystem() *string
	SetUserSpecifiedSystem(val *string)
	UserSpecifiedSystemInput() *string
	UserSpecifiedType() *string
	SetUserSpecifiedType(val *string)
	UserSpecifiedTypeInput() *string
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
	PutGcsFilesetSpec(value *GoogleDataCatalogEntryGcsFilesetSpec)
	PutTimeouts(value *GoogleDataCatalogEntryTimeouts)
	ResetDescription()
	ResetDisplayName()
	ResetGcsFilesetSpec()
	ResetId()
	ResetLinkedResource()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchema()
	ResetTimeouts()
	ResetType()
	ResetUserSpecifiedSystem()
	ResetUserSpecifiedType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleDataCatalogEntry
type jsiiProxy_GoogleDataCatalogEntry struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDataCatalogEntry) BigqueryDateShardedSpec() GoogleDataCatalogEntryBigqueryDateShardedSpecList {
	var returns GoogleDataCatalogEntryBigqueryDateShardedSpecList
	_jsii_.Get(
		j,
		"bigqueryDateShardedSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) BigqueryTableSpec() GoogleDataCatalogEntryBigqueryTableSpecList {
	var returns GoogleDataCatalogEntryBigqueryTableSpecList
	_jsii_.Get(
		j,
		"bigqueryTableSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) EntryGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) EntryGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) EntryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) EntryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) GcsFilesetSpec() GoogleDataCatalogEntryGcsFilesetSpecOutputReference {
	var returns GoogleDataCatalogEntryGcsFilesetSpecOutputReference
	_jsii_.Get(
		j,
		"gcsFilesetSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) GcsFilesetSpecInput() *GoogleDataCatalogEntryGcsFilesetSpec {
	var returns *GoogleDataCatalogEntryGcsFilesetSpec
	_jsii_.Get(
		j,
		"gcsFilesetSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) IntegratedSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integratedSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) LinkedResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) LinkedResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Timeouts() GoogleDataCatalogEntryTimeoutsOutputReference {
	var returns GoogleDataCatalogEntryTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) UserSpecifiedSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSpecifiedSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) UserSpecifiedSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSpecifiedSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) UserSpecifiedType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSpecifiedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntry) UserSpecifiedTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSpecifiedTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_data_catalog_entry google_data_catalog_entry} Resource.
func NewGoogleDataCatalogEntry(scope constructs.Construct, id *string, config *GoogleDataCatalogEntryConfig) GoogleDataCatalogEntry {
	_init_.Initialize()

	if err := validateNewGoogleDataCatalogEntryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataCatalogEntry{}

	_jsii_.Create(
		"google-beta.googleDataCatalogEntry.GoogleDataCatalogEntry",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_data_catalog_entry google_data_catalog_entry} Resource.
func NewGoogleDataCatalogEntry_Override(g GoogleDataCatalogEntry, scope constructs.Construct, id *string, config *GoogleDataCatalogEntryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataCatalogEntry.GoogleDataCatalogEntry",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetEntryGroup(val *string) {
	if err := j.validateSetEntryGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetEntryId(val *string) {
	if err := j.validateSetEntryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryId",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetLinkedResource(val *string) {
	if err := j.validateSetLinkedResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetUserSpecifiedSystem(val *string) {
	if err := j.validateSetUserSpecifiedSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userSpecifiedSystem",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntry)SetUserSpecifiedType(val *string) {
	if err := j.validateSetUserSpecifiedTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userSpecifiedType",
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
func GoogleDataCatalogEntry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataCatalogEntry_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDataCatalogEntry.GoogleDataCatalogEntry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataCatalogEntry_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataCatalogEntry_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDataCatalogEntry.GoogleDataCatalogEntry",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataCatalogEntry_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataCatalogEntry_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDataCatalogEntry.GoogleDataCatalogEntry",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDataCatalogEntry_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleDataCatalogEntry.GoogleDataCatalogEntry",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDataCatalogEntry) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataCatalogEntry) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) PutGcsFilesetSpec(value *GoogleDataCatalogEntryGcsFilesetSpec) {
	if err := g.validatePutGcsFilesetSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsFilesetSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) PutTimeouts(value *GoogleDataCatalogEntryTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetGcsFilesetSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsFilesetSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetLinkedResource() {
	_jsii_.InvokeVoid(
		g,
		"resetLinkedResource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetUserSpecifiedSystem() {
	_jsii_.InvokeVoid(
		g,
		"resetUserSpecifiedSystem",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ResetUserSpecifiedType() {
	_jsii_.InvokeVoid(
		g,
		"resetUserSpecifiedType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntry) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataCatalogEntry) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

