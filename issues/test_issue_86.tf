// Test case for issue 86:
// - https://github.com/terraform-providers/terraform-provider-nutanix/issues/86
//   CPU and RAM change does not restart VM automatically
// Explanation:
//   Current implementation of the TF provider relies on hotadd of resources (example: adding memory without requiring a reboot).
//   Some operating systems do not support adding resources without reboot. The provider should support this functionality.
// Steps:
//  - Create image based on Debian 8.11 (https://cdimage.debian.org/cdimage/archive/8.11.1/amd64/iso-cd/debian-8.11.1-amd64-netinst.iso)
//  - Run this TF file to create a Debian 8.11 VM
//  - Login on the VM and verify if the amount of memory is correct (free -h)
//  - Modify the `memory_size_mb` to a larger value
//  - Re-run TF
//  - Validate if the amount of memory is modifed in Prism (should be the case)
//  - Validate if the amount of memory is modifed in the OS (should NOT be the case)
//  - Reboot virtual machine via Prism and verify if the memory is now correctly reflected in the OS

locals {
  vmname_prefix = "yst-terraform-cat-debugging"
  amount        = 1
  image_name    = "debian_8.11_nohotadd"
  cluster_name  = "ntnx-belux-dr"
  subnet_id     = "bcfbc635-ff07-4234-a40b-a3bab85b2e9f"
  num_sockets   = "2"
  #   memory_size_mib = 2048
  memory_size_mib = 4096
}

data "nutanix_image" "debian" {
  image_name = local.image_name
}

data "nutanix_cluster" "cluster_by_name" {
  name = local.cluster_name
}

resource "nutanix_virtual_machine" "vm" {
  name                 = "${local.vmname_prefix}-1"
  cluster_uuid         = data.nutanix_cluster.cluster_by_name.id
  num_vcpus_per_socket = "1"
  num_sockets          = local.num_sockets
  memory_size_mib      = local.memory_size_mib

  nic_list {
    subnet_uuid = "${local.subnet_id}"
  }

  disk_list {
    data_source_reference = {
      kind = "image"
      uuid = data.nutanix_image.debian.id
    }
  }
}
