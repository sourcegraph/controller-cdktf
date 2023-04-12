package googlebigquerydataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlebigquerydataset/internal"
)

type GoogleBigqueryDatasetAccessOutputReference interface {
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
	Dataset() GoogleBigqueryDatasetAccessDatasetOutputReference
	DatasetInput() *GoogleBigqueryDatasetAccessDataset
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	Fqn() *string
	GroupByEmail() *string
	SetGroupByEmail(val *string)
	GroupByEmailInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	SpecialGroup() *string
	SetSpecialGroup(val *string)
	SpecialGroupInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserByEmail() *string
	SetUserByEmail(val *string)
	UserByEmailInput() *string
	View() GoogleBigqueryDatasetAccessViewOutputReference
	ViewInput() *GoogleBigqueryDatasetAccessView
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
	PutDataset(value *GoogleBigqueryDatasetAccessDataset)
	PutView(value *GoogleBigqueryDatasetAccessView)
	ResetDataset()
	ResetDomain()
	ResetGroupByEmail()
	ResetRole()
	ResetSpecialGroup()
	ResetUserByEmail()
	ResetView()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryDatasetAccessOutputReference
type jsiiProxy_GoogleBigqueryDatasetAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) Dataset() GoogleBigqueryDatasetAccessDatasetOutputReference {
	var returns GoogleBigqueryDatasetAccessDatasetOutputReference
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) DatasetInput() *GoogleBigqueryDatasetAccessDataset {
	var returns *GoogleBigqueryDatasetAccessDataset
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GroupByEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GroupByEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) SpecialGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specialGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) SpecialGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specialGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) UserByEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userByEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) UserByEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userByEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) View() GoogleBigqueryDatasetAccessViewOutputReference {
	var returns GoogleBigqueryDatasetAccessViewOutputReference
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ViewInput() *GoogleBigqueryDatasetAccessView {
	var returns *GoogleBigqueryDatasetAccessView
	_jsii_.Get(
		j,
		"viewInput",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryDatasetAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleBigqueryDatasetAccessOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryDatasetAccessOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryDatasetAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBigqueryDataset.GoogleBigqueryDatasetAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleBigqueryDatasetAccessOutputReference_Override(g GoogleBigqueryDatasetAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBigqueryDataset.GoogleBigqueryDatasetAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetGroupByEmail(val *string) {
	if err := j.validateSetGroupByEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupByEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetSpecialGroup(val *string) {
	if err := j.validateSetSpecialGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"specialGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference)SetUserByEmail(val *string) {
	if err := j.validateSetUserByEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userByEmail",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) PutDataset(value *GoogleBigqueryDatasetAccessDataset) {
	if err := g.validatePutDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataset",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) PutView(value *GoogleBigqueryDatasetAccessView) {
	if err := g.validatePutViewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putView",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ResetDataset() {
	_jsii_.InvokeVoid(
		g,
		"resetDataset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ResetGroupByEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupByEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ResetRole() {
	_jsii_.InvokeVoid(
		g,
		"resetRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ResetSpecialGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetSpecialGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ResetUserByEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetUserByEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ResetView() {
	_jsii_.InvokeVoid(
		g,
		"resetView",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

