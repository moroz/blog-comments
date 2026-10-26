package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dynamodb.NewTable(ctx, "comments", &dynamodb.TableArgs{
			Ttl: &dynamodb.TableTtlArgs{
				AttributeName: pulumi.String("ttl"),
				Enabled:       pulumi.Bool(true),
			},
			Attributes: &dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("url"),
					Type: pulumi.String("S"),
				},
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
}
