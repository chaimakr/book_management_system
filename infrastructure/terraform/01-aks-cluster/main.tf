data "azurerm_resource_group" "terra-book-ms" {
  name     = "terra-book-ms"
}


resource "azurerm_kubernetes_cluster" "bms-cluster" {
  name = "terra-bms-aks"
  resource_group_name = data.azurerm_resource_group.terra-book-ms.name
  location = data.azurerm_resource_group.terra-book-ms.location
  http_application_routing_enabled = true
  dns_prefix = "bms"
  network_profile {
    network_plugin = "kubenet"
    network_policy = "calico"
  }
  default_node_pool {
    name = "default"
    node_count = 2
    vm_size = "Standard_B2s"
    enable_auto_scaling = false
  }

  identity {
    type = "SystemAssigned"
  }
}