package managedkafkaconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/managedkafkaconnector/internal"
)

type ManagedKafkaConnectorTaskRestartPolicyOutputReference interface {
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
	InternalValue() *ManagedKafkaConnectorTaskRestartPolicy
	SetInternalValue(val *ManagedKafkaConnectorTaskRestartPolicy)
	MaximumBackoff() *string
	SetMaximumBackoff(val *string)
	MaximumBackoffInput() *string
	MinimumBackoff() *string
	SetMinimumBackoff(val *string)
	MinimumBackoffInput() *string
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
	ResetMaximumBackoff()
	ResetMinimumBackoff()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedKafkaConnectorTaskRestartPolicyOutputReference
type jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) InternalValue() *ManagedKafkaConnectorTaskRestartPolicy {
	var returns *ManagedKafkaConnectorTaskRestartPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) MaximumBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) MaximumBackoffInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumBackoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) MinimumBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) MinimumBackoffInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumBackoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewManagedKafkaConnectorTaskRestartPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedKafkaConnectorTaskRestartPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewManagedKafkaConnectorTaskRestartPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.managedKafkaConnector.ManagedKafkaConnectorTaskRestartPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedKafkaConnectorTaskRestartPolicyOutputReference_Override(m ManagedKafkaConnectorTaskRestartPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.managedKafkaConnector.ManagedKafkaConnectorTaskRestartPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) SetInternalValue(val *ManagedKafkaConnectorTaskRestartPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) SetMaximumBackoff(val *string) {
	if err := j.validateSetMaximumBackoffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBackoff",
		val,
	)
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) SetMinimumBackoff(val *string) {
	if err := j.validateSetMinimumBackoffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumBackoff",
		val,
	)
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) ResetMaximumBackoff() {
	_jsii_.InvokeVoid(
		m,
		"resetMaximumBackoff",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) ResetMinimumBackoff() {
	_jsii_.InvokeVoid(
		m,
		"resetMinimumBackoff",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		m,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKafkaConnectorTaskRestartPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
