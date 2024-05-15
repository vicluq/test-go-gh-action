# Golang Lambda Deploy

This action will provide a deploy interface to your Golang AWS lambda code. It checks if the labda function (`FUNCTION_NAME`) currently exists, if so, it will update it, if not, it will create it.

You can edit the workflow file to add more configuration flags to your lambda, such as VPC, subnets, etc.

## Secrets
These are the AWS Secret Key, Key ID and role, which must be added as a **repository secret**.

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_ROLE` (the arn string representing the role that will be given to your lambda)

## Environment Variables
 - `AWS_REGION`: your-aws-region, default: us-east-1
 - `ENTRY_FILE`: your-go-entry-file, default: main.go
 - `FUNCTION_NAME`: your-lambda-function-name

The entry file should be the one containing your `lambda.Start()` call.

Check the following repository with the template in action: [vicluq/test-go-aws-lambda-gh-action](https://github.com/vicluq/test-go-gh-action)
