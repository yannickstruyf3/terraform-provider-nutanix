// Test case for issue 28:
// - https://github.com/terraform-providers/terraform-provider-nutanix/issues/28
//   Terraform crashes when trying to specify boot_device_order_list for a vm resource
// Explanation:
//   Applying the code changes in PR #53 will prevent the provider to crash when passing boot_device_order_list.
//   However, we should revisit the boot_device_order_list solution. From the UI it is not possible to change to order of the devices.
//   It is only possible to set a specific device (iscsi.0, iscsi.1,...).
//   Example json:
//      Default values:
//          "boot_config": {
//              "boot_device_order_list": [
//                  "CDROM",
//                  "DISK",
//                  "NETWORK"
//              ],
//              "boot_type": "LEGACY"
//          },
//      Boot from scsi.0:
//          "boot_config": {
//              "boot_device": {
//                  "disk_address": {
//                      "device_index": 0,
//                      "adapter_type": "SCSI"
//                  }
//              },
//              "boot_type": "LEGACY"
//          },
//      Boot UEFI:
//          "boot_config": {
//              "boot_device": {
//                  "disk_address": {
//                      "device_index": 0,
//                      "adapter_type": "SCSI"
//                  }
//              },
//              "boot_type": "UEFI"
//          },
locals {
  vmname_prefix = "yst-terraform-cat-debugging"
  amount        = 1
  image_name    = "CentOS_7_Cloud"
  cluster_name  = "ntnx-belux-dr"
  subnet_id     = "bcfbc635-ff07-4234-a40b-a3bab85b2e9f"
}

data "nutanix_image" "centos" {
  image_name = local.image_name
}

data "nutanix_cluster" "cluster_by_name" {
  name = local.cluster_name
}

resource "nutanix_virtual_machine" "vm" {
  name                 = "${local.vmname_prefix}-1"
  cluster_uuid         = data.nutanix_cluster.cluster_by_name.id
  num_vcpus_per_socket = "1"
  num_sockets          = "2"
  memory_size_mib      = 2048

  #Will not work? Need confirmation
  boot_device_order_list = ["DISK", "CDROM"]

  nic_list {
    subnet_uuid = "${local.subnet_id}"
  }

  disk_list {
    data_source_reference = {
      kind = "image"
      uuid = data.nutanix_image.centos.id
    }
  }
}
