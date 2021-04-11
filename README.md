# Description

The AWS Enumerator was created for service enumeration and info dumping for investigations of penetration testers during Black-Box  testing. The tool should speed up the process of Cloud reviewing in case the security researcher compromised AWS Account Credentials. 

AWS Enumerator counts more than 600 API Calls ( reading actions `Get`,  `List`, `Describe` etc... ) which are going to be extended. 

The tool provides interface for result analysis. All results are saved in json files (one time "Database").

# Tool Install

If you have Go installed and configured (i.e. with `$GOPATH/bin` in your `$PATH`):

```
go get -u github.com/shabarkin/aws-enumerator
```

# Tool Usage

## Credentials setup

To setup credentials you should use `cred` subcommand and supply credentials: 

```bash
./aws-enumerator cred -aws_access_key_id AKIA***********XKU -aws_region us-west-2 -aws_secret_access_key kIm6m********************5JPF
```

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/79476d30-2660-4074-aa83-a1a8b98d5553/Screenshot_2021-04-10_at_14.43.51.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/79476d30-2660-4074-aa83-a1a8b98d5553/Screenshot_2021-04-10_at_14.43.51.png)

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/fb21e481-401a-4752-929d-d2757a60eb90/Screenshot_2021-04-10_at_14.45.51.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/fb21e481-401a-4752-929d-d2757a60eb90/Screenshot_2021-04-10_at_14.45.51.png)

It creates `.env` file which is loaded to global variables each time you call `enum` subcommand.

**WARNING:** If you set these values `AWS_REGION`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN` in global variables manually before running the tool, it will not be able to load AWS Credentials specified in `.env` file ( It can't overwrite global variables ).

## Enumeration

To enumerate all services use `enum` subcommand and supply `all` value or `iam,s3,sts,rds` ( supply without spaces between commas ) etc ...

```bash
./aws-enumerator enum -services all
```

 if you want to check specific services (you can supply up to 10):

```bash
./aws-enumerator enum -services iam,sts,rds
```

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3dd7453e-671d-4c45-af62-7d3f969a463c/Screenshot_2021-04-10_at_13.36.56.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3dd7453e-671d-4c45-af62-7d3f969a463c/Screenshot_2021-04-10_at_13.36.56.png)

(`-speed` is optional, the default value is `normal` ) There is 3 speed choices `slow`, `normal`, `fast` 

```bash
./aws-enumerator enum -services all -speed slow
```

## Analyse

To analyse all gathered information use `dump` subcommand: ( Use `all` for quick overview of available API calls )

```bash
./aws-enumerator dump -services all
```

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2eb7a519-50a0-4909-b6a7-9ed33e52cb4e/Screenshot_2021-04-10_at_13.56.12.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2eb7a519-50a0-4909-b6a7-9ed33e52cb4e/Screenshot_2021-04-10_at_13.56.12.png)

Analyse specific services (you can supply up to 10) `iam,s3,sts` etc ...

```bash
./aws-enumerator dump -services iam,s3,sts
```

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/61f96ceb-5c81-489e-a770-5543af03af0b/Screenshot_2021-04-10_at_14.03.16.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/61f96ceb-5c81-489e-a770-5543af03af0b/Screenshot_2021-04-10_at_14.03.16.png)

To filter the API calls use `-filter` option, start typing the name of API call (`GetA` ...): 

```bash
./aws-enumerator dump -services iam -filter GetA
```

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0cb9420c-adb0-4b59-99b9-4d7174ec0081/Screenshot_2021-04-10_at_14.06.18.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0cb9420c-adb0-4b59-99b9-4d7174ec0081/Screenshot_2021-04-10_at_14.06.18.png)

To retrieve the result of API call use `-print` option

```bash
./aws-enumerator dump -services iam -filter ListS -print
```

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/56344d59-19d0-487b-8afe-c746b8d7348e/Screenshot_2021-04-10_at_14.08.01.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/56344d59-19d0-487b-8afe-c746b8d7348e/Screenshot_2021-04-10_at_14.08.01.png)

## Demo Video

[https://twitter.com/shabarkin/status/1379717210650861573](https://twitter.com/shabarkin/status/1379717210650861573)

# Support

[AWS Enumerator](https://www.notion.so/AWS-Enumerator-88d7bce36d78416c8f67c089ad9ec5dc)
