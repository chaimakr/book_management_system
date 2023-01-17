provider "azurerm" {
  features {}
}
terraform {
  required_version = "~>1.3.4"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>3.31.0"
    }
  }
  backend "azurerm" {
    resource_group_name = "terra-book-ms"
    storage_account_name = "bmsbackend"
    container_name = "bmstfstate"
    key = "dev.terraform.bmstfstate"
  }
}