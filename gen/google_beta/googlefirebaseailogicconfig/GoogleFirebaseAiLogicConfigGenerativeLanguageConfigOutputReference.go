package googlefirebaseailogicconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlefirebaseailogicconfig/internal"
)

type GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference interface {
	cdktn.ComplexObject
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
	ApiKeyWo() *string
	SetApiKeyWo(val *string)
	ApiKeyWoInput() *string
	ApiKeyWoVersion() *string
	SetApiKeyWoVersion(val *string)
	ApiKeyWoVersionInput() *string
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
	InternalValue() *GoogleFirebaseAiLogicConfigGenerativeLanguageConfig
	SetInternalValue(val *GoogleFirebaseAiLogicConfigGenerativeLanguageConfig)
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
	ResetApiKey()
	ResetApiKeyWo()
	ResetApiKeyWoVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference
type jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ApiKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ApiKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ApiKeyWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ApiKeyWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) InternalValue() *GoogleFirebaseAiLogicConfigGenerativeLanguageConfig {
	var returns *GoogleFirebaseAiLogicConfigGenerativeLanguageConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleFirebaseAiLogicConfig.GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference_Override(g GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleFirebaseAiLogicConfig.GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference)SetApiKey(val *string) {
	if err := j.validateSetApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference)SetApiKeyWo(val *string) {
	if err := j.validateSetApiKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeyWo",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference)SetApiKeyWoVersion(val *string) {
	if err := j.validateSetApiKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference)SetInternalValue(val *GoogleFirebaseAiLogicConfigGenerativeLanguageConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ResetApiKeyWo() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeyWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ResetApiKeyWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeyWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleFirebaseAiLogicConfigGenerativeLanguageConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

