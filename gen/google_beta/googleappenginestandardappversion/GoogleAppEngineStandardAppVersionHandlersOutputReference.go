package googleappenginestandardappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleappenginestandardappversion/internal"
)

type GoogleAppEngineStandardAppVersionHandlersOutputReference interface {
	cdktf.ComplexObject
	AuthFailAction() *string
	SetAuthFailAction(val *string)
	AuthFailActionInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Login() *string
	SetLogin(val *string)
	LoginInput() *string
	RedirectHttpResponseCode() *string
	SetRedirectHttpResponseCode(val *string)
	RedirectHttpResponseCodeInput() *string
	Script() GoogleAppEngineStandardAppVersionHandlersScriptOutputReference
	ScriptInput() *GoogleAppEngineStandardAppVersionHandlersScript
	SecurityLevel() *string
	SetSecurityLevel(val *string)
	SecurityLevelInput() *string
	StaticFiles() GoogleAppEngineStandardAppVersionHandlersStaticFilesOutputReference
	StaticFilesInput() *GoogleAppEngineStandardAppVersionHandlersStaticFiles
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRegex() *string
	SetUrlRegex(val *string)
	UrlRegexInput() *string
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
	PutScript(value *GoogleAppEngineStandardAppVersionHandlersScript)
	PutStaticFiles(value *GoogleAppEngineStandardAppVersionHandlersStaticFiles)
	ResetAuthFailAction()
	ResetLogin()
	ResetRedirectHttpResponseCode()
	ResetScript()
	ResetSecurityLevel()
	ResetStaticFiles()
	ResetUrlRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAppEngineStandardAppVersionHandlersOutputReference
type jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) AuthFailAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authFailAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) AuthFailActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authFailActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) Login() *string {
	var returns *string
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) LoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) RedirectHttpResponseCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectHttpResponseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) RedirectHttpResponseCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectHttpResponseCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) Script() GoogleAppEngineStandardAppVersionHandlersScriptOutputReference {
	var returns GoogleAppEngineStandardAppVersionHandlersScriptOutputReference
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ScriptInput() *GoogleAppEngineStandardAppVersionHandlersScript {
	var returns *GoogleAppEngineStandardAppVersionHandlersScript
	_jsii_.Get(
		j,
		"scriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) SecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) StaticFiles() GoogleAppEngineStandardAppVersionHandlersStaticFilesOutputReference {
	var returns GoogleAppEngineStandardAppVersionHandlersStaticFilesOutputReference
	_jsii_.Get(
		j,
		"staticFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) StaticFilesInput() *GoogleAppEngineStandardAppVersionHandlersStaticFiles {
	var returns *GoogleAppEngineStandardAppVersionHandlersStaticFiles
	_jsii_.Get(
		j,
		"staticFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) UrlRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) UrlRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlRegexInput",
		&returns,
	)
	return returns
}


func NewGoogleAppEngineStandardAppVersionHandlersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleAppEngineStandardAppVersionHandlersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineStandardAppVersionHandlersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersionHandlersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleAppEngineStandardAppVersionHandlersOutputReference_Override(g GoogleAppEngineStandardAppVersionHandlersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersionHandlersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetAuthFailAction(val *string) {
	if err := j.validateSetAuthFailActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authFailAction",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetLogin(val *string) {
	if err := j.validateSetLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"login",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetRedirectHttpResponseCode(val *string) {
	if err := j.validateSetRedirectHttpResponseCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectHttpResponseCode",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetSecurityLevel(val *string) {
	if err := j.validateSetSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference)SetUrlRegex(val *string) {
	if err := j.validateSetUrlRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlRegex",
		val,
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) PutScript(value *GoogleAppEngineStandardAppVersionHandlersScript) {
	if err := g.validatePutScriptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScript",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) PutStaticFiles(value *GoogleAppEngineStandardAppVersionHandlersStaticFiles) {
	if err := g.validatePutStaticFilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStaticFiles",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ResetAuthFailAction() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthFailAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ResetLogin() {
	_jsii_.InvokeVoid(
		g,
		"resetLogin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ResetRedirectHttpResponseCode() {
	_jsii_.InvokeVoid(
		g,
		"resetRedirectHttpResponseCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ResetScript() {
	_jsii_.InvokeVoid(
		g,
		"resetScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ResetSecurityLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ResetStaticFiles() {
	_jsii_.InvokeVoid(
		g,
		"resetStaticFiles",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ResetUrlRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionHandlersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

