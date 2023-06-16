package googledataprocjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocjob/internal"
)

type GoogleDataprocJobPrestoConfigOutputReference interface {
	cdktf.ComplexObject
	ClientTags() *[]*string
	SetClientTags(val *[]*string)
	ClientTagsInput() *[]*string
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
	ContinueOnFailure() interface{}
	SetContinueOnFailure(val interface{})
	ContinueOnFailureInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataprocJobPrestoConfig
	SetInternalValue(val *GoogleDataprocJobPrestoConfig)
	LoggingConfig() GoogleDataprocJobPrestoConfigLoggingConfigOutputReference
	LoggingConfigInput() *GoogleDataprocJobPrestoConfigLoggingConfig
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	QueryFileUri() *string
	SetQueryFileUri(val *string)
	QueryFileUriInput() *string
	QueryList() *[]*string
	SetQueryList(val *[]*string)
	QueryListInput() *[]*string
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
	PutLoggingConfig(value *GoogleDataprocJobPrestoConfigLoggingConfig)
	ResetClientTags()
	ResetContinueOnFailure()
	ResetLoggingConfig()
	ResetOutputFormat()
	ResetProperties()
	ResetQueryFileUri()
	ResetQueryList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocJobPrestoConfigOutputReference
type jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ClientTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ClientTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ContinueOnFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ContinueOnFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) InternalValue() *GoogleDataprocJobPrestoConfig {
	var returns *GoogleDataprocJobPrestoConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) LoggingConfig() GoogleDataprocJobPrestoConfigLoggingConfigOutputReference {
	var returns GoogleDataprocJobPrestoConfigLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) LoggingConfigInput() *GoogleDataprocJobPrestoConfigLoggingConfig {
	var returns *GoogleDataprocJobPrestoConfigLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) QueryFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) QueryFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) QueryList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) QueryListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocJobPrestoConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocJobPrestoConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocJobPrestoConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocJob.GoogleDataprocJobPrestoConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocJobPrestoConfigOutputReference_Override(g GoogleDataprocJobPrestoConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocJob.GoogleDataprocJobPrestoConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetClientTags(val *[]*string) {
	if err := j.validateSetClientTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTags",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetContinueOnFailure(val interface{}) {
	if err := j.validateSetContinueOnFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continueOnFailure",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetInternalValue(val *GoogleDataprocJobPrestoConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetQueryFileUri(val *string) {
	if err := j.validateSetQueryFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFileUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetQueryList(val *[]*string) {
	if err := j.validateSetQueryListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryList",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) PutLoggingConfig(value *GoogleDataprocJobPrestoConfigLoggingConfig) {
	if err := g.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ResetClientTags() {
	_jsii_.InvokeVoid(
		g,
		"resetClientTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ResetContinueOnFailure() {
	_jsii_.InvokeVoid(
		g,
		"resetContinueOnFailure",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ResetQueryFileUri() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryFileUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ResetQueryList() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocJobPrestoConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

