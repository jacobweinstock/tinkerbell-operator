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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	PhasePending = "PENDING"
	PhaseRunning = "RUNNING"
	PhaseDone    = "DONE"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Data struct {
	ID       string   `json:"id"`
	Metadata Metadata `json:"metadata"`
	Network  Network  `json:"network"`
}
type PreinstalledOperatingSystemVersion struct {
}
type Custom struct {
	PreinstalledOperatingSystemVersion PreinstalledOperatingSystemVersion `json:"preinstalled_operating_system_version"`
	PrivateSubnets                     []string                           `json:"private_subnets"`
}
type Facility struct {
	FacilityCode    string `json:"facility_code"`
	PlanSlug        string `json:"plan_slug"`
	PlanVersionSlug string `json:"plan_version_slug"`
}
type OperatingSystemVersion struct {
	Distro  string `json:"distro"`
	OsSlug  string `json:"os_slug"`
	Version string `json:"version"`
}
type Partitions struct {
	Label  string `json:"label"`
	Number int    `json:"number"`
	Size   int    `json:"size"`
}
type Disks struct {
	Device     string       `json:"device"`
	Partitions []Partitions `json:"partitions"`
	WipeTable  bool         `json:"wipe_table"`
}
type Create struct {
	Options []string `json:"options"`
}
type Mount struct {
	Create Create `json:"create"`
	Device string `json:"device"`
	Format string `json:"format"`
	Point  string `json:"point"`
}
type Filesystems struct {
	Mount Mount `json:"mount"`
}
type Storage struct {
	Disks       []Disks       `json:"disks"`
	Filesystems []Filesystems `json:"filesystems"`
}
type Instance struct {
	CryptedRootPassword    string                 `json:"crypted_root_password"`
	OperatingSystemVersion OperatingSystemVersion `json:"operating_system_version"`
	Storage                Storage                `json:"storage"`
}
type Manufacturer struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}
type Metadata struct {
	BondingMode  int          `json:"bonding_mode"`
	Custom       Custom       `json:"custom"`
	Facility     Facility     `json:"facility"`
	Instance     Instance     `json:"instance"`
	Manufacturer Manufacturer `json:"manufacturer"`
	State        string       `json:"state"`
}
type IP struct {
	Address string `json:"address"`
	Gateway string `json:"gateway"`
	Netmask string `json:"netmask"`
}
type Dhcp struct {
	Arch        string   `json:"arch"`
	Hostname    string   `json:"hostname"`
	IP          IP       `json:"ip"`
	LeaseTime   int      `json:"lease_time"`
	Mac         string   `json:"mac"`
	NameServers []string `json:"name_servers"`
	TimeServers []string `json:"time_servers"`
	Uefi        bool     `json:"uefi"`
}
type Ipxe struct {
	Contents string `json:"contents"`
	URL      string `json:"url"`
}
type Osie struct {
	BaseURL string `json:"base_url"`
	Initrd  string `json:"initrd"`
	Kernel  string `json:"kernel"`
}
type Netboot struct {
	AllowPxe      bool `json:"allow_pxe"`
	AllowWorkflow bool `json:"allow_workflow"`
	Ipxe          Ipxe `json:"ipxe"`
	Osie          Osie `json:"osie"`
}
type Interfaces struct {
	Dhcp    Dhcp    `json:"dhcp"`
	Netboot Netboot `json:"netboot"`
}
type Network struct {
	Interfaces []Interfaces `json:"interfaces"`
}

// HardwareSpec defines the desired state of Hardware
type HardwareSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Hardware. Edit Hardware_types.go to remove/update
	Data Data `json:"data"`
}

// HardwareStatus defines the observed state of Hardware
type HardwareStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true

// Hardware is the Schema for the hardwares API
type Hardware struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HardwareSpec   `json:"spec,omitempty"`
	Status HardwareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HardwareList contains a list of Hardware
type HardwareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hardware `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Hardware{}, &HardwareList{})
}
