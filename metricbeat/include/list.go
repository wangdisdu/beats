/*
Package include imports all Module and MetricSet packages so that they register
their factories with the global registry. This package can be imported in the
main package to automatically register all of the standard supported Metricbeat
modules.
*/
package include

import (
	// This list is automatically generated by `make imports`
	_ "github.com/elastic/beats/metricbeat/module/aerospike"
	_ "github.com/elastic/beats/metricbeat/module/aerospike/namespace"
	_ "github.com/elastic/beats/metricbeat/module/apache"
	_ "github.com/elastic/beats/metricbeat/module/apache/status"
	_ "github.com/elastic/beats/metricbeat/module/ceph"
	_ "github.com/elastic/beats/metricbeat/module/ceph/cluster_disk"
	_ "github.com/elastic/beats/metricbeat/module/ceph/cluster_health"
	_ "github.com/elastic/beats/metricbeat/module/ceph/cluster_status"
	_ "github.com/elastic/beats/metricbeat/module/ceph/monitor_health"
	_ "github.com/elastic/beats/metricbeat/module/ceph/osd_df"
	_ "github.com/elastic/beats/metricbeat/module/ceph/osd_tree"
	_ "github.com/elastic/beats/metricbeat/module/ceph/pool_disk"
	_ "github.com/elastic/beats/metricbeat/module/couchbase"
	_ "github.com/elastic/beats/metricbeat/module/couchbase/bucket"
	_ "github.com/elastic/beats/metricbeat/module/couchbase/cluster"
	_ "github.com/elastic/beats/metricbeat/module/couchbase/node"
	_ "github.com/elastic/beats/metricbeat/module/docker"
	_ "github.com/elastic/beats/metricbeat/module/docker/container"
	_ "github.com/elastic/beats/metricbeat/module/docker/cpu"
	_ "github.com/elastic/beats/metricbeat/module/docker/diskio"
	_ "github.com/elastic/beats/metricbeat/module/docker/healthcheck"
	_ "github.com/elastic/beats/metricbeat/module/docker/image"
	_ "github.com/elastic/beats/metricbeat/module/docker/info"
	_ "github.com/elastic/beats/metricbeat/module/docker/memory"
	_ "github.com/elastic/beats/metricbeat/module/docker/network"
	_ "github.com/elastic/beats/metricbeat/module/dropwizard"
	_ "github.com/elastic/beats/metricbeat/module/dropwizard/collector"
	_ "github.com/elastic/beats/metricbeat/module/elasticsearch"
	_ "github.com/elastic/beats/metricbeat/module/elasticsearch/cluster_stats"
	_ "github.com/elastic/beats/metricbeat/module/elasticsearch/node"
	_ "github.com/elastic/beats/metricbeat/module/elasticsearch/node_stats"
	_ "github.com/elastic/beats/metricbeat/module/etcd"
	_ "github.com/elastic/beats/metricbeat/module/etcd/leader"
	_ "github.com/elastic/beats/metricbeat/module/etcd/self"
	_ "github.com/elastic/beats/metricbeat/module/etcd/store"
	_ "github.com/elastic/beats/metricbeat/module/golang"
	_ "github.com/elastic/beats/metricbeat/module/golang/expvar"
	_ "github.com/elastic/beats/metricbeat/module/golang/heap"
	_ "github.com/elastic/beats/metricbeat/module/graphite"
	_ "github.com/elastic/beats/metricbeat/module/graphite/server"
	_ "github.com/elastic/beats/metricbeat/module/haproxy"
	_ "github.com/elastic/beats/metricbeat/module/haproxy/info"
	_ "github.com/elastic/beats/metricbeat/module/haproxy/stat"
	_ "github.com/elastic/beats/metricbeat/module/http"
	_ "github.com/elastic/beats/metricbeat/module/http/json"
	_ "github.com/elastic/beats/metricbeat/module/http/server"
	_ "github.com/elastic/beats/metricbeat/module/jolokia"
	_ "github.com/elastic/beats/metricbeat/module/jolokia/jmx"
	_ "github.com/elastic/beats/metricbeat/module/kafka"
	_ "github.com/elastic/beats/metricbeat/module/kafka/consumergroup"
	_ "github.com/elastic/beats/metricbeat/module/kafka/partition"
	_ "github.com/elastic/beats/metricbeat/module/kibana"
	_ "github.com/elastic/beats/metricbeat/module/kibana/status"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/container"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/event"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/node"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/pod"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/state_container"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/state_deployment"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/state_node"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/state_pod"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/state_replicaset"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/system"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/util"
	_ "github.com/elastic/beats/metricbeat/module/kubernetes/volume"
	_ "github.com/elastic/beats/metricbeat/module/logstash"
	_ "github.com/elastic/beats/metricbeat/module/logstash/node"
	_ "github.com/elastic/beats/metricbeat/module/logstash/node_stats"
	_ "github.com/elastic/beats/metricbeat/module/memcached"
	_ "github.com/elastic/beats/metricbeat/module/memcached/stats"
	_ "github.com/elastic/beats/metricbeat/module/mongodb"
	_ "github.com/elastic/beats/metricbeat/module/mongodb/dbstats"
	_ "github.com/elastic/beats/metricbeat/module/mongodb/status"
	_ "github.com/elastic/beats/metricbeat/module/mysql"
	_ "github.com/elastic/beats/metricbeat/module/mysql/status"
	_ "github.com/elastic/beats/metricbeat/module/nginx"
	_ "github.com/elastic/beats/metricbeat/module/nginx/stubstatus"
	_ "github.com/elastic/beats/metricbeat/module/php_fpm"
	_ "github.com/elastic/beats/metricbeat/module/php_fpm/pool"
	_ "github.com/elastic/beats/metricbeat/module/postgresql"
	_ "github.com/elastic/beats/metricbeat/module/postgresql/activity"
	_ "github.com/elastic/beats/metricbeat/module/postgresql/bgwriter"
	_ "github.com/elastic/beats/metricbeat/module/postgresql/database"
	_ "github.com/elastic/beats/metricbeat/module/prometheus"
	_ "github.com/elastic/beats/metricbeat/module/prometheus/collector"
	_ "github.com/elastic/beats/metricbeat/module/prometheus/stats"
	_ "github.com/elastic/beats/metricbeat/module/rabbitmq"
	_ "github.com/elastic/beats/metricbeat/module/rabbitmq/node"
	_ "github.com/elastic/beats/metricbeat/module/rabbitmq/queue"
	_ "github.com/elastic/beats/metricbeat/module/redis"
	_ "github.com/elastic/beats/metricbeat/module/redis/info"
	_ "github.com/elastic/beats/metricbeat/module/redis/keyspace"
	_ "github.com/elastic/beats/metricbeat/module/system"
	_ "github.com/elastic/beats/metricbeat/module/system/core"
	_ "github.com/elastic/beats/metricbeat/module/system/cpu"
	_ "github.com/elastic/beats/metricbeat/module/system/diskio"
	_ "github.com/elastic/beats/metricbeat/module/system/filesystem"
	_ "github.com/elastic/beats/metricbeat/module/system/fsstat"
	_ "github.com/elastic/beats/metricbeat/module/system/load"
	_ "github.com/elastic/beats/metricbeat/module/system/memory"
	_ "github.com/elastic/beats/metricbeat/module/system/network"
	_ "github.com/elastic/beats/metricbeat/module/system/process"
	_ "github.com/elastic/beats/metricbeat/module/system/process_summary"
	_ "github.com/elastic/beats/metricbeat/module/system/raid"
	_ "github.com/elastic/beats/metricbeat/module/system/socket"
	_ "github.com/elastic/beats/metricbeat/module/system/uptime"
	_ "github.com/elastic/beats/metricbeat/module/vsphere"
	_ "github.com/elastic/beats/metricbeat/module/vsphere/datastore"
	_ "github.com/elastic/beats/metricbeat/module/vsphere/host"
	_ "github.com/elastic/beats/metricbeat/module/vsphere/virtualmachine"
	_ "github.com/elastic/beats/metricbeat/module/windows"
	_ "github.com/elastic/beats/metricbeat/module/windows/perfmon"
	_ "github.com/elastic/beats/metricbeat/module/windows/service"
	_ "github.com/elastic/beats/metricbeat/module/zookeeper"
	_ "github.com/elastic/beats/metricbeat/module/zookeeper/mntr"
)
