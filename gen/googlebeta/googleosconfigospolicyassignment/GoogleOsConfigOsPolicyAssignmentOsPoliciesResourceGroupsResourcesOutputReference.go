package googleosconfigospolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleosconfigospolicyassignment/internal"
)

type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference interface {
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
	Exec() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference
	ExecInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	File() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileOutputReference
	FileInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Pkg() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference
	PkgInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	Repository() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference
	RepositoryInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
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
	PutExec(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec)
	PutFile(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile)
	PutPkg(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg)
	PutRepository(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository)
	ResetExec()
	ResetFile()
	ResetPkg()
	ResetRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference
type jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Exec() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ExecInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) File() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) FileInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Pkg() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference
	_jsii_.Get(
		j,
		"pkg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PkgInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	_jsii_.Get(
		j,
		"pkgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Repository() GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference {
	var returns GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) RepositoryInput() *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	var returns *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleOsConfigOsPolicyAssignment.GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference_Override(g GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleOsConfigOsPolicyAssignment.GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PutExec(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) {
	if err := g.validatePutExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PutFile(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) {
	if err := g.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PutPkg(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) {
	if err := g.validatePutPkgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPkg",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PutRepository(value *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) {
	if err := g.validatePutRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ResetExec() {
	_jsii_.InvokeVoid(
		g,
		"resetExec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		g,
		"resetFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ResetPkg() {
	_jsii_.InvokeVoid(
		g,
		"resetPkg",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ResetRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

