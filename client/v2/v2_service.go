package v2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/terraform-providers/terraform-provider-nutanix/client"
)

// Operations ...
type Operations struct {
	client *client.Client
}

// Service ...
type Service interface {
	GetVM(uuid string) (*VMResponse, error)
}

/*GetVM Gets a VM
 * This operation gets a op.
 *
 * @param uuid The uuid of the entity.
 * @return *VMIntentResponse
 */
func (op Operations) GetVM(uuid string) (*VMResponse, error) {
	ctx := context.TODO()

	path := fmt.Sprintf("/vms/%s?include_vm_disk_config=true", uuid)

	req, err := op.client.NewRequest(ctx, http.MethodGet, path, nil)
	vmIntentResponse := new(VMResponse)

	if err != nil {
		return nil, err
	}

	return vmIntentResponse, op.client.Do(ctx, req, vmIntentResponse)
}
