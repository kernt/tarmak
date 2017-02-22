class puppernetes::master(
  $disable_kubelet = true,
  $disable_proxy = true,
){

  $apiserver_alt_names='kubernetes.default'
  $apiserver_ip_sans='10.254.0.1'
  include ::puppernetes
  include ::vault_client

  Class['vault_client'] -> Class['puppernetes::master']

  $service_account_key_path = "${::puppernetes::kubernetes_ssl_dir}/service-account-key.pem"
  vault_client::secret_service { 'kube-service-account-key':
    field       => 'key',
    secret_path => "${::puppernetes::cluster_name}/secrets/service-accounts",
    user        => $::puppernetes::kubernetes_user,
    dest_path   => $service_account_key_path,
  }

  $controller_manager_base_path = "${::puppernetes::kubernetes_ssl_dir}/kube-controller-manager"
  vault_client::cert_service { 'kube-controller-manager':
    base_path   => $controller_manager_base_path,
    common_name => 'kube-controller-manager',
    role        => "${::puppernetes::cluster_name}/pki/${::puppernetes::kubernetes_ca_name}/sign/kube-controller-manager",
    user        => $::puppernetes::kubernetes_user,
  }

  $scheduler_base_path = "${::puppernetes::kubernetes_ssl_dir}/kube-scheduler"
  vault_client::cert_service { 'kube-scheduler':
    base_path   => $scheduler_base_path,
    common_name => 'kube-scheduler',
    role        => "${::puppernetes::cluster_name}/pki/${::puppernetes::kubernetes_ca_name}/sign/kube-scheduler",
    user        => $::puppernetes::kubernetes_user,
  }

  $apiserver_base_path = "${::puppernetes::kubernetes_ssl_dir}/kube-apiserver"
  vault_client::cert_service { 'kube-apiserver':
    base_path   => $apiserver_base_path,
    common_name => "kube-apiserver.${::puppernetes::cluster_name}.${::puppernetes::dns_root}",
    role        => "${::puppernetes::cluster_name}/pki/${::puppernetes::kubernetes_ca_name}/sign/kube-apiserver",
    user        => $::puppernetes::kubernetes_user,
    ip_sans     => "${apiserver_ip_sans},${::puppernetes::ipaddress}",
    alt_names   => $apiserver_alt_names,
  }

  $admin_base_path = "${::puppernetes::kubernetes_ssl_dir}/kube-admin"
  vault_client::cert_service { 'kube-admin':
    base_path   => $admin_base_path,
    common_name => 'admin',
    role        => "${::puppernetes::cluster_name}/pki/${::puppernetes::kubernetes_ca_name}/sign/admin",
    user        => $::puppernetes::kubernetes_user,
  }

  $etcd_apiserver_base_path = "${::puppernetes::kubernetes_ssl_dir}/${::puppernetes::etcd_k8s_main_ca_name}"
  vault_client::cert_service { 'etcd-apiserver':
    base_path   => $etcd_apiserver_base_path,
    common_name => "${::hostname}.${::puppernetes::cluster_name}.${::puppernetes::dns_root}",
    role        => "${::puppernetes::cluster_name}/pki/${::puppernetes::etcd_k8s_main_ca_name}/sign/client",
    user        => $::puppernetes::kubernetes_user,
    ip_sans     => $::puppernetes::ipaddress,
  }

  class { 'kubernetes::apiserver':
      ca_file          => "${apiserver_base_path}-ca.pem",
      key_file         => "${apiserver_base_path}-key.pem",
      cert_file        => "${apiserver_base_path}.pem",
      etcd_ca_file     => "${etcd_apiserver_base_path}-ca.pem",
      etcd_key_file    => "${etcd_apiserver_base_path}-key.pem",
      etcd_cert_file   => "${etcd_apiserver_base_path}.pem",
      etcd_port        => $::puppernetes::etcd_k8s_main_client_port,
      etcd_events_port => $::puppernetes::etcd_k8s_events_client_port,
      systemd_after    => ['kube-apiserver-cert.service', 'kube-service-account-key-secret.service'],
      systemd_requires => ['kube-apiserver-cert.service', 'kube-service-account-key-secret.service'],
  }

  class { 'kubernetes::controller_manager':
      ca_file          => "${controller_manager_base_path}-ca.pem",
      key_file         => "${controller_manager_base_path}-key.pem",
      cert_file        => "${controller_manager_base_path}.pem",
      systemd_after    => ['kube-controller-manager-cert.service', 'kube-service-account-key-secret.service'],
      systemd_requires => ['kube-controller-manager-cert.service', 'kube-service-account-key-secret.service'],
  }

  class { 'kubernetes::scheduler':
      ca_file          => "${scheduler_base_path}-ca.pem",
      key_file         => "${scheduler_base_path}-key.pem",
      cert_file        => "${scheduler_base_path}.pem",
      systemd_after    => ['kube-scheduler-cert.service'],
      systemd_requires => ['kube-scheduler-cert.service'],
  }

  class { 'kubernetes::kubectl':
      ca_file   => "${admin_base_path}-ca.pem",
      key_file  => "${admin_base_path}-key.pem",
      cert_file => "${admin_base_path}.pem",
  }

  Service['kube-scheduler-cert.service'] -> Service['kube-scheduler.service']
  Service['kube-controller-manager-cert.service'] -> Service['kube-controller-manager.service']
  Service['kube-apiserver-cert.service'] -> Service['kube-apiserver.service']
  Service['kube-service-account-key-secret.service'] -> Service['kube-controller-manager.service']
  Service['kube-service-account-key-secret.service'] -> Service['kube-apiserver.service']
  Service['kube-admin-cert.service'] -> Kubernetes::Apply <||>

  class { 'kubernetes::master':
    disable_kubelet => $disable_kubelet,
    disable_proxy   => $disable_proxy,
  }

}
