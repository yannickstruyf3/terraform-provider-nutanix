package nutanix

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

const (
	PERMISSIONNAME = "Access_Console_Virtual_Machine"
	PERMISSINOUUID = "178e1e72-a46f-4b1d-b12b-c53c2c2db83c"
)

func TestAccNutanixPermissionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPermissionDataSourceConfig(PERMISSINOUUID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "name", PERMISSIONNAME),
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "operation", "console_access"),
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "fields.0.field_mode", "DISALLOWED"),
				),
			},
		},
	})
}

func testAccPermissionDataSourceConfig(uuid string) string {
	return fmt.Sprintf(`
data "nutanix_permission" "test" {
	permission_id = "%s"
}
`, uuid)
}

func TestAccNutanixPermissionDataSource_basicByName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPermissionDataSourceConfigByName(PERMISSIONNAME),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "name", PERMISSIONNAME),
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "operation", "console_access"),
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "fields.0.field_mode", "DISALLOWED"),
				),
			},
		},
	})
}

func testAccPermissionDataSourceConfigByName(name string) string {
	return fmt.Sprintf(`
data "nutanix_permission" "test" {
	permission_name = "%s"
}
`, name)
}
