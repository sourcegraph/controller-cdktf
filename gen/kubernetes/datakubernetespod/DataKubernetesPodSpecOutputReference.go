package datakubernetespod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/datakubernetespod/internal"
)

type DataKubernetesPodSpecOutputReference interface {
	cdktn.ComplexObject
	ActiveDeadlineSeconds() *float64
	Affinity() DataKubernetesPodSpecAffinityList
	AutomountServiceAccountToken() cdktn.IResolvable
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
	Container() DataKubernetesPodSpecContainerList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsConfig() DataKubernetesPodSpecDnsConfigList
	DnsPolicy() *string
	EnableServiceLinks() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	HostAliases() DataKubernetesPodSpecHostAliasesList
	HostIpc() cdktn.IResolvable
	Hostname() *string
	HostNetwork() cdktn.IResolvable
	HostPid() cdktn.IResolvable
	ImagePullSecrets() DataKubernetesPodSpecImagePullSecretsList
	InitContainer() DataKubernetesPodSpecInitContainerList
	InternalValue() *DataKubernetesPodSpec
	SetInternalValue(val *DataKubernetesPodSpec)
	NodeName() *string
	NodeSelector() cdktn.StringMap
	PriorityClassName() *string
	ReadinessGate() DataKubernetesPodSpecReadinessGateList
	RestartPolicy() *string
	SecurityContext() DataKubernetesPodSpecSecurityContextList
	ServiceAccountName() *string
	ShareProcessNamespace() cdktn.IResolvable
	Subdomain() *string
	TerminationGracePeriodSeconds() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Toleration() DataKubernetesPodSpecTolerationList
	TopologySpreadConstraint() DataKubernetesPodSpecTopologySpreadConstraintList
	Volume() DataKubernetesPodSpecVolumeList
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataKubernetesPodSpecOutputReference
type jsiiProxy_DataKubernetesPodSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) Affinity() DataKubernetesPodSpecAffinityList {
	var returns DataKubernetesPodSpecAffinityList
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) AutomountServiceAccountToken() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) Container() DataKubernetesPodSpecContainerList {
	var returns DataKubernetesPodSpecContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) DnsConfig() DataKubernetesPodSpecDnsConfigList {
	var returns DataKubernetesPodSpecDnsConfigList
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) EnableServiceLinks() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableServiceLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) HostAliases() DataKubernetesPodSpecHostAliasesList {
	var returns DataKubernetesPodSpecHostAliasesList
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) HostIpc() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) HostNetwork() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) HostPid() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) ImagePullSecrets() DataKubernetesPodSpecImagePullSecretsList {
	var returns DataKubernetesPodSpecImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) InitContainer() DataKubernetesPodSpecInitContainerList {
	var returns DataKubernetesPodSpecInitContainerList
	_jsii_.Get(
		j,
		"initContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) InternalValue() *DataKubernetesPodSpec {
	var returns *DataKubernetesPodSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) NodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) NodeSelector() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) PriorityClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) ReadinessGate() DataKubernetesPodSpecReadinessGateList {
	var returns DataKubernetesPodSpecReadinessGateList
	_jsii_.Get(
		j,
		"readinessGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) SecurityContext() DataKubernetesPodSpecSecurityContextList {
	var returns DataKubernetesPodSpecSecurityContextList
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) ShareProcessNamespace() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) TerminationGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) Toleration() DataKubernetesPodSpecTolerationList {
	var returns DataKubernetesPodSpecTolerationList
	_jsii_.Get(
		j,
		"toleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) TopologySpreadConstraint() DataKubernetesPodSpecTopologySpreadConstraintList {
	var returns DataKubernetesPodSpecTopologySpreadConstraintList
	_jsii_.Get(
		j,
		"topologySpreadConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference) Volume() DataKubernetesPodSpecVolumeList {
	var returns DataKubernetesPodSpecVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}


func NewDataKubernetesPodSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataKubernetesPodSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataKubernetesPodSpecOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesPodSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-kubernetes.dataKubernetesPod.DataKubernetesPodSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataKubernetesPodSpecOutputReference_Override(d DataKubernetesPodSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-kubernetes.dataKubernetesPod.DataKubernetesPodSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference)SetInternalValue(val *DataKubernetesPodSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

