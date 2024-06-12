# pulumi

## cli setup

```sh
brew update && brew upgrade pulumi
```

## aws setup

```sh
export CAAS_ACCESS=""
export CAAS_SECRET=""
export CAAS_REGION=us-east-1
pulumi login 's3://pulumi-demo-local?region=us-east-1'    
pulumi stack init --secrets-provider="awskms://<key>?region=us-east-1" 
```

## config setup

```sh
pulumi config set vpcName local-pulumi-vpc
pulumi config set clusterName local-pulumi-eks
pulumi config set minSize 2
pulumi config set maxSize 3
pulumi config set instanceType t3.medium
```

## preview

```sh
pulumi preview
```

## apply

```sh
pulumi up
```
