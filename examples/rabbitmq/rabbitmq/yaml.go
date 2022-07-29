package rabbitmq

const ConfigYaml = `
exchange:
  apply:
    name: apply
    type: direct

queue:
  apply_vp_validation:
    name: apply.vp.validation
    durable: true

binding:
  apply_vp_validation:
    queue: apply_vp_validation
    exchange: apply
    bind_options:
      routing_key: vp_validation

channel:
  apply_vp_validation:
    prefetch: 20
`
