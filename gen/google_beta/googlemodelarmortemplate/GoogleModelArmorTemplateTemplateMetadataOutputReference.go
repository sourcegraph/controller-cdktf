package googlemodelarmortemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemodelarmortemplate/internal"
)

type GoogleModelArmorTemplateTemplateMetadataOutputReference interface {
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
	CustomLlmResponseSafetyErrorCode() *float64
	SetCustomLlmResponseSafetyErrorCode(val *float64)
	CustomLlmResponseSafetyErrorCodeInput() *float64
	CustomLlmResponseSafetyErrorMessage() *string
	SetCustomLlmResponseSafetyErrorMessage(val *string)
	CustomLlmResponseSafetyErrorMessageInput() *string
	CustomPromptSafetyErrorCode() *float64
	SetCustomPromptSafetyErrorCode(val *float64)
	CustomPromptSafetyErrorCodeInput() *float64
	CustomPromptSafetyErrorMessage() *string
	SetCustomPromptSafetyErrorMessage(val *string)
	CustomPromptSafetyErrorMessageInput() *string
	EnforcementType() *string
	SetEnforcementType(val *string)
	EnforcementTypeInput() *string
	// Experimental.
	Fqn() *string
	IgnorePartialInvocationFailures() interface{}
	SetIgnorePartialInvocationFailures(val interface{})
	IgnorePartialInvocationFailuresInput() interface{}
	InternalValue() *GoogleModelArmorTemplateTemplateMetadata
	SetInternalValue(val *GoogleModelArmorTemplateTemplateMetadata)
	LogSanitizeOperations() interface{}
	SetLogSanitizeOperations(val interface{})
	LogSanitizeOperationsInput() interface{}
	LogTemplateOperations() interface{}
	SetLogTemplateOperations(val interface{})
	LogTemplateOperationsInput() interface{}
	MultiLanguageDetection() GoogleModelArmorTemplateTemplateMetadataMultiLanguageDetectionOutputReference
	MultiLanguageDetectionInput() *GoogleModelArmorTemplateTemplateMetadataMultiLanguageDetection
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
	PutMultiLanguageDetection(value *GoogleModelArmorTemplateTemplateMetadataMultiLanguageDetection)
	ResetCustomLlmResponseSafetyErrorCode()
	ResetCustomLlmResponseSafetyErrorMessage()
	ResetCustomPromptSafetyErrorCode()
	ResetCustomPromptSafetyErrorMessage()
	ResetEnforcementType()
	ResetIgnorePartialInvocationFailures()
	ResetLogSanitizeOperations()
	ResetLogTemplateOperations()
	ResetMultiLanguageDetection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleModelArmorTemplateTemplateMetadataOutputReference
type jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) CustomLlmResponseSafetyErrorCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customLlmResponseSafetyErrorCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) CustomLlmResponseSafetyErrorCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customLlmResponseSafetyErrorCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) CustomLlmResponseSafetyErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLlmResponseSafetyErrorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) CustomLlmResponseSafetyErrorMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLlmResponseSafetyErrorMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) CustomPromptSafetyErrorCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customPromptSafetyErrorCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) CustomPromptSafetyErrorCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customPromptSafetyErrorCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) CustomPromptSafetyErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPromptSafetyErrorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) CustomPromptSafetyErrorMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPromptSafetyErrorMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) EnforcementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) EnforcementTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) IgnorePartialInvocationFailures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePartialInvocationFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) IgnorePartialInvocationFailuresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePartialInvocationFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) InternalValue() *GoogleModelArmorTemplateTemplateMetadata {
	var returns *GoogleModelArmorTemplateTemplateMetadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) LogSanitizeOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logSanitizeOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) LogSanitizeOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logSanitizeOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) LogTemplateOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logTemplateOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) LogTemplateOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logTemplateOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) MultiLanguageDetection() GoogleModelArmorTemplateTemplateMetadataMultiLanguageDetectionOutputReference {
	var returns GoogleModelArmorTemplateTemplateMetadataMultiLanguageDetectionOutputReference
	_jsii_.Get(
		j,
		"multiLanguageDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) MultiLanguageDetectionInput() *GoogleModelArmorTemplateTemplateMetadataMultiLanguageDetection {
	var returns *GoogleModelArmorTemplateTemplateMetadataMultiLanguageDetection
	_jsii_.Get(
		j,
		"multiLanguageDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleModelArmorTemplateTemplateMetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleModelArmorTemplateTemplateMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleModelArmorTemplateTemplateMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleModelArmorTemplate.GoogleModelArmorTemplateTemplateMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleModelArmorTemplateTemplateMetadataOutputReference_Override(g GoogleModelArmorTemplateTemplateMetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleModelArmorTemplate.GoogleModelArmorTemplateTemplateMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetCustomLlmResponseSafetyErrorCode(val *float64) {
	if err := j.validateSetCustomLlmResponseSafetyErrorCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLlmResponseSafetyErrorCode",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetCustomLlmResponseSafetyErrorMessage(val *string) {
	if err := j.validateSetCustomLlmResponseSafetyErrorMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLlmResponseSafetyErrorMessage",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetCustomPromptSafetyErrorCode(val *float64) {
	if err := j.validateSetCustomPromptSafetyErrorCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPromptSafetyErrorCode",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetCustomPromptSafetyErrorMessage(val *string) {
	if err := j.validateSetCustomPromptSafetyErrorMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPromptSafetyErrorMessage",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetEnforcementType(val *string) {
	if err := j.validateSetEnforcementTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcementType",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetIgnorePartialInvocationFailures(val interface{}) {
	if err := j.validateSetIgnorePartialInvocationFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignorePartialInvocationFailures",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetInternalValue(val *GoogleModelArmorTemplateTemplateMetadata) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetLogSanitizeOperations(val interface{}) {
	if err := j.validateSetLogSanitizeOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logSanitizeOperations",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetLogTemplateOperations(val interface{}) {
	if err := j.validateSetLogTemplateOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logTemplateOperations",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) PutMultiLanguageDetection(value *GoogleModelArmorTemplateTemplateMetadataMultiLanguageDetection) {
	if err := g.validatePutMultiLanguageDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMultiLanguageDetection",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ResetCustomLlmResponseSafetyErrorCode() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomLlmResponseSafetyErrorCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ResetCustomLlmResponseSafetyErrorMessage() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomLlmResponseSafetyErrorMessage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ResetCustomPromptSafetyErrorCode() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomPromptSafetyErrorCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ResetCustomPromptSafetyErrorMessage() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomPromptSafetyErrorMessage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ResetEnforcementType() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforcementType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ResetIgnorePartialInvocationFailures() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnorePartialInvocationFailures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ResetLogSanitizeOperations() {
	_jsii_.InvokeVoid(
		g,
		"resetLogSanitizeOperations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ResetLogTemplateOperations() {
	_jsii_.InvokeVoid(
		g,
		"resetLogTemplateOperations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ResetMultiLanguageDetection() {
	_jsii_.InvokeVoid(
		g,
		"resetMultiLanguageDetection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleModelArmorTemplateTemplateMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

