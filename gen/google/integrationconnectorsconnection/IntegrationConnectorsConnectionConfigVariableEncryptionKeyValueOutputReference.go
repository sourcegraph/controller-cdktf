package integrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/integrationconnectorsconnection/internal"
)

type IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference interface {
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
	InternalValue() *IntegrationConnectorsConnectionConfigVariableEncryptionKeyValue
	SetInternalValue(val *IntegrationConnectorsConnectionConfigVariableEncryptionKeyValue)
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
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
	ResetKmsKeyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference
type jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) InternalValue() *IntegrationConnectorsConnectionConfigVariableEncryptionKeyValue {
	var returns *IntegrationConnectorsConnectionConfigVariableEncryptionKeyValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewIntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference_Override(i IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference",
		[]any{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) SetInternalValue(val *IntegrationConnectorsConnectionConfigVariableEncryptionKeyValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		i,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		i,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableEncryptionKeyValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
