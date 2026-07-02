package databasemigrationservicemigrationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/databasemigrationservicemigrationjob/internal"
)

type DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfig
	SetInternalValue(val *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfig)
	ObjectConfigs() DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsList
	ObjectConfigsInput() any
	ObjectsSelectionType() *string
	SetObjectsSelectionType(val *string)
	ObjectsSelectionTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutObjectConfigs(value any)
	ResetObjectConfigs()
	ResetObjectsSelectionType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference
type jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) InternalValue() *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfig {
	var returns *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ObjectConfigs() DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsList {
	var returns DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsList
	_jsii_.Get(
		j,
		"objectConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ObjectConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"objectConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ObjectsSelectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectsSelectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ObjectsSelectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectsSelectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference_Override(d DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) SetInternalValue(val *DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) SetObjectsSelectionType(val *string) {
	if err := j.validateSetObjectsSelectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectsSelectionType",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) PutObjectConfigs(value any) {
	if err := d.validatePutObjectConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putObjectConfigs",
		[]any{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ResetObjectConfigs() {
	_jsii_.InvokeVoid(
		d,
		"resetObjectConfigs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ResetObjectsSelectionType() {
	_jsii_.InvokeVoid(
		d,
		"resetObjectsSelectionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
