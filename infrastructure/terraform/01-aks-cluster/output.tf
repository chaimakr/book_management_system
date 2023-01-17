output "kube_config" {
  value = azurerm_kubernetes_cluster.bms-cluster.kube_config
  sensitive = true
}