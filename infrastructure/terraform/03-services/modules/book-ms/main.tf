resource "helm_release" "mybookms" {
  name = var.release_name
  chart = var.chart
  repository = var.repository
}