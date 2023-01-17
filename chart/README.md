# Helm Chart

Here you will find the helm chart for the application.

## How to install

To install the helm chart, first add it to the helm repository

```console
$ helm repo add mychart https://github.com/chaimakr/book_management_system
```

Then update the local repository
```console
$ helm repo update
```

To install the helm chart use the following command
```console
$ helm install <RELEASE_NAME> mychart/book-management-system
```