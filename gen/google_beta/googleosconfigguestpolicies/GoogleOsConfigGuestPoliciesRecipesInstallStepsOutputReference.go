package googleosconfigguestpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigguestpolicies/internal"
)

type GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference interface {
	cdktf.ComplexObject
	ArchiveExtraction() GoogleOsConfigGuestPoliciesRecipesInstallStepsArchiveExtractionOutputReference
	ArchiveExtractionInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsArchiveExtraction
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
	DpkgInstallation() GoogleOsConfigGuestPoliciesRecipesInstallStepsDpkgInstallationOutputReference
	DpkgInstallationInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsDpkgInstallation
	FileCopy() GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopyOutputReference
	FileCopyInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopy
	FileExec() GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference
	FileExecInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MsiInstallation() GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallationOutputReference
	MsiInstallationInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallation
	RpmInstallation() GoogleOsConfigGuestPoliciesRecipesInstallStepsRpmInstallationOutputReference
	RpmInstallationInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsRpmInstallation
	ScriptRun() GoogleOsConfigGuestPoliciesRecipesInstallStepsScriptRunOutputReference
	ScriptRunInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsScriptRun
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
	PutArchiveExtraction(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsArchiveExtraction)
	PutDpkgInstallation(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsDpkgInstallation)
	PutFileCopy(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopy)
	PutFileExec(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec)
	PutMsiInstallation(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallation)
	PutRpmInstallation(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsRpmInstallation)
	PutScriptRun(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsScriptRun)
	ResetArchiveExtraction()
	ResetDpkgInstallation()
	ResetFileCopy()
	ResetFileExec()
	ResetMsiInstallation()
	ResetRpmInstallation()
	ResetScriptRun()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference
type jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ArchiveExtraction() GoogleOsConfigGuestPoliciesRecipesInstallStepsArchiveExtractionOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesInstallStepsArchiveExtractionOutputReference
	_jsii_.Get(
		j,
		"archiveExtraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ArchiveExtractionInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsArchiveExtraction {
	var returns *GoogleOsConfigGuestPoliciesRecipesInstallStepsArchiveExtraction
	_jsii_.Get(
		j,
		"archiveExtractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) DpkgInstallation() GoogleOsConfigGuestPoliciesRecipesInstallStepsDpkgInstallationOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesInstallStepsDpkgInstallationOutputReference
	_jsii_.Get(
		j,
		"dpkgInstallation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) DpkgInstallationInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsDpkgInstallation {
	var returns *GoogleOsConfigGuestPoliciesRecipesInstallStepsDpkgInstallation
	_jsii_.Get(
		j,
		"dpkgInstallationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) FileCopy() GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopyOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopyOutputReference
	_jsii_.Get(
		j,
		"fileCopy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) FileCopyInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopy {
	var returns *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopy
	_jsii_.Get(
		j,
		"fileCopyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) FileExec() GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference
	_jsii_.Get(
		j,
		"fileExec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) FileExecInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec {
	var returns *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec
	_jsii_.Get(
		j,
		"fileExecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) MsiInstallation() GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallationOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallationOutputReference
	_jsii_.Get(
		j,
		"msiInstallation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) MsiInstallationInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallation {
	var returns *GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallation
	_jsii_.Get(
		j,
		"msiInstallationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) RpmInstallation() GoogleOsConfigGuestPoliciesRecipesInstallStepsRpmInstallationOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesInstallStepsRpmInstallationOutputReference
	_jsii_.Get(
		j,
		"rpmInstallation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) RpmInstallationInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsRpmInstallation {
	var returns *GoogleOsConfigGuestPoliciesRecipesInstallStepsRpmInstallation
	_jsii_.Get(
		j,
		"rpmInstallationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ScriptRun() GoogleOsConfigGuestPoliciesRecipesInstallStepsScriptRunOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesInstallStepsScriptRunOutputReference
	_jsii_.Get(
		j,
		"scriptRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ScriptRunInput() *GoogleOsConfigGuestPoliciesRecipesInstallStepsScriptRun {
	var returns *GoogleOsConfigGuestPoliciesRecipesInstallStepsScriptRun
	_jsii_.Get(
		j,
		"scriptRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference_Override(g GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) PutArchiveExtraction(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsArchiveExtraction) {
	if err := g.validatePutArchiveExtractionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putArchiveExtraction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) PutDpkgInstallation(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsDpkgInstallation) {
	if err := g.validatePutDpkgInstallationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDpkgInstallation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) PutFileCopy(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopy) {
	if err := g.validatePutFileCopyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFileCopy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) PutFileExec(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec) {
	if err := g.validatePutFileExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFileExec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) PutMsiInstallation(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallation) {
	if err := g.validatePutMsiInstallationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMsiInstallation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) PutRpmInstallation(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsRpmInstallation) {
	if err := g.validatePutRpmInstallationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRpmInstallation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) PutScriptRun(value *GoogleOsConfigGuestPoliciesRecipesInstallStepsScriptRun) {
	if err := g.validatePutScriptRunParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScriptRun",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ResetArchiveExtraction() {
	_jsii_.InvokeVoid(
		g,
		"resetArchiveExtraction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ResetDpkgInstallation() {
	_jsii_.InvokeVoid(
		g,
		"resetDpkgInstallation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ResetFileCopy() {
	_jsii_.InvokeVoid(
		g,
		"resetFileCopy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ResetFileExec() {
	_jsii_.InvokeVoid(
		g,
		"resetFileExec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ResetMsiInstallation() {
	_jsii_.InvokeVoid(
		g,
		"resetMsiInstallation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ResetRpmInstallation() {
	_jsii_.InvokeVoid(
		g,
		"resetRpmInstallation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ResetScriptRun() {
	_jsii_.InvokeVoid(
		g,
		"resetScriptRun",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

