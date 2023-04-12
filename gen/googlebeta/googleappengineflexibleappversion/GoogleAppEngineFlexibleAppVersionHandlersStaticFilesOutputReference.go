package googleappengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleappengineflexibleappversion/internal"
)

type GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference interface {
	cdktf.ComplexObject
	ApplicationReadable() interface{}
	SetApplicationReadable(val interface{})
	ApplicationReadableInput() interface{}
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
	Expiration() *string
	SetExpiration(val *string)
	ExpirationInput() *string
	// Experimental.
	Fqn() *string
	HttpHeaders() *map[string]*string
	SetHttpHeaders(val *map[string]*string)
	HttpHeadersInput() *map[string]*string
	InternalValue() *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles
	SetInternalValue(val *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles)
	MimeType() *string
	SetMimeType(val *string)
	MimeTypeInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	RequireMatchingFile() interface{}
	SetRequireMatchingFile(val interface{})
	RequireMatchingFileInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UploadPathRegex() *string
	SetUploadPathRegex(val *string)
	UploadPathRegexInput() *string
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
	ResetApplicationReadable()
	ResetExpiration()
	ResetHttpHeaders()
	ResetMimeType()
	ResetPath()
	ResetRequireMatchingFile()
	ResetUploadPathRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference
type jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ApplicationReadable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationReadable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ApplicationReadableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationReadableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) Expiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ExpirationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) HttpHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) HttpHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) InternalValue() *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles {
	var returns *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) MimeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) MimeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) RequireMatchingFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMatchingFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) RequireMatchingFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMatchingFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) UploadPathRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadPathRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) UploadPathRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadPathRegexInput",
		&returns,
	)
	return returns
}


func NewGoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference_Override(g GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetApplicationReadable(val interface{}) {
	if err := j.validateSetApplicationReadableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationReadable",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetExpiration(val *string) {
	if err := j.validateSetExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiration",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetHttpHeaders(val *map[string]*string) {
	if err := j.validateSetHttpHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetInternalValue(val *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetMimeType(val *string) {
	if err := j.validateSetMimeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mimeType",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetRequireMatchingFile(val interface{}) {
	if err := j.validateSetRequireMatchingFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireMatchingFile",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetUploadPathRegex(val *string) {
	if err := j.validateSetUploadPathRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadPathRegex",
		val,
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetApplicationReadable() {
	_jsii_.InvokeVoid(
		g,
		"resetApplicationReadable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		g,
		"resetExpiration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetHttpHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetMimeType() {
	_jsii_.InvokeVoid(
		g,
		"resetMimeType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		g,
		"resetPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetRequireMatchingFile() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireMatchingFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetUploadPathRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetUploadPathRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

