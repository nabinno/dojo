---
title: "Custom Providers with Terraform"
tags: google-cloud-platform, terraform, go
url: https://www.qwiklabs.com/focuses/1204
---

# Goal
- Custom Providers with Terraform

# Task
- [x] Requirement of Custom Providers
- [x] Setup
- [x] Install Terraform
- [x] The provider schema
- [x] Building the Plugin
- [x] Defining Resources
- [x] Invoking the Provider
- [x] Implement Create
- [x] Error handling & partial state
- [x] Implementing Destroy
- [x] Implementing Read

# Supplement
## Install Terraform
```sh
terraform version
wget https://github.com/warrensbox/terraform-switcher/releases/download/0.7.737/terraform-switcher_0.7.737_linux_amd64.tar.gz
mkdir -p ${HOME}/bin
tar -xvf terraform-switcher_0.7.737_linux_amd64.tar.gz -C ${HOME}/bin
export PATH=$PATH:${HOME}/bin
tfswitch -b ${HOME}/bin/terraform 0.11.14
echo "0.11.14" >> .tfswitchrc
exec -l $SHELL
terraform version
```

## The provider schema
```sh
cat <<EOF >provider.go
package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
        return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{},
        }
}
EOF
```

## Building the Plugin
```sh
cat <<EOF >main.go
package main

import (
        "github.com/hashicorp/terraform/plugin"
        "github.com/hashicorp/terraform/terraform"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return Provider()
                },
        })
}
EOF

go get github.com/hashicorp/terraform/helper/schema
go get github.com/hashicorp/terraform/plugin
go get github.com/hashicorp/terraform/terraform
go build -o terraform-provider-example
ls
./terraform-provider-example
```

## Defining Resources
```sh
cat <<EOF >resource_server.go
package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func resourceServer() *schema.Resource {
        return &schema.Resource{
                Create: resourceServerCreate,
                Read:   resourceServerRead,
                Update: resourceServerUpdate,
                Delete: resourceServerDelete,

                Schema: map[string]*schema.Schema{
                        "address": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}
EOF

cat <<EOF >>resource_server.go
func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}
EOF

cat <<EOF >provider.go
package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
        return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
                        "example_server": resourceServer(),
                },
        }
}
EOF

go build -o terraform-provider-example
./terraform-provider-example
```

## Invoking the Provider
```sh
cat <<EOF >main.tf
resource "example_server" "my-server" {}
EOF

terraform init
terraform plan

cat <<EOF >main.tf
resource "example_server" "my-server" {
  address = "1.2.3.4"
}
EOF

terraform plan
```

## Implement Create
```sh
nano resource_server.go
nano> func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
        address := d.Get("address").(string)
        d.SetId(address)
        return nil
}

go build -o terraform-provider-example
terraform init
terraform plan
terraform apply

terraform plan

terraform plan
terraform apply
```

## Error handling & partial state
```sh
nano resource_server.go
nano> func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
        // Enable partial state mode
        d.Partial(true)

        if d.HasChange("address") {
                // Try updating the address
                if err := updateAddress(d, m); err != nil {
                        return err
                }

                d.SetPartial("address")
        }

        // If we were to return here, before disabling partial mode below,
        // then only the "address" field would be saved.

        // We succeeded, disable partial mode. This causes Terraform to save
        // all fields again.
        d.Partial(false)

        return nil
}
```

## Implementing Destroy
```sh
nano resource_server.go
nano> func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
  // d.SetId("") is automatically called assuming delete returns no errors, but
  // it is added here for explicitness.
        d.SetId("")
        return nil
}

go build -o terraform-provider-example
terraform init
terraform destroy
```

## Implementing Read
```sh
nano resource_server.go
nano> func resourceServerRead(d *schema.ResourceData, m interface{}) error {
  client := m.(*MyClient)

  // Attempt to read from an upstream API
  obj, ok := client.Get(d.Id())

  // If the resource does not exist, inform Terraform. We want to immediately
  // return here to prevent further processing.
  if !ok {
    d.SetId("")
    return nil
  }

  d.Set("address", obj.Address)
  return nil
}
```

## Reference
- https://godoc.org/github.com/hashicorp/terraform/helper/schema
- https://www.terraform.io/docs/extend/how-terraform-works.html#terraform-core
