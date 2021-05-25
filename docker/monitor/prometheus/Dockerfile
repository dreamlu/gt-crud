FROM prom/prometheus
COPY prometheus.yml /etc/prometheus/
COPY alert/memory_over.yml /etc/prometheus/
COPY alert/node_down.yml /etc/prometheus/
COPY alert/disk_over.yml /etc/prometheus/
COPY alert/network_over.yml /etc/prometheus/

EXPOSE 9090

