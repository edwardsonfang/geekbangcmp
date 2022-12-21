# Update envoy.yaml
> fix the fellow bug:
> create deployment with the original envoy.yaml, will report:
> Didn‘t find a registered implementation for ‘envoy.filters.http.router‘ with type URL: ‘‘

    root@cncamp:~/101/module4# cat envoy.yaml
    admin:
      address:
        socket_address: { address: 127.0.0.1, port_value: 9901 }
    
    static_resources:
      listeners:
        - name: listener_0
          address:
            socket_address: { address: 0.0.0.0, port_value: 10000 }
          filter_chains:
            - filters:
                - name: envoy.filters.network.http_connection_manager
                  typed_config:
                    "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
                    stat_prefix: ingress_http
                    codec_type: AUTO
                    route_config:
                      name: local_route
                      virtual_hosts:
                        - name: local_service
                          domains: ["*"]
                          routes:
                            - match: { prefix: "/" }
                              route: { cluster: some_service }
                    http_filters:
                      - name: envoy.filters.http.router
                        typed_config:
                                "@type": type.googleapis.com/envoy.extensions.filters.http.router.v3.Router
      clusters:
        - name: some_service
          connect_timeout: 0.25s
          type: LOGICAL_DNS
          lb_policy: ROUND_ROBIN
          load_assignment:
            cluster_name: some_service
            endpoints:
              - lb_endpoints:
                  - endpoint:
                      address:
                        socket_address:
                          address: nginx
                          port_value: 80

# create envoy configmap
    root@cncamp:~/101/module4# kubectl create configmap envoy-config --from-file=envoy.yaml
    configmap/envoy-config created
    root@cncamp:~/101/module4# kubectl get configmap
    NAME               DATA   AGE
    envoy-config       1      7s
    game-config        1      2d
    game-env-config    7      2d
    kube-root-ca.crt   1      13d
    special-config     2      2d

# create envoy deployment
    root@cncamp:~/101/module4# kubectl create -f envoy-deploy.yaml 
    deployment.apps/envoy created
    root@cncamp:~/101/module4# kubectl get pod
    NAME                      READY   STATUS      RESTARTS      AGE
    centos-7844d8ffc8-r78dv   0/1     Running     3 (43m ago)   4d2h
    configmap-volume-pod      0/1     Completed   0             2d
    envoy-56c8c785d6-xvv7d    1/1     Running     0             5s
    nginx-kube                1/1     Running     3 (43m ago)   4d4h

# expose envoy service with port 10000 and test with curl
    root@cncamp:~/101/module4# kubectl expose deploy envoy --selector run=envoy --port=10000 --type=NodePort
    service/envoy exposed
    root@cncamp:~/101/module4# kubectl get service
    NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)           AGE
    envoy        NodePort    10.110.198.95   <none>        10000:32284/TCP   7s
    kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP           13d
    nginx-kube   ClusterIP   10.110.225.72   <none>        80/TCP            2d
    root@cncamp:~/101/module4# curl 10.110.198.95:10000
    no healthy upstream

# enter envoy pod. check the process and its configuration
    root@cncamp:~/101/module4# kubectl exec -it envoy-56c8c785d6-xvv7d bash
    kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl exec [POD] -- [COMMAND] instead.
    root@envoy-56c8c785d6-xvv7d:/# ls   
    bin   dev                   etc   lib    lib64   media  opt   root  sbin  sys  usr
    boot  docker-entrypoint.sh  home  lib32  libx32  mnt    proc  run   srv   tmp  var
    root@envoy-56c8c785d6-xvv7d:/# pwd
    /
    root@envoy-56c8c785d6-xvv7d:/# ps -ef
    UID          PID    PPID  C STIME TTY          TIME CMD
    envoy          1       0  0 03:42 ?        00:00:00 envoy -c /etc/envoy/envoy.yaml
    root          20       0  0 03:44 pts/0    00:00:00 bash
    root          30      20  0 03:45 pts/0    00:00:00 ps -ef
    root@envoy-56c8c785d6-xvv7d:/# cat /etc/envoy/envoy.yaml 
    admin:
      address:
        socket_address: { address: 127.0.0.1, port_value: 9901 }
    
    static_resources:
      listeners:
        - name: listener_0
          address:
            socket_address: { address: 0.0.0.0, port_value: 10000 }
          filter_chains:
            - filters:
                - name: envoy.filters.network.http_connection_manager
                  typed_config:
                    "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
                    stat_prefix: ingress_http
                    codec_type: AUTO
                    route_config:
                      name: local_route
                      virtual_hosts:
                        - name: local_service
                          domains: ["*"]
                          routes:
                            - match: { prefix: "/" }
                              route: { cluster: some_service }
                    http_filters:
                      - name: envoy.filters.http.router
                        typed_config:
                                "@type": type.googleapis.com/envoy.extensions.filters.http.router.v3.Router
      clusters:
        - name: some_service
          connect_timeout: 0.25s
          type: LOGICAL_DNS
          lb_policy: ROUND_ROBIN
          load_assignment:
            cluster_name: some_service
            endpoints:
              - lb_endpoints:
                  - endpoint:
                      address:
                        socket_address:
                          address: nginx
                          port_value: 80