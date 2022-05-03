# check_traefik - a monitoring plugin to check Trafik instances

This monitoring plugin checks [Traefik](https://traefik.io/traefik/) by querying the [Ping](https://doc.traefik.io/traefik/operations/ping/) health-check URL.

## Example:
Configure your Traefik service to activate ping, e.g. in docker-compose
```yaml
services:
  frontend:
    image: traefik:v2.6
    command:
    # ...
    # Enable ping and use custom routing
    - --ping=true
    - --ping.manualrouting=true
    labels:
      traefik.enable: "true"
      traefik.http.middlewares.auth-monitoring.basicauth.users: monitoring:$$2y$$05$$9kDJQAJckPliR0Px7Qcxs.LRPpeC4G.cF7F87Fa1NJXW6/9YOKTLa
      traefik.http.routers.ping.entrypoints: https
      traefik.http.routers.ping.middlewares: auth-monitoring
      traefik.http.routers.ping.rule: Host(`traefik.domain.tld`) && PathPrefix(`/ping`)
      traefik.http.routers.ping.service: ping@internal
      traefik.http.routers.ping.tls: "true"
```

```shell
check_traefik health -I 192.0.2.101 -H traefik.domain.tld --user="monitoring" --password="password"
```

## Coming soon:
Query the [API](https://doc.traefik.io/traefik/operations/api/#endpoints).
* Compare the response of `/api/version` against latest GitHub release or a static version.
* Check the response of `/api/overview` for total, warnings and errors.
