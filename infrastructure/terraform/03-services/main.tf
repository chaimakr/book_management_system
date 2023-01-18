data "terraform_remote_state" "aks" {
  backend = "azurerm"
  config = {
    resource_group_name = "terra-book-mb"
    container_name = "bmsbackend"
    storage_account_name = "bmstfstate"
    key = "dev.terraform.bmstfstate"
   }
}

locals {
  kube_config = one(data.terraform_remote_state.aks.outputs.kube_config)
  host                   = local.kube_config.host
  username               = local.kube_config.username
  password               = local.kube_config.password
  client_certificate     = base64decode(local.kube_config.client_certificate)
  client_key             = base64decode(local.kube_config.client_key)
  cluster_ca_certificate = base64decode(local.kube_config.cluster_ca_certificate)
}

provider "kubernetes" {
  alias = "bms-cluster"
  host                   = local.host
  username               = local.username
  password               = local.password
  client_certificate     = local.client_certificate
  client_key             = local.client_key
  cluster_ca_certificate = local.cluster_ca_certificate
}

provider "helm" {
  kubernetes {
    host                   = local.host
    username               = local.username
    password               = local.password
    client_certificate     = local.client_certificate
    client_key             = local.client_key
    cluster_ca_certificate = local.cluster_ca_certificate
  }
}

module "book-ms" {
  source = "./modules/book-ms"
  providers = {
    helm = helm
  }
  chart = var.application_helm_chart
  repository = var.application_helm_repo
  release_name = var.application_helm_release_name
}

module "ingress_controller" {
  source = "./modules/ingress-controller"
  providers = {
    helm = helm
  }
}