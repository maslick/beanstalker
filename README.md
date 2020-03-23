# =beanstalker=
AWS Beanstalk Multi-container Docker platform Demo

## Installation
```
eb init -i --profile personal
eb create testing --profile personal -r eu-central-1 --elb-type application
eb deploy testing --profile personal
```
