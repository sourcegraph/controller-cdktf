package dialogflowcxtoolversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxtoolversion/internal"
)

type DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference interface {
	cdktf.ComplexObject
	ApiKeyConfig() DialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfigOutputReference
	ApiKeyConfigInput() *DialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig
	BearerTokenConfig() DialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference
	BearerTokenConfigInput() *DialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig
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
	InternalValue() *DialogflowCxToolVersionToolOpenApiSpecAuthentication
	SetInternalValue(val *DialogflowCxToolVersionToolOpenApiSpecAuthentication)
	OauthConfig() DialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfigOutputReference
	OauthConfigInput() *DialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig
	ServiceAgentAuthConfig() DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference
	ServiceAgentAuthConfigInput() *DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig
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
	PutApiKeyConfig(value *DialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig)
	PutBearerTokenConfig(value *DialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig)
	PutOauthConfig(value *DialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig)
	PutServiceAgentAuthConfig(value *DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig)
	ResetApiKeyConfig()
	ResetBearerTokenConfig()
	ResetOauthConfig()
	ResetServiceAgentAuthConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference
type jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ApiKeyConfig() DialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfigOutputReference {
	var returns DialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ApiKeyConfigInput() *DialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig {
	var returns *DialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) BearerTokenConfig() DialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference {
	var returns DialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"bearerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) BearerTokenConfigInput() *DialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig {
	var returns *DialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig
	_jsii_.Get(
		j,
		"bearerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) InternalValue() *DialogflowCxToolVersionToolOpenApiSpecAuthentication {
	var returns *DialogflowCxToolVersionToolOpenApiSpecAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) OauthConfig() DialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfigOutputReference {
	var returns DialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) OauthConfigInput() *DialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig {
	var returns *DialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ServiceAgentAuthConfig() DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference {
	var returns DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAgentAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ServiceAgentAuthConfigInput() *DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig {
	var returns *DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig
	_jsii_.Get(
		j,
		"serviceAgentAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxToolVersion.DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference_Override(d DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxToolVersion.DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) SetInternalValue(val *DialogflowCxToolVersionToolOpenApiSpecAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) PutApiKeyConfig(value *DialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig) {
	if err := d.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApiKeyConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) PutBearerTokenConfig(value *DialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig) {
	if err := d.validatePutBearerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBearerTokenConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) PutOauthConfig(value *DialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig) {
	if err := d.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOauthConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) PutServiceAgentAuthConfig(value *DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig) {
	if err := d.validatePutServiceAgentAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceAgentAuthConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ResetBearerTokenConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetBearerTokenConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ResetServiceAgentAuthConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAgentAuthConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
