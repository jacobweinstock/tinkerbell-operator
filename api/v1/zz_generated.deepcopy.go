// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Create) DeepCopyInto(out *Create) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Create.
func (in *Create) DeepCopy() *Create {
	if in == nil {
		return nil
	}
	out := new(Create)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Custom) DeepCopyInto(out *Custom) {
	*out = *in
	out.PreinstalledOperatingSystemVersion = in.PreinstalledOperatingSystemVersion
	if in.PrivateSubnets != nil {
		in, out := &in.PrivateSubnets, &out.PrivateSubnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Custom.
func (in *Custom) DeepCopy() *Custom {
	if in == nil {
		return nil
	}
	out := new(Custom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Data) DeepCopyInto(out *Data) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Network.DeepCopyInto(&out.Network)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Data.
func (in *Data) DeepCopy() *Data {
	if in == nil {
		return nil
	}
	out := new(Data)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dhcp) DeepCopyInto(out *Dhcp) {
	*out = *in
	out.IP = in.IP
	if in.NameServers != nil {
		in, out := &in.NameServers, &out.NameServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeServers != nil {
		in, out := &in.TimeServers, &out.TimeServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dhcp.
func (in *Dhcp) DeepCopy() *Dhcp {
	if in == nil {
		return nil
	}
	out := new(Dhcp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disks) DeepCopyInto(out *Disks) {
	*out = *in
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = make([]Partitions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disks.
func (in *Disks) DeepCopy() *Disks {
	if in == nil {
		return nil
	}
	out := new(Disks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Facility) DeepCopyInto(out *Facility) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Facility.
func (in *Facility) DeepCopy() *Facility {
	if in == nil {
		return nil
	}
	out := new(Facility)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filesystems) DeepCopyInto(out *Filesystems) {
	*out = *in
	in.Mount.DeepCopyInto(&out.Mount)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filesystems.
func (in *Filesystems) DeepCopy() *Filesystems {
	if in == nil {
		return nil
	}
	out := new(Filesystems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hardware) DeepCopyInto(out *Hardware) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hardware.
func (in *Hardware) DeepCopy() *Hardware {
	if in == nil {
		return nil
	}
	out := new(Hardware)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hardware) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareList) DeepCopyInto(out *HardwareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hardware, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareList.
func (in *HardwareList) DeepCopy() *HardwareList {
	if in == nil {
		return nil
	}
	out := new(HardwareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HardwareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareSpec) DeepCopyInto(out *HardwareSpec) {
	*out = *in
	in.Data.DeepCopyInto(&out.Data)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareSpec.
func (in *HardwareSpec) DeepCopy() *HardwareSpec {
	if in == nil {
		return nil
	}
	out := new(HardwareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareStatus) DeepCopyInto(out *HardwareStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareStatus.
func (in *HardwareStatus) DeepCopy() *HardwareStatus {
	if in == nil {
		return nil
	}
	out := new(HardwareStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IP) DeepCopyInto(out *IP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IP.
func (in *IP) DeepCopy() *IP {
	if in == nil {
		return nil
	}
	out := new(IP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.OperatingSystemVersion = in.OperatingSystemVersion
	in.Storage.DeepCopyInto(&out.Storage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interfaces) DeepCopyInto(out *Interfaces) {
	*out = *in
	in.Dhcp.DeepCopyInto(&out.Dhcp)
	out.Netboot = in.Netboot
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Interfaces.
func (in *Interfaces) DeepCopy() *Interfaces {
	if in == nil {
		return nil
	}
	out := new(Interfaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ipxe) DeepCopyInto(out *Ipxe) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ipxe.
func (in *Ipxe) DeepCopy() *Ipxe {
	if in == nil {
		return nil
	}
	out := new(Ipxe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Manufacturer) DeepCopyInto(out *Manufacturer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Manufacturer.
func (in *Manufacturer) DeepCopy() *Manufacturer {
	if in == nil {
		return nil
	}
	out := new(Manufacturer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	in.Custom.DeepCopyInto(&out.Custom)
	out.Facility = in.Facility
	in.Instance.DeepCopyInto(&out.Instance)
	out.Manufacturer = in.Manufacturer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mount) DeepCopyInto(out *Mount) {
	*out = *in
	in.Create.DeepCopyInto(&out.Create)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mount.
func (in *Mount) DeepCopy() *Mount {
	if in == nil {
		return nil
	}
	out := new(Mount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Netboot) DeepCopyInto(out *Netboot) {
	*out = *in
	out.Ipxe = in.Ipxe
	out.Osie = in.Osie
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Netboot.
func (in *Netboot) DeepCopy() *Netboot {
	if in == nil {
		return nil
	}
	out := new(Netboot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]Interfaces, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
func (in *OperatingSystemVersion) DeepCopyInto(out *OperatingSystemVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatingSystemVersion.
func (in *OperatingSystemVersion) DeepCopy() *OperatingSystemVersion {
	if in == nil {
		return nil
	}
	out := new(OperatingSystemVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Osie) DeepCopyInto(out *Osie) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Osie.
func (in *Osie) DeepCopy() *Osie {
	if in == nil {
		return nil
	}
	out := new(Osie)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Partitions) DeepCopyInto(out *Partitions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Partitions.
func (in *Partitions) DeepCopy() *Partitions {
	if in == nil {
		return nil
	}
	out := new(Partitions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreinstalledOperatingSystemVersion) DeepCopyInto(out *PreinstalledOperatingSystemVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreinstalledOperatingSystemVersion.
func (in *PreinstalledOperatingSystemVersion) DeepCopy() *PreinstalledOperatingSystemVersion {
	if in == nil {
		return nil
	}
	out := new(PreinstalledOperatingSystemVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]Disks, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Filesystems != nil {
		in, out := &in.Filesystems, &out.Filesystems
		*out = make([]Filesystems, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}