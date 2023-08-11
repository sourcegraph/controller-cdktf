package pagesproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/pagesproject/internal"
)

type PagesProjectDeploymentConfigsPreviewOutputReference interface {
	cdktf.ComplexObject
	AlwaysUseLatestCompatibilityDate() interface{}
	SetAlwaysUseLatestCompatibilityDate(val interface{})
	AlwaysUseLatestCompatibilityDateInput() interface{}
	CompatibilityDate() *string
	SetCompatibilityDate(val *string)
	CompatibilityDateInput() *string
	CompatibilityFlags() *[]*string
	SetCompatibilityFlags(val *[]*string)
	CompatibilityFlagsInput() *[]*string
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
	D1Databases() *map[string]*string
	SetD1Databases(val *map[string]*string)
	D1DatabasesInput() *map[string]*string
	DurableObjectNamespaces() *map[string]*string
	SetDurableObjectNamespaces(val *map[string]*string)
	DurableObjectNamespacesInput() *map[string]*string
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	FailOpen() interface{}
	SetFailOpen(val interface{})
	FailOpenInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *PagesProjectDeploymentConfigsPreview
	SetInternalValue(val *PagesProjectDeploymentConfigsPreview)
	KvNamespaces() *map[string]*string
	SetKvNamespaces(val *map[string]*string)
	KvNamespacesInput() *map[string]*string
	R2Buckets() *map[string]*string
	SetR2Buckets(val *map[string]*string)
	R2BucketsInput() *map[string]*string
	ServiceBinding() PagesProjectDeploymentConfigsPreviewServiceBindingList
	ServiceBindingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageModel() *string
	SetUsageModel(val *string)
	UsageModelInput() *string
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
	PutServiceBinding(value interface{})
	ResetAlwaysUseLatestCompatibilityDate()
	ResetCompatibilityDate()
	ResetCompatibilityFlags()
	ResetD1Databases()
	ResetDurableObjectNamespaces()
	ResetEnvironmentVariables()
	ResetFailOpen()
	ResetKvNamespaces()
	ResetR2Buckets()
	ResetServiceBinding()
	ResetUsageModel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PagesProjectDeploymentConfigsPreviewOutputReference
type jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) AlwaysUseLatestCompatibilityDate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysUseLatestCompatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) AlwaysUseLatestCompatibilityDateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysUseLatestCompatibilityDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CompatibilityDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CompatibilityDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CompatibilityFlags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CompatibilityFlagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) D1Databases() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"d1Databases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) D1DatabasesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"d1DatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) DurableObjectNamespaces() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"durableObjectNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) DurableObjectNamespacesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"durableObjectNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) FailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) FailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) InternalValue() *PagesProjectDeploymentConfigsPreview {
	var returns *PagesProjectDeploymentConfigsPreview
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) KvNamespaces() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"kvNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) KvNamespacesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"kvNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) R2Buckets() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"r2Buckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) R2BucketsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"r2BucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ServiceBinding() PagesProjectDeploymentConfigsPreviewServiceBindingList {
	var returns PagesProjectDeploymentConfigsPreviewServiceBindingList
	_jsii_.Get(
		j,
		"serviceBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ServiceBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) UsageModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) UsageModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageModelInput",
		&returns,
	)
	return returns
}


func NewPagesProjectDeploymentConfigsPreviewOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PagesProjectDeploymentConfigsPreviewOutputReference {
	_init_.Initialize()

	if err := validateNewPagesProjectDeploymentConfigsPreviewOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectDeploymentConfigsPreviewOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPagesProjectDeploymentConfigsPreviewOutputReference_Override(p PagesProjectDeploymentConfigsPreviewOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectDeploymentConfigsPreviewOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetAlwaysUseLatestCompatibilityDate(val interface{}) {
	if err := j.validateSetAlwaysUseLatestCompatibilityDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysUseLatestCompatibilityDate",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetCompatibilityDate(val *string) {
	if err := j.validateSetCompatibilityDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityDate",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetCompatibilityFlags(val *[]*string) {
	if err := j.validateSetCompatibilityFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityFlags",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetD1Databases(val *map[string]*string) {
	if err := j.validateSetD1DatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"d1Databases",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetDurableObjectNamespaces(val *map[string]*string) {
	if err := j.validateSetDurableObjectNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durableObjectNamespaces",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetFailOpen(val interface{}) {
	if err := j.validateSetFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOpen",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetInternalValue(val *PagesProjectDeploymentConfigsPreview) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetKvNamespaces(val *map[string]*string) {
	if err := j.validateSetKvNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kvNamespaces",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetR2Buckets(val *map[string]*string) {
	if err := j.validateSetR2BucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"r2Buckets",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference)SetUsageModel(val *string) {
	if err := j.validateSetUsageModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageModel",
		val,
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutServiceBinding(value interface{}) {
	if err := p.validatePutServiceBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putServiceBinding",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetAlwaysUseLatestCompatibilityDate() {
	_jsii_.InvokeVoid(
		p,
		"resetAlwaysUseLatestCompatibilityDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetCompatibilityDate() {
	_jsii_.InvokeVoid(
		p,
		"resetCompatibilityDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetCompatibilityFlags() {
	_jsii_.InvokeVoid(
		p,
		"resetCompatibilityFlags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetD1Databases() {
	_jsii_.InvokeVoid(
		p,
		"resetD1Databases",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetDurableObjectNamespaces() {
	_jsii_.InvokeVoid(
		p,
		"resetDurableObjectNamespaces",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetFailOpen() {
	_jsii_.InvokeVoid(
		p,
		"resetFailOpen",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetKvNamespaces() {
	_jsii_.InvokeVoid(
		p,
		"resetKvNamespaces",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetR2Buckets() {
	_jsii_.InvokeVoid(
		p,
		"resetR2Buckets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetServiceBinding() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceBinding",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetUsageModel() {
	_jsii_.InvokeVoid(
		p,
		"resetUsageModel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

