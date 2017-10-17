ansible-twemproxy role
======================

Tested on 16.04

# Usage

# Variables

All variables can be found at twemproxy documentation.

- `twemproxy_mbuf_size`: 16384
- `twemproxy_output`: /var/log/twemproxy.log
- `twemproxy_stats_port`: 22221
- `twemproxy_stats_addr`: 127.0.0.1
- `twemproxy_stats_interval`: 30000 # msecs

# Pools

Pools are defined the same way it is in the twemproxy yaml structure, except for the server part.

The server key pust contain an ansible group, a port and a network interface.
Servers will be created from this information: all servers for the group will
be added in the pool with the IP address corresponding to the designated
interface, and on the configured port.

Weight can not be currently changed.

Example:

```yaml
twemproxy_pools:
  redis:
    listen: 127.0.0.1:6379
    hash: fnv1a_64
    distribution: ketama
    auto_eject_hosts: "true"
    redis: "true"
    server_retry_timeout: 30000
    server_failure_limit: 1
    servers:
      group: redis
      port: "{{ redis_port }}"
      interface: "{{ redis_interface }}"
```



Michel Blanc <mb@mbnet.fr>
