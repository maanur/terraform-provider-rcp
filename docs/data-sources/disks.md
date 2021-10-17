---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "rustack_disks Data Source - terraform-provider-rustack"
subcategory: ""
description: |-
  
---

# rustack_disks (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **vdc_id** (String) id of the VDC

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **disks** (List of Object) (see [below for nested schema](#nestedatt--disks))

<a id="nestedatt--disks"></a>
### Nested Schema for `disks`

Read-Only:

- **id** (String)
- **name** (String)
- **size** (Number)
- **storage_profile_id** (String)
- **storage_profile_name** (String)

