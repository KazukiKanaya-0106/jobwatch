package config

const DefaultFilename = "jobwatch.yaml"

const DefaultTemplate = `projects:
  name: my_project
  tags:
    - monitoring

run:
  log_tail: 80

notify:
  on_success: true
  on_failure: true
  channels:
    - type: slack
      webhook_url: ${JOBWATCH_SLACK_WEBHOOK_URL}
`
