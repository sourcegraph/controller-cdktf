package googleosconfigguestpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigguestpolicies/internal"
)

type GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference interface {
	cdktf.ComplexObject
	ArchiveExtraction() GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtractionOutputReference
	ArchiveExtractionInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtraction
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
	DpkgInstallation() GoogleOsConfigGuestPoliciesRecipesUpdateStepsDpkgInstallationOutputReference
	DpkgInstallationInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsDpkgInstallation
	FileCopy() GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileCopyOutputReference
	FileCopyInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileCopy
	FileExec() GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExecOutputReference
	FileExecInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExec
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MsiInstallation() GoogleOsConfigGuestPoliciesRecipesUpdateStepsMsiInstallationOutputReference
	MsiInstallationInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsMsiInstallation
	RpmInstallation() GoogleOsConfigGuestPoliciesRecipesUpdateStepsRpmInstallationOutputReference
	RpmInstallationInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsRpmInstallation
	ScriptRun() GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRunOutputReference
	ScriptRunInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRun
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
	PutArchiveExtraction(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtraction)
	PutDpkgInstallation(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsDpkgInstallation)
	PutFileCopy(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileCopy)
	PutFileExec(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExec)
	PutMsiInstallation(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsMsiInstallation)
	PutRpmInstallation(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsRpmInstallation)
	PutScriptRun(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRun)
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

// The jsii proxy struct for GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference
type jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ArchiveExtraction() GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtractionOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtractionOutputReference
	_jsii_.Get(
		j,
		"archiveExtraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ArchiveExtractionInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtraction {
	var returns *GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtraction
	_jsii_.Get(
		j,
		"archiveExtractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) DpkgInstallation() GoogleOsConfigGuestPoliciesRecipesUpdateStepsDpkgInstallationOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesUpdateStepsDpkgInstallationOutputReference
	_jsii_.Get(
		j,
		"dpkgInstallation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) DpkgInstallationInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsDpkgInstallation {
	var returns *GoogleOsConfigGuestPoliciesRecipesUpdateStepsDpkgInstallation
	_jsii_.Get(
		j,
		"dpkgInstallationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) FileCopy() GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileCopyOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileCopyOutputReference
	_jsii_.Get(
		j,
		"fileCopy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) FileCopyInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileCopy {
	var returns *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileCopy
	_jsii_.Get(
		j,
		"fileCopyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) FileExec() GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExecOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExecOutputReference
	_jsii_.Get(
		j,
		"fileExec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) FileExecInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExec {
	var returns *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExec
	_jsii_.Get(
		j,
		"fileExecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) MsiInstallation() GoogleOsConfigGuestPoliciesRecipesUpdateStepsMsiInstallationOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesUpdateStepsMsiInstallationOutputReference
	_jsii_.Get(
		j,
		"msiInstallation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) MsiInstallationInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsMsiInstallation {
	var returns *GoogleOsConfigGuestPoliciesRecipesUpdateStepsMsiInstallation
	_jsii_.Get(
		j,
		"msiInstallationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) RpmInstallation() GoogleOsConfigGuestPoliciesRecipesUpdateStepsRpmInstallationOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesUpdateStepsRpmInstallationOutputReference
	_jsii_.Get(
		j,
		"rpmInstallation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) RpmInstallationInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsRpmInstallation {
	var returns *GoogleOsConfigGuestPoliciesRecipesUpdateStepsRpmInstallation
	_jsii_.Get(
		j,
		"rpmInstallationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ScriptRun() GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRunOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRunOutputReference
	_jsii_.Get(
		j,
		"scriptRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ScriptRunInput() *GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRun {
	var returns *GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRun
	_jsii_.Get(
		j,
		"scriptRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference{}

	_jsii_.Create(
		"google-beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference_Override(g GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) PutArchiveExtraction(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtraction) {
	if err := g.validatePutArchiveExtractionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putArchiveExtraction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) PutDpkgInstallation(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsDpkgInstallation) {
	if err := g.validatePutDpkgInstallationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDpkgInstallation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) PutFileCopy(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileCopy) {
	if err := g.validatePutFileCopyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFileCopy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) PutFileExec(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExec) {
	if err := g.validatePutFileExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFileExec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) PutMsiInstallation(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsMsiInstallation) {
	if err := g.validatePutMsiInstallationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMsiInstallation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) PutRpmInstallation(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsRpmInstallation) {
	if err := g.validatePutRpmInstallationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRpmInstallation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) PutScriptRun(value *GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRun) {
	if err := g.validatePutScriptRunParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScriptRun",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ResetArchiveExtraction() {
	_jsii_.InvokeVoid(
		g,
		"resetArchiveExtraction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ResetDpkgInstallation() {
	_jsii_.InvokeVoid(
		g,
		"resetDpkgInstallation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ResetFileCopy() {
	_jsii_.InvokeVoid(
		g,
		"resetFileCopy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ResetFileExec() {
	_jsii_.InvokeVoid(
		g,
		"resetFileExec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ResetMsiInstallation() {
	_jsii_.InvokeVoid(
		g,
		"resetMsiInstallation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ResetRpmInstallation() {
	_jsii_.InvokeVoid(
		g,
		"resetRpmInstallation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ResetScriptRun() {
	_jsii_.InvokeVoid(
		g,
		"resetScriptRun",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesUpdateStepsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

