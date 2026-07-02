package spotinstancerequest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/spotinstancerequest/internal"
)

type SpotInstanceRequestCapacityReservationSpecificationOutputReference interface {
	cdktf.ComplexObject
	CapacityReservationPreference() *string
	SetCapacityReservationPreference(val *string)
	CapacityReservationPreferenceInput() *string
	CapacityReservationTarget() SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTargetOutputReference
	CapacityReservationTargetInput() *SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTarget
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
	InternalValue() *SpotInstanceRequestCapacityReservationSpecification
	SetInternalValue(val *SpotInstanceRequestCapacityReservationSpecification)
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
	PutCapacityReservationTarget(value *SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTarget)
	ResetCapacityReservationPreference()
	ResetCapacityReservationTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpotInstanceRequestCapacityReservationSpecificationOutputReference
type jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) CapacityReservationPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) CapacityReservationPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) CapacityReservationTarget() SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTargetOutputReference {
	var returns SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTargetOutputReference
	_jsii_.Get(
		j,
		"capacityReservationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) CapacityReservationTargetInput() *SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTarget {
	var returns *SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTarget
	_jsii_.Get(
		j,
		"capacityReservationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) InternalValue() *SpotInstanceRequestCapacityReservationSpecification {
	var returns *SpotInstanceRequestCapacityReservationSpecification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSpotInstanceRequestCapacityReservationSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpotInstanceRequestCapacityReservationSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewSpotInstanceRequestCapacityReservationSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.spotInstanceRequest.SpotInstanceRequestCapacityReservationSpecificationOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpotInstanceRequestCapacityReservationSpecificationOutputReference_Override(s SpotInstanceRequestCapacityReservationSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.spotInstanceRequest.SpotInstanceRequestCapacityReservationSpecificationOutputReference",
		[]any{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) SetCapacityReservationPreference(val *string) {
	if err := j.validateSetCapacityReservationPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationPreference",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) SetInternalValue(val *SpotInstanceRequestCapacityReservationSpecification) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) PutCapacityReservationTarget(value *SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTarget) {
	if err := s.validatePutCapacityReservationTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCapacityReservationTarget",
		[]any{value},
	)
}

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) ResetCapacityReservationPreference() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacityReservationPreference",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) ResetCapacityReservationTarget() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacityReservationTarget",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (s *jsiiProxy_SpotInstanceRequestCapacityReservationSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
