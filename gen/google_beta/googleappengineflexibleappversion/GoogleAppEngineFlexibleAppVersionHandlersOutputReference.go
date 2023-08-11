package googleappengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleappengineflexibleappversion/internal"
)

type GoogleAppEngineFlexibleAppVersionHandlersOutputReference interface {
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
	Script() GoogleAppEngineFlexibleAppVersionHandlersScriptOutputReference
	ScriptInput() *GoogleAppEngineFlexibleAppVersionHandlersScript
	SecurityLevel() *string
	SetSecurityLevel(val *string)
	SecurityLevelInput() *string
	StaticFiles() GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference
	StaticFilesInput() *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles
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
	PutScript(value *GoogleAppEngineFlexibleAppVersionHandlersScript)
	PutStaticFiles(value *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles)
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

// The jsii proxy struct for GoogleAppEngineFlexibleAppVersionHandlersOutputReference
type jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) AuthFailAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authFailAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) AuthFailActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authFailActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) Login() *string {
	var returns *string
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) LoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) RedirectHttpResponseCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectHttpResponseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) RedirectHttpResponseCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectHttpResponseCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) Script() GoogleAppEngineFlexibleAppVersionHandlersScriptOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionHandlersScriptOutputReference
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ScriptInput() *GoogleAppEngineFlexibleAppVersionHandlersScript {
	var returns *GoogleAppEngineFlexibleAppVersionHandlersScript
	_jsii_.Get(
		j,
		"scriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) SecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) StaticFiles() GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference
	_jsii_.Get(
		j,
		"staticFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) StaticFilesInput() *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles {
	var returns *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles
	_jsii_.Get(
		j,
		"staticFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) UrlRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) UrlRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlRegexInput",
		&returns,
	)
	return returns
}


func NewGoogleAppEngineFlexibleAppVersionHandlersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleAppEngineFlexibleAppVersionHandlersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineFlexibleAppVersionHandlersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionHandlersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleAppEngineFlexibleAppVersionHandlersOutputReference_Override(g GoogleAppEngineFlexibleAppVersionHandlersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionHandlersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetAuthFailAction(val *string) {
	if err := j.validateSetAuthFailActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authFailAction",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetLogin(val *string) {
	if err := j.validateSetLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"login",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetRedirectHttpResponseCode(val *string) {
	if err := j.validateSetRedirectHttpResponseCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectHttpResponseCode",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetSecurityLevel(val *string) {
	if err := j.validateSetSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference)SetUrlRegex(val *string) {
	if err := j.validateSetUrlRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlRegex",
		val,
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) PutScript(value *GoogleAppEngineFlexibleAppVersionHandlersScript) {
	if err := g.validatePutScriptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScript",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) PutStaticFiles(value *GoogleAppEngineFlexibleAppVersionHandlersStaticFiles) {
	if err := g.validatePutStaticFilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStaticFiles",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ResetAuthFailAction() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthFailAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ResetLogin() {
	_jsii_.InvokeVoid(
		g,
		"resetLogin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ResetRedirectHttpResponseCode() {
	_jsii_.InvokeVoid(
		g,
		"resetRedirectHttpResponseCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ResetScript() {
	_jsii_.InvokeVoid(
		g,
		"resetScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ResetSecurityLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ResetStaticFiles() {
	_jsii_.InvokeVoid(
		g,
		"resetStaticFiles",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ResetUrlRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionHandlersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

