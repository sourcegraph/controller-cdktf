package dataawsredshiftclustercredentials

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/dataawsredshiftclustercredentials/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/redshift_cluster_credentials aws_redshift_cluster_credentials}.
type DataAwsRedshiftClusterCredentials interface {
	cdktf.TerraformDataSource
	AutoCreate() any
	SetAutoCreate(val any)
	AutoCreateInput() any
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterIdentifier() *string
	SetClusterIdentifier(val *string)
	ClusterIdentifierInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	DbGroups() *[]*string
	SetDbGroups(val *[]*string)
	DbGroupsInput() *[]*string
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	DbPassword() *string
	DbUser() *string
	SetDbUser(val *string)
	DbUserInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DurationSeconds() *float64
	SetDurationSeconds(val *float64)
	DurationSecondsInput() *float64
	Expiration() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAutoCreate()
	ResetDbGroups()
	ResetDbName()
	ResetDurationSeconds()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsRedshiftClusterCredentials
type jsiiProxy_DataAwsRedshiftClusterCredentials struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) AutoCreate() any {
	var returns any
	_jsii_.Get(
		j,
		"autoCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) AutoCreateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"autoCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) ClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DbGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DbGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DbPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DbUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DbUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) DurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) Expiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/redshift_cluster_credentials aws_redshift_cluster_credentials} Data Source.
func NewDataAwsRedshiftClusterCredentials(scope constructs.Construct, id *string, config *DataAwsRedshiftClusterCredentialsConfig) DataAwsRedshiftClusterCredentials {
	_init_.Initialize()

	if err := validateNewDataAwsRedshiftClusterCredentialsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRedshiftClusterCredentials{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRedshiftClusterCredentials.DataAwsRedshiftClusterCredentials",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/redshift_cluster_credentials aws_redshift_cluster_credentials} Data Source.
func NewDataAwsRedshiftClusterCredentials_Override(d DataAwsRedshiftClusterCredentials, scope constructs.Construct, id *string, config *DataAwsRedshiftClusterCredentialsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRedshiftClusterCredentials.DataAwsRedshiftClusterCredentials",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetAutoCreate(val any) {
	if err := j.validateSetAutoCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCreate",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetClusterIdentifier(val *string) {
	if err := j.validateSetClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetDbGroups(val *[]*string) {
	if err := j.validateSetDbGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbGroups",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetDbUser(val *string) {
	if err := j.validateSetDbUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbUser",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetDurationSeconds(val *float64) {
	if err := j.validateSetDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationSeconds",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsRedshiftClusterCredentials) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsRedshiftClusterCredentials resource upon running "cdktf plan <stack-name>".
func DataAwsRedshiftClusterCredentials_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsRedshiftClusterCredentials_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRedshiftClusterCredentials.DataAwsRedshiftClusterCredentials",
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
func DataAwsRedshiftClusterCredentials_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDataAwsRedshiftClusterCredentials_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRedshiftClusterCredentials.DataAwsRedshiftClusterCredentials",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRedshiftClusterCredentials_IsTerraformDataSource(x any) *bool {
	_init_.Initialize()

	if err := validateDataAwsRedshiftClusterCredentials_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRedshiftClusterCredentials.DataAwsRedshiftClusterCredentials",
		"isTerraformDataSource",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRedshiftClusterCredentials_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDataAwsRedshiftClusterCredentials_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRedshiftClusterCredentials.DataAwsRedshiftClusterCredentials",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

func DataAwsRedshiftClusterCredentials_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsRedshiftClusterCredentials.DataAwsRedshiftClusterCredentials",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ResetAutoCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ResetDbGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetDbGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ResetDbName() {
	_jsii_.InvokeVoid(
		d,
		"resetDbName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ResetDurationSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetDurationSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRedshiftClusterCredentials) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
