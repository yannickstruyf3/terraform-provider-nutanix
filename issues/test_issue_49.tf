// Test case for issue 49:
// - https://github.com/terraform-providers/terraform-provider-nutanix/issues/49
//   Terraform Unable to use Provisioners in VM Resource when DHCP is used for IP Address
// Explanation:
//  This TF example will create a Windows virtual machine on a network without IPAM.
//  If the VM is provisioned, the IP of the VM will be used to connect via WinRM and create a folder
locals {
  vmname_prefix = "yst-tf-cat"
  amount        = 1
  image_name    = "windows2k12r2"
  cluster_name  = "ntnx-belux-dr"
  # No ipam
  subnet_id   = "138980ea-2418-45e1-a12b-4944dccf44e4"
  vm_username = "administrator"
}

variable "vm_password" {}

data "nutanix_image" "centos" {
  image_name = local.image_name
}

data "nutanix_cluster" "cluster_by_name" {
  name = local.cluster_name
}

data "template_file" "unattend" {
  template = "${file("unattend.xml")}"
  vars = {
    vmname_prefix = local.vmname_prefix
    index         = 1
    vm_password   = var.vm_password
  }
}

resource "nutanix_virtual_machine" "vm" {
  name                 = "${local.vmname_prefix}-1"
  cluster_uuid         = data.nutanix_cluster.cluster_by_name.id
  num_vcpus_per_socket = "1"
  num_sockets          = "2"
  memory_size_mib      = 2048

  nic_list {
    subnet_uuid = "${local.subnet_id}"
  }

  guest_customization_sysprep = {
    install_type = "PREPARED"
    unattend_xml = base64encode(data.template_file.unattend.rendered)
  }

  disk_list {
    data_source_reference = {
      kind = "image"
      uuid = data.nutanix_image.centos.id
    }
  }

  connection {
    type     = "winrm"
    user     = local.vm_username
    password = var.vm_password
    use_ntlm = true
    insecure = true
    host     = "${nutanix_virtual_machine.vm.nic_list_status.0.ip_endpoint_list.0.ip}"
  }

  provisioner "remote-exec" {
    inline = [
      "powershell.exe mkdir c:/testing"
    ]
  }
}
