package googleosconfigospolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigospolicyassignment/internal"
)

type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference interface {
	cdktf.ComplexObject
	Apt() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptOutputReference
	AptInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt
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
	Deb() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebOutputReference
	DebInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	// Experimental.
	Fqn() *string
	Googet() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetOutputReference
	GoogetInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget
	InternalValue() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	SetInternalValue(val *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg)
	Msi() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiOutputReference
	MsiInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi
	Rpm() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmOutputReference
	RpmInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Yum() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumOutputReference
	YumInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum
	Zypper() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperOutputReference
	ZypperInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper
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
	PutApt(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt)
	PutDeb(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb)
	PutGooget(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget)
	PutMsi(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi)
	PutRpm(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm)
	PutYum(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum)
	PutZypper(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper)
	ResetApt()
	ResetDeb()
	ResetGooget()
	ResetMsi()
	ResetRpm()
	ResetYum()
	ResetZypper()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference
type jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) Apt() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptOutputReference
	_jsii_.Get(
		j,
		"apt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) AptInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt
	_jsii_.Get(
		j,
		"aptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) Deb() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebOutputReference
	_jsii_.Get(
		j,
		"deb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) DebInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb
	_jsii_.Get(
		j,
		"debInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) Googet() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetOutputReference
	_jsii_.Get(
		j,
		"googet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GoogetInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget
	_jsii_.Get(
		j,
		"googetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) InternalValue() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) Msi() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiOutputReference
	_jsii_.Get(
		j,
		"msi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) MsiInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi
	_jsii_.Get(
		j,
		"msiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) Rpm() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmOutputReference
	_jsii_.Get(
		j,
		"rpm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) RpmInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm
	_jsii_.Get(
		j,
		"rpmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) Yum() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumOutputReference
	_jsii_.Get(
		j,
		"yum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) YumInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum
	_jsii_.Get(
		j,
		"yumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) Zypper() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperOutputReference
	_jsii_.Get(
		j,
		"zypper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ZypperInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper
	_jsii_.Get(
		j,
		"zypperInput",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference{}

	_jsii_.Create(
		"google-beta.googleOsConfigOsPolicyAssignment.GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference_Override(g GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleOsConfigOsPolicyAssignment.GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference)SetInternalValue(val *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) PutApt(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) {
	if err := g.validatePutAptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApt",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) PutDeb(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) {
	if err := g.validatePutDebParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeb",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) PutGooget(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) {
	if err := g.validatePutGoogetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGooget",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) PutMsi(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) {
	if err := g.validatePutMsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMsi",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) PutRpm(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) {
	if err := g.validatePutRpmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRpm",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) PutYum(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) {
	if err := g.validatePutYumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putYum",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) PutZypper(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) {
	if err := g.validatePutZypperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putZypper",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetApt() {
	_jsii_.InvokeVoid(
		g,
		"resetApt",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetDeb() {
	_jsii_.InvokeVoid(
		g,
		"resetDeb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetGooget() {
	_jsii_.InvokeVoid(
		g,
		"resetGooget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetMsi() {
	_jsii_.InvokeVoid(
		g,
		"resetMsi",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetRpm() {
	_jsii_.InvokeVoid(
		g,
		"resetRpm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetYum() {
	_jsii_.InvokeVoid(
		g,
		"resetYum",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetZypper() {
	_jsii_.InvokeVoid(
		g,
		"resetZypper",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

