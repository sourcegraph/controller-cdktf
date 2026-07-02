package deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/deployment/internal"
)

type DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConfigMap() DeploymentSpecTemplateSpecVolumeProjectedSourcesConfigMapList
	ConfigMapInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DownwardApi() DeploymentSpecTemplateSpecVolumeProjectedSourcesDownwardApiOutputReference
	DownwardApiInput() *DeploymentSpecTemplateSpecVolumeProjectedSourcesDownwardApi
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Secret() DeploymentSpecTemplateSpecVolumeProjectedSourcesSecretList
	SecretInput() any
	ServiceAccountToken() DeploymentSpecTemplateSpecVolumeProjectedSourcesServiceAccountTokenOutputReference
	ServiceAccountTokenInput() *DeploymentSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutConfigMap(value any)
	PutDownwardApi(value *DeploymentSpecTemplateSpecVolumeProjectedSourcesDownwardApi)
	PutSecret(value any)
	PutServiceAccountToken(value *DeploymentSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken)
	ResetConfigMap()
	ResetDownwardApi()
	ResetSecret()
	ResetServiceAccountToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference
type jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ConfigMap() DeploymentSpecTemplateSpecVolumeProjectedSourcesConfigMapList {
	var returns DeploymentSpecTemplateSpecVolumeProjectedSourcesConfigMapList
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ConfigMapInput() any {
	var returns any
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) DownwardApi() DeploymentSpecTemplateSpecVolumeProjectedSourcesDownwardApiOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeProjectedSourcesDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) DownwardApiInput() *DeploymentSpecTemplateSpecVolumeProjectedSourcesDownwardApi {
	var returns *DeploymentSpecTemplateSpecVolumeProjectedSourcesDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) Secret() DeploymentSpecTemplateSpecVolumeProjectedSourcesSecretList {
	var returns DeploymentSpecTemplateSpecVolumeProjectedSourcesSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) SecretInput() any {
	var returns any
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ServiceAccountToken() DeploymentSpecTemplateSpecVolumeProjectedSourcesServiceAccountTokenOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeProjectedSourcesServiceAccountTokenOutputReference
	_jsii_.Get(
		j,
		"serviceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ServiceAccountTokenInput() *DeploymentSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken {
	var returns *DeploymentSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken
	_jsii_.Get(
		j,
		"serviceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference_Override(d DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) PutConfigMap(value any) {
	if err := d.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfigMap",
		[]any{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) PutDownwardApi(value *DeploymentSpecTemplateSpecVolumeProjectedSourcesDownwardApi) {
	if err := d.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDownwardApi",
		[]any{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) PutSecret(value any) {
	if err := d.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecret",
		[]any{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) PutServiceAccountToken(value *DeploymentSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken) {
	if err := d.validatePutServiceAccountTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceAccountToken",
		[]any{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		d,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		d,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetServiceAccountToken() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccountToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeProjectedSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
