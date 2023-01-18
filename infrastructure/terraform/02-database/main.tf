data "terraform_remote_state" "aks" {
  backend = "azurerm"
  config = {
    resource_group_name = "terra-book-ms"
    container_name = "bmstfstate"
    storage_account_name = "bmsbackend"
    key = "dev.terraform.bmstfstate"
  }
}


resource "kubernetes_persistent_volume" "mongo-pv" {
    metadata {
      name = "mongo-pv"
      labels = {
        type = "local"
        app = "mongodb-pv"
      }
    }

    spec {
      storage_class_name = "manual"
      capacity = {
        storage = "1Gi"
      }
      access_modes = [ "ReadWriteOnce" ]
      persistent_volume_source {
        host_path {
          path = "/mnt/data"
        }
      }
    }
}

resource "kubernetes_persistent_volume_claim_v1" "mongo-pvc" {
  metadata {
    name = "mongo-pvc"
    labels = {
      app = "mongodb"
    }
  }
  spec {
    storage_class_name = "manual"
    access_modes = [ "ReadWriteOnce" ]
    resources {
      requests = {
        storage = "1Gi"
      }
    }
  }
}

resource "kubernetes_deployment" "db_deploy" {
  metadata {
    name = "mongodb"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app" = "database"
      }
    }
    template {
      metadata {
        labels = {
          app = "database"
        }
      }
      spec {
        container {
          name = "mongodb"
          image = "mongo:4.4"
          image_pull_policy = "IfNotPresent"
          volume_mount {
            name = "mongodb-data"
            mount_path = "/data/db"
          }
        }
        volume {
          name = "mongodb-data"
          persistent_volume_claim {
            claim_name = "mongo-pvc"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "db_service" {
  metadata {
    name = "mongodb"
    labels = {
      app = "database"
    }
  }
  spec {
    type = "NodePort"
    port {
      port = 27017
    }
    selector = {
      app = 30005
    }
  }
}