resource "azurerm_resource_group" "terra-book-ms" {
  name     = "terra-book-ms"
  location = "France Central"
}
resource "azurerm_storage_account" "bms-storage-account" {
  name                     = "bmsbackend"
  resource_group_name      = azurerm_resource_group.terra-book-ms.name
  location                 = azurerm_resource_group.terra-book-ms.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    environment = "staging"
  }
}

resource "azurerm_storage_container" "bms-storage-container" {
  name                  = "bmstfstate"
  storage_account_name  = azurerm_storage_account.bms-storage-account.name
  container_access_type = "private"
}
