package bigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/bigquerytable/internal"
)

type BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference interface {
	cdktf.ComplexObject
	ColumnFamily() BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyList
	ColumnFamilyInput() any
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
	IgnoreUnspecifiedColumnFamilies() any
	SetIgnoreUnspecifiedColumnFamilies(val any)
	IgnoreUnspecifiedColumnFamiliesInput() any
	InternalValue() *BigqueryTableExternalDataConfigurationBigtableOptions
	SetInternalValue(val *BigqueryTableExternalDataConfigurationBigtableOptions)
	OutputColumnFamiliesAsJson() any
	SetOutputColumnFamiliesAsJson(val any)
	OutputColumnFamiliesAsJsonInput() any
	ReadRowkeyAsString() any
	SetReadRowkeyAsString(val any)
	ReadRowkeyAsStringInput() any
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
	PutColumnFamily(value any)
	ResetColumnFamily()
	ResetIgnoreUnspecifiedColumnFamilies()
	ResetOutputColumnFamiliesAsJson()
	ResetReadRowkeyAsString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference
type jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ColumnFamily() BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyList {
	var returns BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyList
	_jsii_.Get(
		j,
		"columnFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ColumnFamilyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"columnFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) IgnoreUnspecifiedColumnFamilies() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreUnspecifiedColumnFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) IgnoreUnspecifiedColumnFamiliesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreUnspecifiedColumnFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) InternalValue() *BigqueryTableExternalDataConfigurationBigtableOptions {
	var returns *BigqueryTableExternalDataConfigurationBigtableOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) OutputColumnFamiliesAsJson() any {
	var returns any
	_jsii_.Get(
		j,
		"outputColumnFamiliesAsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) OutputColumnFamiliesAsJsonInput() any {
	var returns any
	_jsii_.Get(
		j,
		"outputColumnFamiliesAsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ReadRowkeyAsString() any {
	var returns any
	_jsii_.Get(
		j,
		"readRowkeyAsString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ReadRowkeyAsStringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readRowkeyAsStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBigqueryTableExternalDataConfigurationBigtableOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryTableExternalDataConfigurationBigtableOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryTableExternalDataConfigurationBigtableOptionsOutputReference_Override(b BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) SetIgnoreUnspecifiedColumnFamilies(val any) {
	if err := j.validateSetIgnoreUnspecifiedColumnFamiliesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreUnspecifiedColumnFamilies",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) SetInternalValue(val *BigqueryTableExternalDataConfigurationBigtableOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) SetOutputColumnFamiliesAsJson(val any) {
	if err := j.validateSetOutputColumnFamiliesAsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputColumnFamiliesAsJson",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) SetReadRowkeyAsString(val any) {
	if err := j.validateSetReadRowkeyAsStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readRowkeyAsString",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) PutColumnFamily(value any) {
	if err := b.validatePutColumnFamilyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putColumnFamily",
		[]any{value},
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ResetColumnFamily() {
	_jsii_.InvokeVoid(
		b,
		"resetColumnFamily",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ResetIgnoreUnspecifiedColumnFamilies() {
	_jsii_.InvokeVoid(
		b,
		"resetIgnoreUnspecifiedColumnFamilies",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ResetOutputColumnFamiliesAsJson() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputColumnFamiliesAsJson",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ResetReadRowkeyAsString() {
	_jsii_.InvokeVoid(
		b,
		"resetReadRowkeyAsString",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		b,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
