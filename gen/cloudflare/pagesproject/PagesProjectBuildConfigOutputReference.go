package pagesproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/pagesproject/internal"
)

type PagesProjectBuildConfigOutputReference interface {
	cdktf.ComplexObject
	BuildCommand() *string
	SetBuildCommand(val *string)
	BuildCommandInput() *string
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
	DestinationDir() *string
	SetDestinationDir(val *string)
	DestinationDirInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *PagesProjectBuildConfig
	SetInternalValue(val *PagesProjectBuildConfig)
	RootDir() *string
	SetRootDir(val *string)
	RootDirInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebAnalyticsTag() *string
	SetWebAnalyticsTag(val *string)
	WebAnalyticsTagInput() *string
	WebAnalyticsToken() *string
	SetWebAnalyticsToken(val *string)
	WebAnalyticsTokenInput() *string
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
	ResetBuildCommand()
	ResetDestinationDir()
	ResetRootDir()
	ResetWebAnalyticsTag()
	ResetWebAnalyticsToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PagesProjectBuildConfigOutputReference
type jsiiProxy_PagesProjectBuildConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) BuildCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) BuildCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) DestinationDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) DestinationDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) InternalValue() *PagesProjectBuildConfig {
	var returns *PagesProjectBuildConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) RootDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) RootDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) WebAnalyticsTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAnalyticsTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) WebAnalyticsTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAnalyticsTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) WebAnalyticsToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAnalyticsToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference) WebAnalyticsTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAnalyticsTokenInput",
		&returns,
	)
	return returns
}


func NewPagesProjectBuildConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PagesProjectBuildConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPagesProjectBuildConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PagesProjectBuildConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectBuildConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPagesProjectBuildConfigOutputReference_Override(p PagesProjectBuildConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectBuildConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetBuildCommand(val *string) {
	if err := j.validateSetBuildCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildCommand",
		val,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetDestinationDir(val *string) {
	if err := j.validateSetDestinationDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationDir",
		val,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetInternalValue(val *PagesProjectBuildConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetRootDir(val *string) {
	if err := j.validateSetRootDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootDir",
		val,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetWebAnalyticsTag(val *string) {
	if err := j.validateSetWebAnalyticsTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAnalyticsTag",
		val,
	)
}

func (j *jsiiProxy_PagesProjectBuildConfigOutputReference)SetWebAnalyticsToken(val *string) {
	if err := j.validateSetWebAnalyticsTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAnalyticsToken",
		val,
	)
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) ResetBuildCommand() {
	_jsii_.InvokeVoid(
		p,
		"resetBuildCommand",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) ResetDestinationDir() {
	_jsii_.InvokeVoid(
		p,
		"resetDestinationDir",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) ResetRootDir() {
	_jsii_.InvokeVoid(
		p,
		"resetRootDir",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) ResetWebAnalyticsTag() {
	_jsii_.InvokeVoid(
		p,
		"resetWebAnalyticsTag",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) ResetWebAnalyticsToken() {
	_jsii_.InvokeVoid(
		p,
		"resetWebAnalyticsToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectBuildConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

