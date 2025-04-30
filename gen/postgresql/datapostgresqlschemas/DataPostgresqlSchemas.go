package datapostgresqlschemas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/postgresql/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/postgresql/datapostgresqlschemas/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/schemas postgresql_schemas}.
type DataPostgresqlSchemas interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	IncludeSystemSchemas() interface{}
	SetIncludeSystemSchemas(val interface{})
	IncludeSystemSchemasInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RegexPattern() *string
	SetRegexPattern(val *string)
	RegexPatternInput() *string
	Schemas() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetId()
	ResetIncludeSystemSchemas()
	ResetLikeAllPatterns()
	ResetLikeAnyPatterns()
	ResetNotLikeAllPatterns()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegexPattern()
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
}

// The jsii proxy struct for DataPostgresqlSchemas
type jsiiProxy_DataPostgresqlSchemas struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataPostgresqlSchemas) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) IncludeSystemSchemas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSystemSchemas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) IncludeSystemSchemasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSystemSchemasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) LikeAllPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAllPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) LikeAllPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAllPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) LikeAnyPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAnyPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) LikeAnyPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAnyPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) NotLikeAllPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notLikeAllPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) NotLikeAllPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notLikeAllPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) RegexPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) RegexPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) Schemas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSchemas) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/schemas postgresql_schemas} Data Source.
func NewDataPostgresqlSchemas(scope constructs.Construct, id *string, config *DataPostgresqlSchemasConfig) DataPostgresqlSchemas {
	_init_.Initialize()

	if err := validateNewDataPostgresqlSchemasParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataPostgresqlSchemas{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.dataPostgresqlSchemas.DataPostgresqlSchemas",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/schemas postgresql_schemas} Data Source.
func NewDataPostgresqlSchemas_Override(d DataPostgresqlSchemas, scope constructs.Construct, id *string, config *DataPostgresqlSchemasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.dataPostgresqlSchemas.DataPostgresqlSchemas",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetIncludeSystemSchemas(val interface{}) {
	if err := j.validateSetIncludeSystemSchemasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSystemSchemas",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetLikeAllPatterns(val *[]*string) {
	if err := j.validateSetLikeAllPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likeAllPatterns",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetLikeAnyPatterns(val *[]*string) {
	if err := j.validateSetLikeAnyPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likeAnyPatterns",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetNotLikeAllPatterns(val *[]*string) {
	if err := j.validateSetNotLikeAllPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notLikeAllPatterns",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSchemas)SetRegexPattern(val *string) {
	if err := j.validateSetRegexPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regexPattern",
		val,
	)
}

// Generates CDKTF code for importing a DataPostgresqlSchemas resource upon running "cdktf plan <stack-name>".
func DataPostgresqlSchemas_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataPostgresqlSchemas_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.dataPostgresqlSchemas.DataPostgresqlSchemas",
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
func DataPostgresqlSchemas_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPostgresqlSchemas_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.dataPostgresqlSchemas.DataPostgresqlSchemas",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPostgresqlSchemas_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPostgresqlSchemas_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.dataPostgresqlSchemas.DataPostgresqlSchemas",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPostgresqlSchemas_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPostgresqlSchemas_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.dataPostgresqlSchemas.DataPostgresqlSchemas",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataPostgresqlSchemas_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.dataPostgresqlSchemas.DataPostgresqlSchemas",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataPostgresqlSchemas) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataPostgresqlSchemas) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataPostgresqlSchemas) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSchemas) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataPostgresqlSchemas) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataPostgresqlSchemas) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataPostgresqlSchemas) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataPostgresqlSchemas) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataPostgresqlSchemas) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataPostgresqlSchemas) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataPostgresqlSchemas) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSchemas) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataPostgresqlSchemas) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSchemas) ResetIncludeSystemSchemas() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeSystemSchemas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSchemas) ResetLikeAllPatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetLikeAllPatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSchemas) ResetLikeAnyPatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetLikeAnyPatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSchemas) ResetNotLikeAllPatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetNotLikeAllPatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSchemas) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSchemas) ResetRegexPattern() {
	_jsii_.InvokeVoid(
		d,
		"resetRegexPattern",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSchemas) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSchemas) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSchemas) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSchemas) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSchemas) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSchemas) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

