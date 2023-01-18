data "terraform_remote_state" "aks" {
  backend = "azurerm"
  config = {
    resource_group_name = "terra-book-ms"
    container_name = "bmstfstate"
    storage_account_name = "bmsbackend"
    key = "dev.terraform.tfstate"
  }
}


resource "helm_release" "datadog" {
  name = "datadog"
  chart = "datadog"
  repository = "https://helm.datadoghq.com"

  set {
    name = "datadog.apiKey"
    value = var.apiKey
    type = "string"
  }
  values = [
    "${file("values/datadog-values.yaml")}"
  ]
}

resource "helm_release" "prometheus" {
  name = "prometheus"
  chart = "prometheus"
  repository = "https://prometheus-community.github.io/helm-charts"
  values = [ "${file("values/prometheus-values.yaml")}" ]
}

resource "helm_release" "grafana" {
  name = "grafana"
  chart = "grafana"
  repository = "https://grafana.github.io/helm-charts"
}