require 'spec_helper'

describe 'kubernetes::apiserver' do
  let :service_file do
    '/etc/systemd/system/kube-apiserver.service'
  end

  context 'with default values for all parameters' do
    it { should contain_class('kubernetes::apiserver') }
    it do
      should contain_file(service_file).with_content(/After=network.target/)
      should contain_file(service_file).with_content(/User=kubernetes/)
      should contain_file(service_file).with_content(/Group=kubernetes/)
      should contain_file(service_file).with_content(/#{Regexp.escape('"--etcd-servers=http://localhost:2379"')}/)
      should contain_file(service_file).with_content(%r{--service-cluster-ip-range=10\.254\.0\.0/16})
      should contain_file(service_file).with_content(%r{--allow-privileged=true})
    end
  end

  context 'with etcd override for events' do
    let(:params) { {'etcd_events_port' => 1234 } }
    it 'should have an etcd overrides line' do
      should contain_file(service_file).with_content(/#{Regexp.escape('"--etcd-servers-overrides=/events#http://localhost:1234"')}/)
    end
  end

  context 'insecure bind address' do
    context 'is specified' do
      let(:params) { {'insecure_bind_address' => '127.0.0.1' } }
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--insecure-bind-address=127.0.0.1')}/)}
    end
    context 'not specified' do
      it { should_not contain_file(service_file).with_content(/#{Regexp.escape('--insecure-bind-address=')}/)}
    end
  end

  context 'cloud provider' do
    context 'default' do
      it { should_not contain_file(service_file).with_content(%r{--cloud-provider}) }
    end

    context 'aws' do
      let(:pre_condition) {[
        """
        class{'kubernetes': cloud_provider => 'aws'}
        """
      ]}
      it { should contain_file(service_file).with_content(%r{--cloud-provider=aws}) }
    end
  end

  context 'iam authenticator' do
    context 'default' do
      it do
        should contain_class('kubernetes::aws_iam_authenticator_init').with('file_ensure' => 'absent')
        should_not contain_file(service_file).with_content(%r{aws-iam-authenticator-init.service})
      end
    end

    context 'iam authenticator enabled' do
      let(:pre_condition) {[
        """
        class{'kubernetes::apiserver': aws_iam_authenticator_init => true}
        """
      ]}
      it { should contain_file(service_file).with_content(/#{Regexp.escape('Requires=')}.*#{Regexp.escape('aws-iam-authenticator-init.service')}/) }
      it { should contain_class('kubernetes::aws_iam_authenticator_init').with('file_ensure' => 'file') }
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--authentication-token-webhook-config-file=/etc/kubernetes/aws-iam-authenticator/kubeconfig.yaml')}/) }
    end
  end

  context 'admission controllers' do
    context 'customized' do
      let(:params) { {'admission_control' => ['Test1'] } }
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--admission-control=Test1')}/)}
    end

    context 'default pre 1.4' do
      let(:pre_condition) {[
        """
        class{'kubernetes': version => '1.3.5'}
        """
      ]}
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--admission-control=NamespaceLifecycle,LimitRanger,ServiceAccount,ResourceQuota')}/)}
    end

    context 'default 1.4+' do
      let(:pre_condition) {[
        """
        class{'kubernetes': version => '1.4.0'}
        """
      ]}
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--admission-control=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota')}/)}
    end

    context 'default 1.13' do
      let(:pre_condition) {[
        """
        class{'kubernetes': version => '1.13.0'}
        """
      ]}
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--enable-admission-plugins=Initializers,NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,PodSecurityPolicy,NodeRestriction')}/)}
    end

    context 'default 1.14' do
      let(:pre_condition) {[
        """
        class{'kubernetes': version => '1.14.0'}
        """
      ]}
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,PodSecurityPolicy,NodeRestriction')}/)}
    end
  end

  context 'storage backend' do
    context 'customized' do
      let(:params) { {'storage_backend' => 'consulator' } }
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--storage-backend=consulator')}/)}
    end

    context 'default pre 1.5' do
      let(:pre_condition) {[
        """
        class{'kubernetes': version => '1.4.8'}
        """
      ]}
      it { should_not contain_file(service_file).with_content(/#{Regexp.escape('--storage-backend=etcd2')}/)}
    end

    context 'default 1.5+' do
      let(:pre_condition) {[
        """
        class{'kubernetes': version => '1.5.0'}
        """
      ]}
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--storage-backend=etcd3')}/)}
    end
  end

  context 'auth token webhook file' do
    let(:pre_condition) {[
      """
      class{'kubernetes::apiserver': auth_token_webhook_file => '/foo/bar/baz'}
      """
    ]}
    it { should contain_file(service_file).with_content(/#{Regexp.escape('--authentication-token-webhook-config-file=/foo/bar/baz')}/)}
    let(:pre_condition) {[
      """
      class{'kubernetes::apiserver':
        aws_iam_authenticator_init => true,
        auth_token_webhook_file    => '/foo/bar/baz'}
      """
    ]}
    it { should contain_file(service_file).with_content(/#{Regexp.escape('--authentication-token-webhook-config-file=/foo/bar/baz')}/)}
    it { should contain_class('kubernetes::aws_iam_authenticator_init').with('auth_token_webhook_file' => '/foo/bar/baz')}
  end

  context 'runtime_config' do
    let :kubernetes_version do
      '1.6.2'
    end

    let :authorization_mode do
      '[\'RBAC\']'
    end

    let(:pre_condition) {[
      """
        class{'kubernetes':
          version => '#{kubernetes_version}',
          authorization_mode => #{authorization_mode},
        }
      """
    ]}

    context 'default' do
      it 'should not set a runtime config' do
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--runtime-config=')}/)
      end
    end

    context 'RBAC before 1.6.0' do
      let :kubernetes_version do
        '1.5.7'
      end

      it 'should activate RBAC API via a runtime config' do
        should contain_file(service_file).with_content(/#{Regexp.escape('--runtime-config=')}/)
      end
    end
  end

  context 'request header and proxy client options' do
    let(:params) { {
      'requestheader_client_ca_file' => '/tmp/proxy-ca.pem',
      'proxy_client_cert_file' => '/tmp/proxy.pem',
      'proxy_client_key_file' => '/tmp/proxy-key.pem',
    } }

    let(:pre_condition) {[
      """
        class{'kubernetes':
          version => '#{kubernetes_version}',
        }
      """
    ]}

    context 'all necessary parameters k8s version 1.5.x' do
      let(:kubernetes_version) { '1.5.7' }
      it "should not setup request header options" do
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--requestheader-client-ca-file=/tmp/proxy-ca.pem')}/)
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--proxy-client-cert-file=/tmp/proxy.pem')}/)
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--proxy-client-key-file=/tmp/proxy-key.pem')}/)
      end
    end

    context 'all necessary parameters k8s version 1.6.x' do
      let(:kubernetes_version) { '1.6.12' }
      it "should not setup request header options" do
        should contain_file(service_file).with_content(/#{Regexp.escape('--requestheader-client-ca-file=/tmp/proxy-ca.pem')}/)
        should contain_file(service_file).with_content(/#{Regexp.escape('--proxy-client-cert-file=/tmp/proxy.pem')}/)
        should contain_file(service_file).with_content(/#{Regexp.escape('--proxy-client-key-file=/tmp/proxy-key.pem')}/)
      end
    end

    context 'missing parameters k8s version 1.6.x' do
      let(:params) { {
        'proxy_client_cert_file' => '/tmp/proxy.pem',
        'proxy_client_key_file' => '/tmp/proxy-key.pem',
      } }

      let(:kubernetes_version) { '1.6.12' }
      it "should not setup request header options" do
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--requestheader-client-ca-file=/tmp/proxy-ca.pem')}/)
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--proxy-client-cert-file=/tmp/proxy.pem')}/)
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--proxy-client-key-file=/tmp/proxy-key.pem')}/)
      end
    end
  end

  context 'authorization_mode' do
    let :kubernetes_version do
      '1.6.2'
    end

    let :authorization_mode do
      '[]'
    end

    let(:pre_condition) {[
      """
        class{'kubernetes':
          version => '#{kubernetes_version}',
          authorization_mode => #{authorization_mode},
        }
      """
    ]}

    let :policy_file do
      '/etc/kubernetes/kube-apiserver-abac-policy.json'
    end

    context 'default k8s version 1.6+' do
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--authorization-mode=RBAC')}/)}
      it { should_not contain_file(service_file).with_content(/#{Regexp.escape("--authorization-policy-file=#{policy_file}")}/)}
      it { should_not contain_file(policy_file) }
    end

    context 'default k8s version 1.5.x' do
      let(:kubernetes_version) { '1.5.7' }
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--authorization-mode=ABAC')}/)}
      it { should contain_file(service_file).with_content(/#{Regexp.escape("--authorization-policy-file=#{policy_file}")}/)}
      it 'contains rules for important subjects' do
        should contain_file(policy_file).with_content(/#{Regexp.escape('"system:node"')}/)
        should contain_file(policy_file).with_content(/#{Regexp.escape('"system:kube-controller-manager"')}/)
        should contain_file(policy_file).with_content(/#{Regexp.escape('"system:kube-scheduler"')}/)
        should contain_file(policy_file).with_content(/#{Regexp.escape('"admin"')}/)
        should contain_file(policy_file).with_content(/#{Regexp.escape('generic endpoint')}/)
      end
    end

    context 'default k8s version before 1.5' do
      let(:kubernetes_version) { '1.4.3' }
      it { should_not contain_file(policy_file).with_content(/#{Regexp.escape('generic endpoint')}/) }
    end

    context 'set to AlwaysAllow' do
      let(:authorization_mode) { '[\'AlwaysAllow\']' }
      it { should contain_file(service_file).with_content(/#{Regexp.escape('--authorization-mode=AlwaysAllow')}/)}
    end

    context 'deprecated insecure-port flag after 1.11' do
      context 'insecure-port' do
        context 'should exist before 1.11' do
          let(:pre_condition) {[
          """
          class{'kubernetes': version => '1.10.7'}
          """
          ]}

          it {should contain_file(service_file).with_content(/#{Regexp.escape('--insecure-port=')}/)}
        end

        context 'should exist after 1.11' do
          let(:pre_condition) {[
            """
            class{'kubernetes': version => '1.11.0'}
              """
          ]}

          it {should contain_file(service_file).with_content(/#{Regexp.escape('--insecure-port=0')}/)}
        end
      end

      context 'etc-quorum-read' do
        context 'should exist before 1.11' do
          let(:pre_condition) {[
              """
            class{'kubernetes': version => '1.10.7'}
              """
          ]}

          it {should contain_file(service_file).with_content(/#{Regexp.escape('--etcd-quorum-read=true')}/)}
        end

        context 'should not exist after 1.11' do
          let(:pre_condition) {[
              """
          class{'kubernetes': version => '1.11.0'}
              """
          ]}

          it {should_not contain_file(service_file).with_content(/#{Regexp.escape('--etcd-quorum-read=true')}/)}
        end
      end

      context 'admission-control' do
        context 'should exist before 1.11' do
          let(:pre_condition) {[
          """
          class{'kubernetes': version => '1.10.7'}
          """
          ]}

          it {should contain_file(service_file).with_content(/#{Regexp.escape('--admission-control=')}/)}
        end

        context 'should not exist after 1.11' do
          let(:pre_condition) {[
            """
            class{'kubernetes': version => '1.11.0'}
              """
          ]}
          let(:params) { {'admission_control' => ['NamespaceLifecycle', 'LimitRanger', 'foo'] } }

          it {should_not contain_file(service_file).with_content(/#{Regexp.escape('--admission-control=')}/)}
          it {should contain_file(service_file).with_content(/#{Regexp.escape('--enable-admission-plugins=NamespaceLifecycle,LimitRanger,foo')}/)}
          it {should contain_file(service_file).with_content(/#{Regexp.escape('--disable-admission-plugins=')}/)}
        end
      end
    end
  end

  context 'bound service token flags' do
    context 'with kubernetes 1.12+' do
      let(:pre_condition) {[
        """
            class{'kubernetes': version => '1.12.0', service_account_key_file => '/etc/kubernetes/sa.key'}
        """
      ]}
      it do
        should contain_file(service_file).with_content(/#{Regexp.escape('--service-account-signing-key-file=/etc/kubernetes/sa.key')}/)
        should contain_file(service_file).with_content(/#{Regexp.escape('--service-account-issuer=kubernetes.default.svc')}/)
        should contain_file(service_file).with_content(/#{Regexp.escape('--service-account-api-audiences=kubernetes.default.svc')}/)
      end
    end
    context 'with kubernetes 1.12+ and less then 1.13+ with custom audiences' do
      let(:pre_condition) {[
        """
            class{'kubernetes': version => '1.12.0', service_account_key_file => '/etc/kubernetes/sa.key'}
            class{'kubernetes::apiserver': service_account_api_audiences => ['kubernetes.default.svc', 'my.custom.aud.svc']}
        """
      ]}
      it do
        should contain_file(service_file).with_content(/#{Regexp.escape('--service-account-signing-key-file=/etc/kubernetes/sa.key')}/)
        should contain_file(service_file).with_content(/#{Regexp.escape('--service-account-issuer=kubernetes.default.svc')}/)
        should contain_file(service_file).with_content(/#{Regexp.escape('--service-account-api-audiences=kubernetes.default.svc,my.custom.aud.svc')}/)
      end
    end
    context 'with kubernetes 1.13+ with custom audiences' do
      let(:pre_condition) {[
        """
            class{'kubernetes': version => '1.13.0', service_account_key_file => '/etc/kubernetes/sa.key'}
            class{'kubernetes::apiserver': service_account_api_audiences => ['kubernetes.default.svc', 'my.custom.aud.svc']}
        """
      ]}
      it do
        should contain_file(service_file).with_content(/#{Regexp.escape('--service-account-signing-key-file=/etc/kubernetes/sa.key')}/)
        should contain_file(service_file).with_content(/#{Regexp.escape('--service-account-issuer=kubernetes.default.svc')}/)
        should contain_file(service_file).with_content(/#{Regexp.escape('--api-audiences=kubernetes.default.svc,my.custom.aud.svc')}/)
      end
    end
    context 'with kubernetes before 1.12' do
      let(:pre_condition) {[
        """
            class{'kubernetes': version => '1.11.0', service_account_key_file => '/etc/kubernetes/sa.key'}
        """
      ]}
      it do
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--service-account-signing-key-file=/etc/kubernetes/sa.key')}/)
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--service-account-issuer=')}/)
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--service-account-api-audiences=')}/)
      end
    end
  end

  context 'feature gates' do
    context 'without given feature gates and not enabled pod priority' do
      let(:params) { {'feature_gates' => {}}}
      it 'should have default feature gates' do
        should_not contain_file(service_file).with_content(/#{Regexp.escape('--feature-gates=')}/)
      end
    end

    context 'without given feature gates and enabled pod priority' do
      let(:pre_condition) {[
        """
        class{'kubernetes': enable_pod_priority => true}
        """
      ]}
      let(:version) { '1.6.0' }
      it 'should have default feature gates' do
        should contain_file(service_file).with_content(/#{Regexp.escape('--feature-gates=PodPriority=true')}/)
      end
    end

    context 'with given feature gates' do
      let(:params) { {'feature_gates' => {'foo' => true, 'bar' => true}}}
      it 'should have custom feature gates' do
        should contain_file(service_file).with_content(/#{Regexp.escape('--feature-gates=foo=true,bar=true')}/)
      end
    end
  end

end
