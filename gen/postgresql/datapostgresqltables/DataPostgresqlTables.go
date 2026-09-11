package datapostgresqltables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/postgresql/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/postgresql/datapostgresqltables/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables postgresql_tables}.
type DataPostgresqlTables interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LikeAllPatterns() *[]*string
	SetLikeAllPatterns(val *[]*string)
	LikeAllPatternsInput() *[]*string
	LikeAnyPatterns() *[]*string
	SetLikeAnyPatterns(val *[]*string)
	LikeAnyPatternsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	NotLikeAllPatterns() *[]*string
	SetNotLikeAllPatterns(val *[]*string)
	NotLikeAllPatternsInput() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RegexPattern() *string
	SetRegexPattern(val *string)
	RegexPatternInput() *string
	Schemas() *[]*string
	SetSchemas(val *[]*string)
	SchemasInput() *[]*string
	Tables() DataPostgresqlTablesTablesList
	TableTypes() *[]*string
	SetTableTypes(val *[]*string)
	TableTypesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetId()
	ResetLikeAllPatterns()
	ResetLikeAnyPatterns()
	ResetNotLikeAllPatterns()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegexPattern()
	ResetSchemas()
	ResetTableTypes()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DataPostgresqlTables
type jsiiProxy_DataPostgresqlTables struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataPostgresqlTables) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) LikeAllPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAllPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) LikeAllPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAllPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) LikeAnyPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAnyPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) LikeAnyPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAnyPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) NotLikeAllPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notLikeAllPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) NotLikeAllPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notLikeAllPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) RegexPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) RegexPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) Schemas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) SchemasInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) Tables() DataPostgresqlTablesTablesList {
	var returns DataPostgresqlTablesTablesList
	_jsii_.Get(
		j,
		"tables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) TableTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tableTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) TableTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tableTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlTables) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables postgresql_tables} Data Source.
func NewDataPostgresqlTables(scope constructs.Construct, id *string, config *DataPostgresqlTablesConfig) DataPostgresqlTables {
	_init_.Initialize()

	if err := validateNewDataPostgresqlTablesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataPostgresqlTables{}

	_jsii_.Create(
		"@cdktn/provider-postgresql.dataPostgresqlTables.DataPostgresqlTables",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables postgresql_tables} Data Source.
func NewDataPostgresqlTables_Override(d DataPostgresqlTables, scope constructs.Construct, id *string, config *DataPostgresqlTablesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-postgresql.dataPostgresqlTables.DataPostgresqlTables",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetLikeAllPatterns(val *[]*string) {
	if err := j.validateSetLikeAllPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likeAllPatterns",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetLikeAnyPatterns(val *[]*string) {
	if err := j.validateSetLikeAnyPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likeAnyPatterns",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetNotLikeAllPatterns(val *[]*string) {
	if err := j.validateSetNotLikeAllPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notLikeAllPatterns",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetRegexPattern(val *string) {
	if err := j.validateSetRegexPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regexPattern",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetSchemas(val *[]*string) {
	if err := j.validateSetSchemasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemas",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlTables)SetTableTypes(val *[]*string) {
	if err := j.validateSetTableTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableTypes",
		val,
	)
}

// Generates CDKTN code for importing a DataPostgresqlTables resource upon running "cdktn plan <stack-name>".
func DataPostgresqlTables_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataPostgresqlTables_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-postgresql.dataPostgresqlTables.DataPostgresqlTables",
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
func DataPostgresqlTables_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPostgresqlTables_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-postgresql.dataPostgresqlTables.DataPostgresqlTables",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPostgresqlTables_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPostgresqlTables_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-postgresql.dataPostgresqlTables.DataPostgresqlTables",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPostgresqlTables_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPostgresqlTables_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-postgresql.dataPostgresqlTables.DataPostgresqlTables",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataPostgresqlTables_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-postgresql.dataPostgresqlTables.DataPostgresqlTables",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataPostgresqlTables) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataPostgresqlTables) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataPostgresqlTables) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlTables) ResetLikeAllPatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetLikeAllPatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlTables) ResetLikeAnyPatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetLikeAnyPatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlTables) ResetNotLikeAllPatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetNotLikeAllPatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlTables) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlTables) ResetRegexPattern() {
	_jsii_.InvokeVoid(
		d,
		"resetRegexPattern",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlTables) ResetSchemas() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlTables) ResetTableTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetTableTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlTables) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlTables) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

