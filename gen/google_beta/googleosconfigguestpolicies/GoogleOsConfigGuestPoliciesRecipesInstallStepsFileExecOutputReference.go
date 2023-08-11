package googleosconfigguestpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigguestpolicies/internal"
)

type GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference interface {
	cdktf.ComplexObject
	AllowedExitCodes() *string
	SetAllowedExitCodes(val *string)
	AllowedExitCodesInput() *string
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	ArtifactId() *string
	SetArtifactId(val *string)
	ArtifactIdInput() *string
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
	InternalValue() *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec
	SetInternalValue(val *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec)
	LocalPath() *string
	SetLocalPath(val *string)
	LocalPathInput() *string
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
	ResetAllowedExitCodes()
	ResetArgs()
	ResetArtifactId()
	ResetLocalPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference
type jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) AllowedExitCodes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedExitCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) AllowedExitCodesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedExitCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ArtifactIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) InternalValue() *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec {
	var returns *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference_Override(g GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference)SetAllowedExitCodes(val *string) {
	if err := j.validateSetAllowedExitCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedExitCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference)SetArtifactId(val *string) {
	if err := j.validateSetArtifactIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifactId",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference)SetInternalValue(val *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ResetAllowedExitCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedExitCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ResetArtifactId() {
	_jsii_.InvokeVoid(
		g,
		"resetArtifactId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ResetLocalPath() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

