package googlealloydbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlealloydbinstance/internal"
)

type GoogleAlloydbInstanceQueryInsightsConfigOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *GoogleAlloydbInstanceQueryInsightsConfig
	SetInternalValue(val *GoogleAlloydbInstanceQueryInsightsConfig)
	QueryPlansPerMinute() *float64
	SetQueryPlansPerMinute(val *float64)
	QueryPlansPerMinuteInput() *float64
	QueryStringLength() *float64
	SetQueryStringLength(val *float64)
	QueryStringLengthInput() *float64
	RecordApplicationTags() interface{}
	SetRecordApplicationTags(val interface{})
	RecordApplicationTagsInput() interface{}
	RecordClientAddress() interface{}
	SetRecordClientAddress(val interface{})
	RecordClientAddressInput() interface{}
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
	ResetQueryPlansPerMinute()
	ResetQueryStringLength()
	ResetRecordApplicationTags()
	ResetRecordClientAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAlloydbInstanceQueryInsightsConfigOutputReference
type jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) InternalValue() *GoogleAlloydbInstanceQueryInsightsConfig {
	var returns *GoogleAlloydbInstanceQueryInsightsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) QueryPlansPerMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryPlansPerMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) QueryPlansPerMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryPlansPerMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) QueryStringLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryStringLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) QueryStringLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryStringLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) RecordApplicationTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordApplicationTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) RecordApplicationTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordApplicationTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) RecordClientAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordClientAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) RecordClientAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordClientAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAlloydbInstanceQueryInsightsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAlloydbInstanceQueryInsightsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAlloydbInstanceQueryInsightsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstanceQueryInsightsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAlloydbInstanceQueryInsightsConfigOutputReference_Override(g GoogleAlloydbInstanceQueryInsightsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstanceQueryInsightsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference)SetInternalValue(val *GoogleAlloydbInstanceQueryInsightsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference)SetQueryPlansPerMinute(val *float64) {
	if err := j.validateSetQueryPlansPerMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryPlansPerMinute",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference)SetQueryStringLength(val *float64) {
	if err := j.validateSetQueryStringLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringLength",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference)SetRecordApplicationTags(val interface{}) {
	if err := j.validateSetRecordApplicationTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordApplicationTags",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference)SetRecordClientAddress(val interface{}) {
	if err := j.validateSetRecordClientAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordClientAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) ResetQueryPlansPerMinute() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryPlansPerMinute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) ResetQueryStringLength() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryStringLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) ResetRecordApplicationTags() {
	_jsii_.InvokeVoid(
		g,
		"resetRecordApplicationTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) ResetRecordClientAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetRecordClientAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAlloydbInstanceQueryInsightsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

