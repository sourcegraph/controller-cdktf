package googlechronicledatatable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechronicledatatable/internal"
)

type GoogleChronicleDataTableColumnInfoOutputReference interface {
	cdktf.ComplexObject
	ColumnIndex() *float64
	SetColumnIndex(val *float64)
	ColumnIndexInput() *float64
	ColumnType() *string
	SetColumnType(val *string)
	ColumnTypeInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyColumn() interface{}
	SetKeyColumn(val interface{})
	KeyColumnInput() interface{}
	MappedColumnPath() *string
	SetMappedColumnPath(val *string)
	MappedColumnPathInput() *string
	OriginalColumn() *string
	SetOriginalColumn(val *string)
	OriginalColumnInput() *string
	RepeatedValues() interface{}
	SetRepeatedValues(val interface{})
	RepeatedValuesInput() interface{}
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetColumnType()
	ResetKeyColumn()
	ResetMappedColumnPath()
	ResetRepeatedValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleDataTableColumnInfoOutputReference
type jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ColumnIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ColumnIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ColumnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ColumnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) KeyColumn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) KeyColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) MappedColumnPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappedColumnPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) MappedColumnPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappedColumnPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) OriginalColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) OriginalColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) RepeatedValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repeatedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) RepeatedValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repeatedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDataTableColumnInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleChronicleDataTableColumnInfoOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDataTableColumnInfoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleDataTable.GoogleChronicleDataTableColumnInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleChronicleDataTableColumnInfoOutputReference_Override(g GoogleChronicleDataTableColumnInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleDataTable.GoogleChronicleDataTableColumnInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetColumnIndex(val *float64) {
	if err := j.validateSetColumnIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetColumnType(val *string) {
	if err := j.validateSetColumnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnType",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetKeyColumn(val interface{}) {
	if err := j.validateSetKeyColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyColumn",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetMappedColumnPath(val *string) {
	if err := j.validateSetMappedColumnPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappedColumnPath",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetOriginalColumn(val *string) {
	if err := j.validateSetOriginalColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originalColumn",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetRepeatedValues(val interface{}) {
	if err := j.validateSetRepeatedValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeatedValues",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ResetColumnType() {
	_jsii_.InvokeVoid(
		g,
		"resetColumnType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ResetKeyColumn() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyColumn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ResetMappedColumnPath() {
	_jsii_.InvokeVoid(
		g,
		"resetMappedColumnPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ResetRepeatedValues() {
	_jsii_.InvokeVoid(
		g,
		"resetRepeatedValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDataTableColumnInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

