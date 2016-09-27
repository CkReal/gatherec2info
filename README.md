# Name
gatherec2info

# Overview
Default EC2 information is too many information.  
This tool gets main information from EC2 instance, and make TSV file.

# Description
This tool has two mode how to gather information.  
1. Interactive Mode
2. Batch Mode

Interactive Mode is assumed using AWS Begginer or
 User who not learning EC2 detail Information yet.   

Batch Mode is assumed using like daily cron case.  
For example, update EC2 Information to Excel or Google Spreads sheets everyday.

# Install
```
$ go get github.com/CkReal/gatherec2info
```

# Requirement
Prepare AWS Account Key and AWS Secret Access Key each AWS Account.  

## Setting Case 1
Set Environment variable of `AWS_SHARED_CREDENTIALS_FILE`
```
$ export AWS_SHARED_CREDENTIALS_FILE=~/.aws/credentials
```

## Setting Case 2
Write AWS Account Information to `~/.aws/credentials`
```
$ cat ~/.aws/credentials
[sample1]
aws_access_key_id = AK***************
aws_secret_access_key = ***************
[sample2]
aws_access_key_id = AK***************
aws_secret_access_key = ***************
```

# How to Use
## help
```
$ gatherec2info help
This tool is gathering EC2 Information.

Usage:
  gatherec2info [flags]
  gatherec2info [command]

Available Commands:
  batch       gatherec2info Batch Mode
  version     Display This Tools's Version Information

Use "gatherec2info [command] --help" for more information about a command.
```

## version help
```
$ gatherec2info version --help
Display This Tools's Version Information
    *.*.*
    | | |__  Patch Version(ex. typo,refactoring)
    | |____  Minor Version(ex. add func)
    |______  Major Version

Usage:
  gatherec2info version [flags]
```

## batch help
```
$ gatherec2info batch --help
gatherec2info Batch Mode

Usage:
  gatherec2info batch [flags]

Flags:
  -o, --output="": OutputFile PATH(default: ${ProfileId}_EC2.tsv)
  -p, --profile="": Profile Id(ex. sample)
  -r, --region="ap-northeast-1": AWS Region(default: ap-northeast-1)
```

# Usage
## Interactive Mode
This is default mode.
```
$ gatherec2info
#####################################
###### Gather EC2 Information #######
#####################################
Input Profile Id(ex. sample)> sample1
-------------------------------------
[0] default
[1] US East (N. Virginia)
[2] US West (Oregon)
[3] US West (N. California)
[4] EU (Ireland)
[5] EU (Frankfurt)
[6] Asia Pacific (Singapore)
[7] Asia Pacific (Tokyo)
[8] Asia Pacific (Sydney)
[9] Asia Pacific (Seoul)
[10] Asia Pacific (Mumbai)
[11] South America (São Paulo)
Input Region Id(default:Tokyo)> 0
-------------------------------------
[0] ALL
[1] Tags.Name
[2] PrivateIpAddress
[3] PrivateDnsName
[4] PublicIpAddress
[5] PublicDnsName
[6] SubnetId
[7] VpcId
[8] AvailabilityZone
[9] InstanceId
[10] InstanceType
[11] SecurityGroups
[12] ImageId
[13] RootDeviceType
[14] VirtualizationType
[15] KeyName
[16] State
[17] LaunchTime
Input Request Information,comma delimited(ex. 1,2) > 1,2
-------------------------------------
Tags.Name	PrivateIpAddress
sample	10.0.0.1

-------------------------------------
Create [sample1_EC2.tsv] File?(Y/N)> Y
```
### Interactive Result
```
$ cat sample_EC2.tsv
Tags.Name	PrivateIpAddress
sample	10.0.0.1
```

## Batch Mode
Batch mode use `batch` command.

```
$ gatherec2info batch -p sample1 -r ap-northeast-1 -o sample1.txt
```

### Batch Result
```
$ cat sample1.txt
Tags.Name	PrivateIpAddress	PrivateDnsName	PublicIpAddress	PublicDnsName	SubnetId	VpcId	AvailabilityZone	InstanceId	InstanceType	SecurityGroups	ImageId	RootDeviceType	VirtualizationType	KeyName	State	LaunchTime
app1.sample1	10.0.0.1	ip-10-0-0-1.ap-northeast-1.compute.internal	54.238.10.100	ec2-54-238-10-100.ap-northeast-1.compute.amazonaws.com	subnet-12345678	vpc-12345678	ap-northeast-1c	i-12345678	t2.small	app.sample	ami-12345678	ebshvm	sample	running	2016-01-06 12:12:00 +0000 UTC
app2.sample1	10.0.0.1	ip-10-0-0-2.ap-northeast-1.compute.internal	54.199.10.100	ec2-54-199-10-100.ap-northeast-1.compute.amazonaws.com	subnet-12345678	vpc-12345678	ap-northeast-1c	i-12345678	t2.small	app.sample	ami-12345678	ebshvm	sample	running	2016-07-12 18:51:18 +0000 UTC
```
