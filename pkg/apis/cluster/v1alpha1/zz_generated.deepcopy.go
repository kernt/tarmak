// +build !ignore_autogenerated

// Copyright Jetstack Ltd. See LICENSE for details.

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmazonESProxy) DeepCopyInto(out *AmazonESProxy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmazonESProxy.
func (in *AmazonESProxy) DeepCopy() *AmazonESProxy {
	if in == nil {
		return nil
	}
	out := new(AmazonESProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.InstancePools != nil {
		in, out := &in.InstancePools, &out.InstancePools
		*out = make([]InstancePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		if *in == nil {
			*out = nil
		} else {
			*out = new(Network)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.LoggingSinks != nil {
		in, out := &in.LoggingSinks, &out.LoggingSinks
		*out = make([]*LoggingSink, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(LoggingSink)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		if *in == nil {
			*out = nil
		} else {
			*out = new(Values)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.KubernetesAPI != nil {
		in, out := &in.KubernetesAPI, &out.KubernetesAPI
		if *in == nil {
			*out = nil
		} else {
			*out = new(KubernetesAPI)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterKubernetes)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Amazon != nil {
		in, out := &in.Amazon, &out.Amazon
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterAmazon)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAmazon) DeepCopyInto(out *ClusterAmazon) {
	*out = *in
	if in.AdditionalIAMPolicies != nil {
		in, out := &in.AdditionalIAMPolicies, &out.AdditionalIAMPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAmazon.
func (in *ClusterAmazon) DeepCopy() *ClusterAmazon {
	if in == nil {
		return nil
	}
	out := new(ClusterAmazon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetes) DeepCopyInto(out *ClusterKubernetes) {
	*out = *in
	if in.ClusterAutoscaler != nil {
		in, out := &in.ClusterAutoscaler, &out.ClusterAutoscaler
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterKubernetesClusterAutoscaler)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tiller != nil {
		in, out := &in.Tiller, &out.Tiller
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterKubernetesTiller)
			**out = **in
		}
	}
	if in.Dashboard != nil {
		in, out := &in.Dashboard, &out.Dashboard
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterKubernetesDashboard)
			**out = **in
		}
	}
	if in.APIServer != nil {
		in, out := &in.APIServer, &out.APIServer
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterKubernetesAPIServer)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.PodSecurityPolicy != nil {
		in, out := &in.PodSecurityPolicy, &out.PodSecurityPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterPodSecurityPolicy)
			**out = **in
		}
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterKubernetesPrometheus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetes.
func (in *ClusterKubernetes) DeepCopy() *ClusterKubernetes {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesAPIServer) DeepCopyInto(out *ClusterKubernetesAPIServer) {
	*out = *in
	if in.AllowCIDRs != nil {
		in, out := &in.AllowCIDRs, &out.AllowCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterKubernetesAPIServerOIDC)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesAPIServer.
func (in *ClusterKubernetesAPIServer) DeepCopy() *ClusterKubernetesAPIServer {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesAPIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesAPIServerOIDC) DeepCopyInto(out *ClusterKubernetesAPIServerOIDC) {
	*out = *in
	if in.SigningAlgs != nil {
		in, out := &in.SigningAlgs, &out.SigningAlgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesAPIServerOIDC.
func (in *ClusterKubernetesAPIServerOIDC) DeepCopy() *ClusterKubernetesAPIServerOIDC {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesAPIServerOIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesClusterAutoscaler) DeepCopyInto(out *ClusterKubernetesClusterAutoscaler) {
	*out = *in
	if in.Overprovisioning != nil {
		in, out := &in.Overprovisioning, &out.Overprovisioning
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterKubernetesClusterAutoscalerOverprovisioning)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesClusterAutoscaler.
func (in *ClusterKubernetesClusterAutoscaler) DeepCopy() *ClusterKubernetesClusterAutoscaler {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesClusterAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesClusterAutoscalerOverprovisioning) DeepCopyInto(out *ClusterKubernetesClusterAutoscalerOverprovisioning) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesClusterAutoscalerOverprovisioning.
func (in *ClusterKubernetesClusterAutoscalerOverprovisioning) DeepCopy() *ClusterKubernetesClusterAutoscalerOverprovisioning {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesClusterAutoscalerOverprovisioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesDashboard) DeepCopyInto(out *ClusterKubernetesDashboard) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesDashboard.
func (in *ClusterKubernetesDashboard) DeepCopy() *ClusterKubernetesDashboard {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesPrometheus) DeepCopyInto(out *ClusterKubernetesPrometheus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesPrometheus.
func (in *ClusterKubernetesPrometheus) DeepCopy() *ClusterKubernetesPrometheus {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesPrometheus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesTiller) DeepCopyInto(out *ClusterKubernetesTiller) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesTiller.
func (in *ClusterKubernetesTiller) DeepCopy() *ClusterKubernetesTiller {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesTiller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPodSecurityPolicy) DeepCopyInto(out *ClusterPodSecurityPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPodSecurityPolicy.
func (in *ClusterPodSecurityPolicy) DeepCopy() *ClusterPodSecurityPolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterPodSecurityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressRule) DeepCopyInto(out *EgressRule) {
	*out = *in
	in.Shared.DeepCopyInto(&out.Shared)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressRule.
func (in *EgressRule) DeepCopy() *EgressRule {
	if in == nil {
		return nil
	}
	out := new(EgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Firewall) DeepCopyInto(out *Firewall) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.IngressRules != nil {
		in, out := &in.IngressRules, &out.IngressRules
		*out = make([]*IngressRule, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(IngressRule)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.EgressRules != nil {
		in, out := &in.EgressRules, &out.EgressRules
		*out = make([]*EgressRule, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(EgressRule)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Firewall.
func (in *Firewall) DeepCopy() *Firewall {
	if in == nil {
		return nil
	}
	out := new(Firewall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPBasicAuth) DeepCopyInto(out *HTTPBasicAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPBasicAuth.
func (in *HTTPBasicAuth) DeepCopy() *HTTPBasicAuth {
	if in == nil {
		return nil
	}
	out := new(HTTPBasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	*out = *in
	in.Shared.DeepCopyInto(&out.Shared)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRule.
func (in *IngressRule) DeepCopy() *IngressRule {
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancePool) DeepCopyInto(out *InstancePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.BootstrapScripts != nil {
		in, out := &in.BootstrapScripts, &out.BootstrapScripts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]*Subnet, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Subnet)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.Firewalls != nil {
		in, out := &in.Firewalls, &out.Firewalls
		*out = make([]*Firewall, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Firewall)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		if *in == nil {
			*out = nil
		} else {
			*out = new(InstancePoolKubernetes)
			**out = **in
		}
	}
	if in.AllowCIDRs != nil {
		in, out := &in.AllowCIDRs, &out.AllowCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]*Label, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Label)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]*Taint, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Taint)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.Amazon != nil {
		in, out := &in.Amazon, &out.Amazon
		if *in == nil {
			*out = nil
		} else {
			*out = new(InstancePoolAmazon)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancePool.
func (in *InstancePool) DeepCopy() *InstancePool {
	if in == nil {
		return nil
	}
	out := new(InstancePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstancePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancePoolAmazon) DeepCopyInto(out *InstancePoolAmazon) {
	*out = *in
	if in.AdditionalIAMPolicies != nil {
		in, out := &in.AdditionalIAMPolicies, &out.AdditionalIAMPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancePoolAmazon.
func (in *InstancePoolAmazon) DeepCopy() *InstancePoolAmazon {
	if in == nil {
		return nil
	}
	out := new(InstancePoolAmazon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancePoolKubernetes) DeepCopyInto(out *InstancePoolKubernetes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancePoolKubernetes.
func (in *InstancePoolKubernetes) DeepCopy() *InstancePoolKubernetes {
	if in == nil {
		return nil
	}
	out := new(InstancePoolKubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGW) DeepCopyInto(out *InternetGW) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGW.
func (in *InternetGW) DeepCopy() *InternetGW {
	if in == nil {
		return nil
	}
	out := new(InternetGW)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesAPI) DeepCopyInto(out *KubernetesAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesAPI.
func (in *KubernetesAPI) DeepCopy() *KubernetesAPI {
	if in == nil {
		return nil
	}
	out := new(KubernetesAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Label) DeepCopyInto(out *Label) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Label.
func (in *Label) DeepCopy() *Label {
	if in == nil {
		return nil
	}
	out := new(Label)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingSink) DeepCopyInto(out *LoggingSink) {
	*out = *in
	if in.ElasticSearch != nil {
		in, out := &in.ElasticSearch, &out.ElasticSearch
		if *in == nil {
			*out = nil
		} else {
			*out = new(LoggingSinkElasticSearch)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]LoggingSinkType, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingSink.
func (in *LoggingSink) DeepCopy() *LoggingSink {
	if in == nil {
		return nil
	}
	out := new(LoggingSink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingSinkElasticSearch) DeepCopyInto(out *LoggingSinkElasticSearch) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.HTTPBasicAuth != nil {
		in, out := &in.HTTPBasicAuth, &out.HTTPBasicAuth
		if *in == nil {
			*out = nil
		} else {
			*out = new(HTTPBasicAuth)
			**out = **in
		}
	}
	if in.AmazonESProxy != nil {
		in, out := &in.AmazonESProxy, &out.AmazonESProxy
		if *in == nil {
			*out = nil
		} else {
			*out = new(AmazonESProxy)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingSinkElasticSearch.
func (in *LoggingSinkElasticSearch) DeepCopy() *LoggingSinkElasticSearch {
	if in == nil {
		return nil
	}
	out := new(LoggingSinkElasticSearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.InternetGW != nil {
		in, out := &in.InternetGW, &out.InternetGW
		if *in == nil {
			*out = nil
		} else {
			*out = new(InternetGW)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSH) DeepCopyInto(out *SSH) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.PublicKeyData != nil {
		in, out := &in.PublicKeyData, &out.PublicKeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSH.
func (in *SSH) DeepCopy() *SSH {
	if in == nil {
		return nil
	}
	out := new(SSH)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Shared) DeepCopyInto(out *Shared) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Shared.
func (in *Shared) DeepCopy() *Shared {
	if in == nil {
		return nil
	}
	out := new(Shared)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Taint) DeepCopyInto(out *Taint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Taint.
func (in *Taint) DeepCopy() *Taint {
	if in == nil {
		return nil
	}
	out := new(Taint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Values) DeepCopyInto(out *Values) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.ItemMap != nil {
		in, out := &in.ItemMap, &out.ItemMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Values.
func (in *Values) DeepCopy() *Values {
	if in == nil {
		return nil
	}
	out := new(Values)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		if *in == nil {
			*out = nil
		} else {
			*out = new(resource.Quantity)
			**out = (*in).DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Volume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
