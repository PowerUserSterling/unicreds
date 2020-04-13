[![Build Status](https://travis-ci.org/Versent/unicreds.svg?branch=master)](https://travis-ci.org/Versent/unicreds)

# unicreds

This fork removes the CLI as it is just to be used as an import into other projects .  It also updates the aws-sdk imports to aws-sdk-go-v2.

Unicreds is a command line tool to manage secrets within an AWS account, the aim is to keep securely stored
with your systems and data so you don't have to manage them externally. It uses [DynamoDB](https://aws.amazon.com/dynamodb/) and [KMS](https://aws.amazon.com/kms/) to store and
encrypt these secrets. Access to these keys is controlled using [IAM](https://aws.amazon.com/iam/).

Unicreds is written in [Go](https://golang.org/) and is based on [credstash](https://github.com/fugue/credstash).

# setup

1. Create a KMS key in IAM, using an aws profile you have configured in the aws CLI. You can ommit `--profile` if you use the Default profile.
```
aws --region ap-southeast-2 --profile [yourawsprofile] kms create-key --query 'KeyMetadata.KeyId'
```
**Note:** You will also need to assign permission to users other than the root account to access and use the key see [How to Help Protect Sensitive Data with AWS KMS](https://blogs.aws.amazon.com/security/post/Tx79IILINW04DC/How-to-Help-Protect-Sensitive-Data-with-AWS-KMS).
2. Assign the `credstash` alias to the key using the key id printed when you created the KMS key.
```
aws --region ap-southeast-2 --profile [yourawsprofile] kms create-alias --alias-name 'alias/credstash' --target-key-id "xxxx-xxxx-xxxx-xxxx-xxxx"
```
3. Run unicreds setup to create the dynamodb table in your region, ensure you have your credentials configured using the [awscli](https://aws.amazon.com/cli/).
```
unicreds setup --region ap-southeast-2 --profile [yourawsprofile]
```
**Note:** It is really important to tune DynamoDB to your read and write requirements if you're using unicreds with automation!

# references

* [How to Protect the Integrity of Your Encrypted Data by Using AWS Key Management Service and EncryptionContext](https://blogs.aws.amazon.com/security/post/Tx2LZ6WBJJANTNW/How-to-Protect-the-Integrity-of-Your-Encrypted-Data-by-Using-AWS-Key-Management)

# development

I use `scantest` to watch my code and run tests on save.

```
go get github.com/smartystreets/scantest
```

# testing
You can run unit tests which mock out the KMS and DynamoDB backend using `make test`.

There is an integration test you can run using `make integration`. You must set the `AWS_REGION` (default `us-west-2`), `UNICREDS_KEY_ALIAS` (default `alias/unicreds`), and `UNICREDS_TABLE_NAME` (default `credential-store`) environment variables to point to valid AWS resources.

# auto-versioning

If you've been using unicreds auto-versioning before September 2015, Unicreds had the [same](https://github.com/fugue/credstash/issues/51) [bug](https://github.com/Versent/unicreds/issues/34) as credstash when auto-versioning that results in a sorting error after ten versions. You should be able to run the [credstash-migrate-autoversion.py](https://github.com/fugue/credstash/blob/master/credstash-migrate-autoversion.py) script included in the root of the credstash repository to update your versions prior to using the latest version of unicreds.

# todo

* Add the ability to filter list / getall results using DynamoDB filters, at the moment I just use `| grep blah`.
* Work on the output layout.
* Make it easier to import files

# license

This code is Copyright (c) 2015 Versent and released under the MIT license. All rights not explicitly granted in the MIT license are reserved. See the included LICENSE.md file for more details.
