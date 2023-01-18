provider "azurerm" {
  features {}
}

terraform {
  required_providers {
    helm = {
      source  = "hashicorp/helm"
      version = "~>2.8.0"

    }
  }
   backend "azurerm" {
    resource_group_name = "terra-book-ms"
    storage_account_name = "bmsbackend"
    container_name = "bmstfstate"
    key = "bms.terraform.tfstate"
  }
}