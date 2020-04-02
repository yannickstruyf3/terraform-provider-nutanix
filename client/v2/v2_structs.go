package v2

// VMIntentResponse Response object for intentful operations on a vm
type VMResponse struct {
	APIVersion         *string       `json:"api_version,omitempty"`
	AllowLiveMigrate   *bool         `json:"allow_live_migrate,omitempty"`
	GpusAssigned       *bool         `json:"gpus_assigned,omitempty"`
	HaPriority         *int64        `json:"ha_priority,omitempty"`
	HostUUID           *string       `json:"host_uuid,omitempty"`
	MemoryMb           *int64        `json:"memory_mb,omitempty"`
	Name               *string       `json:"name,omitempty"`
	NumCoresPerCpu     *int64        `json:"num_cores_per_vcpu,omitempty"`
	NumVcpu            *int64        `json:"num_vcpus,omitempty"`
	PowerState         *string       `json:"power_state,omitempty"`
	Timezone           *string       `json:"timezone,omitempty"`
	UUID               *string       `json:"uuid,omitempty"`
	VmDiskInfo         *[]VmDiskInfo `json:"vm_disk_info,omitempty"`
	VmFeatures         *VmFeatures   `json:"vm_features,omitempty"`
	VmLogicalTimestamp *int64        `json:"vm_logical_timestamp,omitempty"`
}

type VmFeatures struct {
	AgentVm    *bool `json:"AGENT_VM,omitempty"`
	VgaConsole *bool `json:"VGA_CONSOLE,omitempty"`
}

type VmDiskInfo struct {
	DiskAddress          *DiskAddress       `json:"disk_address,omitempty"`
	IsCdrom              *bool              `json:"is_cdrom,omitempty"`
	IsEmpty              *bool              `json:"is_empty,omitempty"`
	FlashModeEnabled     *bool              `json:"flash_mode_enabled,omitempty"`
	IsScsciPassthrough   *bool              `json:"is_scsi_passthrough,omitempty"`
	IsHotRemoveEnabled   *bool              `json:"is_hot_remove_enabled,omitempty"`
	IsThinProvisioned    *bool              `json:"is_thin_provisioned,omitempty"`
	Shared               *bool              `json:"shared,omitempty"`
	SourceDiskAddress    *SourceDiskAddress `json:"source_disk_address,omitempty"`
	StorageContainerUUID *string            `json:"storage_container_uuid,omitempty"`
	Size                 *int64             `json:"size,omitempty"`
}

type SourceDiskAddress struct {
	VmDiskUUID   *string `json:"vmdisk_uuid,omitempty"`
	NdfsFilepath *string `json:"ndfs_filepath,omitempty"`
}

type DiskAddress struct {
	DeviceBus    *string `json:"device_bus,omitempty"`
	DeviceIndex  *int    `json:"device_index,omitempty"`
	DiskLabel    *string `json:"disk_label,omitempty"`
	NdfsFilepath *string `json:"ndfs_filepath,omitempty"`
	VmdiskUUID   *string `json:"vmdisk_uuid,omitempty"`
	DeviceUUID   *string `json:"device_uuid,omitempty"`
}
