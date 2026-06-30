package sourcedataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/sourcedataset/internal"
)

type SourceDatasetFieldOutputReference interface {
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
	InternalValue() any
	SetInternalValue(val any)
	IsConst() any
	SetIsConst(val any)
	IsConstInput() any
	IsEnum() any
	SetIsEnum(val any)
	IsEnumInput() any
	IsHidden() any
	SetIsHidden(val any)
	IsHiddenInput() any
	IsMetric() any
	SetIsMetric(val any)
	IsMetricInput() any
	IsSearchable() any
	SetIsSearchable(val any)
	IsSearchableInput() any
	Name() *string
	SetName(val *string)
	NameInput() *string
	SqlType() *string
	SetSqlType(val *string)
	SqlTypeInput() *string
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
	ResetIsConst()
	ResetIsEnum()
	ResetIsHidden()
	ResetIsMetric()
	ResetIsSearchable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SourceDatasetFieldOutputReference
type jsiiProxy_SourceDatasetFieldOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsConst() any {
	var returns any
	_jsii_.Get(
		j,
		"isConst",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsConstInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isConstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsEnum() any {
	var returns any
	_jsii_.Get(
		j,
		"isEnum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsEnumInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isEnumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsHidden() any {
	var returns any
	_jsii_.Get(
		j,
		"isHidden",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsHiddenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isHiddenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsMetric() any {
	var returns any
	_jsii_.Get(
		j,
		"isMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsMetricInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsSearchable() any {
	var returns any
	_jsii_.Get(
		j,
		"isSearchable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) IsSearchableInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isSearchableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SqlType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SqlTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewSourceDatasetFieldOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SourceDatasetFieldOutputReference {
	_init_.Initialize()

	if err := validateNewSourceDatasetFieldOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SourceDatasetFieldOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.sourceDataset.SourceDatasetFieldOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSourceDatasetFieldOutputReference_Override(s SourceDatasetFieldOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.sourceDataset.SourceDatasetFieldOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetIsConst(val any) {
	if err := j.validateSetIsConstParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isConst",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetIsEnum(val any) {
	if err := j.validateSetIsEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnum",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetIsHidden(val any) {
	if err := j.validateSetIsHiddenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isHidden",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetIsMetric(val any) {
	if err := j.validateSetIsMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMetric",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetIsSearchable(val any) {
	if err := j.validateSetIsSearchableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSearchable",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetSqlType(val *string) {
	if err := j.validateSetSqlTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlType",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SourceDatasetFieldOutputReference) SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) ResetIsConst() {
	_jsii_.InvokeVoid(
		s,
		"resetIsConst",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) ResetIsEnum() {
	_jsii_.InvokeVoid(
		s,
		"resetIsEnum",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) ResetIsHidden() {
	_jsii_.InvokeVoid(
		s,
		"resetIsHidden",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) ResetIsMetric() {
	_jsii_.InvokeVoid(
		s,
		"resetIsMetric",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) ResetIsSearchable() {
	_jsii_.InvokeVoid(
		s,
		"resetIsSearchable",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		s,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceDatasetFieldOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
