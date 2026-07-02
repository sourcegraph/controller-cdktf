package databasemigrationservicemigrationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/databasemigrationservicemigrationjob/internal"
)

type DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifier
	SetInternalValue(val *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifier)
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	Table() *string
	SetTable(val *string)
	TableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDatabase()
	ResetSchema()
	ResetTable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference
type jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) InternalValue() *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifier {
	var returns *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifier
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewDatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference_Override(d DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SetInternalValue(val *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifier) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) ResetSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		d,
		"resetTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifierOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
