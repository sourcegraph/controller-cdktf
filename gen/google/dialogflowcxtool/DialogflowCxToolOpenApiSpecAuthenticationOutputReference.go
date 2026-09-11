package dialogflowcxtool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxtool/internal"
)

type DialogflowCxToolOpenApiSpecAuthenticationOutputReference interface {
	cdktn.ComplexObject
	ApiKeyConfig() DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfigOutputReference
	ApiKeyConfigInput() *DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfig
	BearerTokenConfig() DialogflowCxToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference
	BearerTokenConfigInput() *DialogflowCxToolOpenApiSpecAuthenticationBearerTokenConfig
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
	InternalValue() *DialogflowCxToolOpenApiSpecAuthentication
	SetInternalValue(val *DialogflowCxToolOpenApiSpecAuthentication)
	OauthConfig() DialogflowCxToolOpenApiSpecAuthenticationOauthConfigOutputReference
	OauthConfigInput() *DialogflowCxToolOpenApiSpecAuthenticationOauthConfig
	ServiceAgentAuthConfig() DialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference
	ServiceAgentAuthConfigInput() *DialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutApiKeyConfig(value *DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfig)
	PutBearerTokenConfig(value *DialogflowCxToolOpenApiSpecAuthenticationBearerTokenConfig)
	PutOauthConfig(value *DialogflowCxToolOpenApiSpecAuthenticationOauthConfig)
	PutServiceAgentAuthConfig(value *DialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfig)
	ResetApiKeyConfig()
	ResetBearerTokenConfig()
	ResetOauthConfig()
	ResetServiceAgentAuthConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxToolOpenApiSpecAuthenticationOutputReference
type jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ApiKeyConfig() DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfigOutputReference {
	var returns DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ApiKeyConfigInput() *DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfig {
	var returns *DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) BearerTokenConfig() DialogflowCxToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference {
	var returns DialogflowCxToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"bearerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) BearerTokenConfigInput() *DialogflowCxToolOpenApiSpecAuthenticationBearerTokenConfig {
	var returns *DialogflowCxToolOpenApiSpecAuthenticationBearerTokenConfig
	_jsii_.Get(
		j,
		"bearerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) InternalValue() *DialogflowCxToolOpenApiSpecAuthentication {
	var returns *DialogflowCxToolOpenApiSpecAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) OauthConfig() DialogflowCxToolOpenApiSpecAuthenticationOauthConfigOutputReference {
	var returns DialogflowCxToolOpenApiSpecAuthenticationOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) OauthConfigInput() *DialogflowCxToolOpenApiSpecAuthenticationOauthConfig {
	var returns *DialogflowCxToolOpenApiSpecAuthenticationOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ServiceAgentAuthConfig() DialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference {
	var returns DialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAgentAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ServiceAgentAuthConfigInput() *DialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfig {
	var returns *DialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfig
	_jsii_.Get(
		j,
		"serviceAgentAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowCxToolOpenApiSpecAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DialogflowCxToolOpenApiSpecAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxToolOpenApiSpecAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCxTool.DialogflowCxToolOpenApiSpecAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxToolOpenApiSpecAuthenticationOutputReference_Override(d DialogflowCxToolOpenApiSpecAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCxTool.DialogflowCxToolOpenApiSpecAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference)SetInternalValue(val *DialogflowCxToolOpenApiSpecAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) PutApiKeyConfig(value *DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfig) {
	if err := d.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) PutBearerTokenConfig(value *DialogflowCxToolOpenApiSpecAuthenticationBearerTokenConfig) {
	if err := d.validatePutBearerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBearerTokenConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) PutOauthConfig(value *DialogflowCxToolOpenApiSpecAuthenticationOauthConfig) {
	if err := d.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) PutServiceAgentAuthConfig(value *DialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfig) {
	if err := d.validatePutServiceAgentAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceAgentAuthConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ResetBearerTokenConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetBearerTokenConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ResetServiceAgentAuthConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAgentAuthConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolOpenApiSpecAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

