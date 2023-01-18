# Book Management system 
The main idea is to implement a crud application following the CQRS design pattern<br> 
### Components 
*Query or Book-getter* : is a gin microservice that retrieve data from the database : get one or all books <br>
*Command or book-setter* : is a gin the microservice that modify the database: insert, update & delete
<br>
*Database* : a mongodb instance
<br>

## Existing Architecture (WIP)
<p align="center">
  <img src="./docs/existing-arch.png" alt="Project architecture">
</p>

## Final Aim 
<p align="center">
  <img src="./docs/architecture.png" alt="Project architecture">
</p>


## Observability

### Logging & Traces
I used datadog to persist and visualise the application logs, the attribute added to each log are request_id and ip_adress(not working for now)
<p align="center">
  <img src="./docs/logs-attr.png" alt="Project architecture">
</p>
The logs dashboard :
<p align="center">
  <img src="./docs/logs-dashboard.png" alt="Project architecture">
</p>
A trace example
<p align="center">
  <img src="./docs/trace.png" alt="Project architecture">
</p>

The traces dashboard :
<p align="center">
  <img src="./docs/traces-dashboard.png" alt="Project architecture">
</p>

<p align="center">
  <img src="./docs/kube.png" alt="Project architecture">
</p>




