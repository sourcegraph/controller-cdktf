package googleintegrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleintegrationconnectorsconnection/internal"
)

type GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference interface {
	cdktf.ComplexObject
	BooleanValue() interface{}
	SetBooleanValue(val interface{})
	BooleanValueInput() interface{}
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
	EncryptionKeyValue() GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValueOutputReference
	EncryptionKeyValueInput() *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValue
	// Experimental.
	Fqn() *string
	IntegerValue() *float64
	SetIntegerValue(val *float64)
	IntegerValueInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	SecretValue() GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableSecretValueOutputReference
	SecretValueInput() *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableSecretValue
	StringValue() *string
	SetStringValue(val *string)
	StringValueInput() *string
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
	PutEncryptionKeyValue(value *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValue)
	PutSecretValue(value *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableSecretValue)
	ResetBooleanValue()
	ResetEncryptionKeyValue()
	ResetIntegerValue()
	ResetSecretValue()
	ResetStringValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference
type jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) BooleanValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) BooleanValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) EncryptionKeyValue() GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValueOutputReference {
	var returns GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValueOutputReference
	_jsii_.Get(
		j,
		"encryptionKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) EncryptionKeyValueInput() *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValue {
	var returns *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValue
	_jsii_.Get(
		j,
		"encryptionKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) IntegerValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"integerValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) IntegerValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"integerValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) SecretValue() GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableSecretValueOutputReference {
	var returns GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableSecretValueOutputReference
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) SecretValueInput() *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableSecretValue {
	var returns *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableSecretValue
	_jsii_.Get(
		j,
		"secretValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference_Override(g GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference)SetBooleanValue(val interface{}) {
	if err := j.validateSetBooleanValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"booleanValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference)SetIntegerValue(val *float64) {
	if err := j.validateSetIntegerValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integerValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) PutEncryptionKeyValue(value *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValue) {
	if err := g.validatePutEncryptionKeyValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionKeyValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) PutSecretValue(value *GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableSecretValue) {
	if err := g.validatePutSecretValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecretValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) ResetBooleanValue() {
	_jsii_.InvokeVoid(
		g,
		"resetBooleanValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) ResetEncryptionKeyValue() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKeyValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) ResetIntegerValue() {
	_jsii_.InvokeVoid(
		g,
		"resetIntegerValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) ResetSecretValue() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		g,
		"resetStringValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionEventingConfigAdditionalVariableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

