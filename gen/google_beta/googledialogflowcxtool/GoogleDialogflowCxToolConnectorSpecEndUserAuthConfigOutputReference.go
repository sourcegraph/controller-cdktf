package googledialogflowcxtool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxtool/internal"
)

type GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfig
	SetInternalValue(val *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfig)
	Oauth2AuthCodeConfig() GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfigOutputReference
	Oauth2AuthCodeConfigInput() *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfig
	Oauth2JwtBearerConfig() GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfigOutputReference
	Oauth2JwtBearerConfigInput() *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfig
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
	PutOauth2AuthCodeConfig(value *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfig)
	PutOauth2JwtBearerConfig(value *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfig)
	ResetOauth2AuthCodeConfig()
	ResetOauth2JwtBearerConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference
type jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) InternalValue() *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfig {
	var returns *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) Oauth2AuthCodeConfig() GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfigOutputReference {
	var returns GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfigOutputReference
	_jsii_.Get(
		j,
		"oauth2AuthCodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) Oauth2AuthCodeConfigInput() *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfig {
	var returns *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfig
	_jsii_.Get(
		j,
		"oauth2AuthCodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) Oauth2JwtBearerConfig() GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfigOutputReference {
	var returns GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfigOutputReference
	_jsii_.Get(
		j,
		"oauth2JwtBearerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) Oauth2JwtBearerConfigInput() *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfig {
	var returns *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfig
	_jsii_.Get(
		j,
		"oauth2JwtBearerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxTool.GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference_Override(g GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxTool.GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference)SetInternalValue(val *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) PutOauth2AuthCodeConfig(value *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfig) {
	if err := g.validatePutOauth2AuthCodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2AuthCodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) PutOauth2JwtBearerConfig(value *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfig) {
	if err := g.validatePutOauth2JwtBearerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2JwtBearerConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) ResetOauth2AuthCodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2AuthCodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) ResetOauth2JwtBearerConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2JwtBearerConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

