// Test case for issue 83 and 112:
// - https://github.com/terraform-providers/terraform-provider-nutanix/issues/83
//   Failed to perform NGT ISO mount operation with error code: kNoFreeCdRomSlot
// - https://github.com/terraform-providers/terraform-provider-nutanix/issues/112
//   Improve error handling on incorrect API calls
// Steps:
//  - Use this TF file to create resources
//  - Go to Prism Central and search for your VM
//  - Click `More` and `Install NGT`
//  - Press `Confirm & Enter Password` -> `Skip and Mount` -> `Done`
//  - When looking at the `Recent Tasks` at the right upper corner of PC, You will see failures
//    Note: This fails because the VM doesn't have a free CDrom device
//  - Re-running the TF file is not possible due to errors
// Workaround:
//  - Retrieve the UUID of the virtual machine (should be in your statefile)
//  - Perform a GET request on https://<< Prism IP >>:9440/api/nutanix/v3/vms/<< VM UUID >>
//  - Take the output, remove the “status” key from the JSON
//  - Perform a PUT request using the retrieved payload on https://<< Prism IP >>:9440/api/nutanix/v3/vms/<< VM UUID >>

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
  name                  = "${local.vmname_prefix}-1"
  cluster_uuid          = data.nutanix_cluster.cluster_by_name.id
  num_vcpus_per_socket  = "1"
  num_sockets           = "2"
  memory_size_mib       = 2048
  power_state_mechanism = "ACPI"

  nic_list {
    subnet_uuid = local.subnet_id
  }

  disk_list {
    data_source_reference = {
      kind = "image"
      uuid = data.nutanix_image.centos.id
    }
  }
}

