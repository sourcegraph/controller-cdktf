package teamsrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/teamsrule/internal"
)

type TeamsRuleRuleSettingsBisoAdminControlsOutputReference interface {
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
	DisableCopyPaste() interface{}
	SetDisableCopyPaste(val interface{})
	DisableCopyPasteInput() interface{}
	DisableDownload() interface{}
	SetDisableDownload(val interface{})
	DisableDownloadInput() interface{}
	DisableKeyboard() interface{}
	SetDisableKeyboard(val interface{})
	DisableKeyboardInput() interface{}
	DisablePrinting() interface{}
	SetDisablePrinting(val interface{})
	DisablePrintingInput() interface{}
	DisableUpload() interface{}
	SetDisableUpload(val interface{})
	DisableUploadInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *TeamsRuleRuleSettingsBisoAdminControls
	SetInternalValue(val *TeamsRuleRuleSettingsBisoAdminControls)
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
	ResetDisableCopyPaste()
	ResetDisableDownload()
	ResetDisableKeyboard()
	ResetDisablePrinting()
	ResetDisableUpload()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamsRuleRuleSettingsBisoAdminControlsOutputReference
type jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisableCopyPaste() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCopyPaste",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisableCopyPasteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCopyPasteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisableDownload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDownload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisableDownloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDownloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisableKeyboard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableKeyboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisableKeyboardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableKeyboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisablePrinting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrinting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisablePrintingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrintingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisableUpload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) DisableUploadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) InternalValue() *TeamsRuleRuleSettingsBisoAdminControls {
	var returns *TeamsRuleRuleSettingsBisoAdminControls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTeamsRuleRuleSettingsBisoAdminControlsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamsRuleRuleSettingsBisoAdminControlsOutputReference {
	_init_.Initialize()

	if err := validateNewTeamsRuleRuleSettingsBisoAdminControlsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsRule.TeamsRuleRuleSettingsBisoAdminControlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamsRuleRuleSettingsBisoAdminControlsOutputReference_Override(t TeamsRuleRuleSettingsBisoAdminControlsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsRule.TeamsRuleRuleSettingsBisoAdminControlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetDisableCopyPaste(val interface{}) {
	if err := j.validateSetDisableCopyPasteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableCopyPaste",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetDisableDownload(val interface{}) {
	if err := j.validateSetDisableDownloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDownload",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetDisableKeyboard(val interface{}) {
	if err := j.validateSetDisableKeyboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableKeyboard",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetDisablePrinting(val interface{}) {
	if err := j.validateSetDisablePrintingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePrinting",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetDisableUpload(val interface{}) {
	if err := j.validateSetDisableUploadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUpload",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetInternalValue(val *TeamsRuleRuleSettingsBisoAdminControls) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) ResetDisableCopyPaste() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableCopyPaste",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) ResetDisableDownload() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableDownload",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) ResetDisableKeyboard() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableKeyboard",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) ResetDisablePrinting() {
	_jsii_.InvokeVoid(
		t,
		"resetDisablePrinting",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) ResetDisableUpload() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableUpload",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsBisoAdminControlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

