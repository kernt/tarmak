define fluent_bit::output (
  Hash $config = {},
){
  include ::fluent_bit

  validate_hash($config)
  $path = $::fluent_bit::path

  if $config['elasticsearch'] {

    $elasticsearch = $config['elasticsearch']

    if $elasticsearch['awsESProxy'] {

      ::aws_es_proxy::instance{ $name:
        tls          => $elasticsearch['tls'],
        dest_port    => $elasticsearch['port'],
        dest_address => $elasticsearch['host'],
        listen_port  => $elasticsearch['awsESProxy']['port'],
      }

    }

    file { "/etc/td-agent-bit/td-agent-bit-output-${name}.conf":
      ensure  => file,
      mode    => '0644',
      owner   => 'root',
      group   => 'root',
      content => template('fluent_bit/td-agent-bit-output.conf.erb'),
    }

  } else {

    file { "/etc/td-agent-bit/td-agent-bit-${name}.conf":
      ensure  => file,
      mode    => '0644',
      owner   => 'root',
      group   => 'root',
      content => template('fluent_bit/td-agent-bit-output.conf.erb'),
    }

  }
}
