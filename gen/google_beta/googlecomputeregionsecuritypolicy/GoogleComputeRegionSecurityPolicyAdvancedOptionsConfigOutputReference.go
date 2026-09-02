package googlecomputeregionsecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionsecuritypolicy/internal"
)

type GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference interface {
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
	InternalValue() *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfig
	SetInternalValue(val *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfig)
	JsonCustomConfig() GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference
	JsonCustomConfigInput() *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig
	JsonParsing() *string
	SetJsonParsing(val *string)
	JsonParsingInput() *string
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	RequestBodyInspectionSize() *string
	SetRequestBodyInspectionSize(val *string)
	RequestBodyInspectionSizeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserIpRequestHeaders() *[]*string
	SetUserIpRequestHeaders(val *[]*string)
	UserIpRequestHeadersInput() *[]*string
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
	PutJsonCustomConfig(value *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig)
	ResetJsonCustomConfig()
	ResetJsonParsing()
	ResetLogLevel()
	ResetRequestBodyInspectionSize()
	ResetUserIpRequestHeaders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference
type jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) InternalValue() *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfig {
	var returns *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) JsonCustomConfig() GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference {
	var returns GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference
	_jsii_.Get(
		j,
		"jsonCustomConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) JsonCustomConfigInput() *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig {
	var returns *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig
	_jsii_.Get(
		j,
		"jsonCustomConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) JsonParsing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParsing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) JsonParsingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParsingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) RequestBodyInspectionSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInspectionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) RequestBodyInspectionSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInspectionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) UserIpRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIpRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) UserIpRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIpRequestHeadersInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionSecurityPolicy.GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference_Override(g GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionSecurityPolicy.GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetInternalValue(val *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetJsonParsing(val *string) {
	if err := j.validateSetJsonParsingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonParsing",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetRequestBodyInspectionSize(val *string) {
	if err := j.validateSetRequestBodyInspectionSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBodyInspectionSize",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetUserIpRequestHeaders(val *[]*string) {
	if err := j.validateSetUserIpRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIpRequestHeaders",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) PutJsonCustomConfig(value *GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig) {
	if err := g.validatePutJsonCustomConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJsonCustomConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ResetJsonCustomConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonCustomConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ResetJsonParsing() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonParsing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ResetRequestBodyInspectionSize() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestBodyInspectionSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ResetUserIpRequestHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetUserIpRequestHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

