package main

import (
	"net/url"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dbUrl := os.Getenv("PRODUCTION_DATABASE_URL")
		authToken := os.Getenv("PRODUCTION_DATABASE_TOKEN")
		qs := url.Values{
			"authToken": {authToken},
		}

		_, err := ssm.NewParameter(ctx, "turso-token", &ssm.ParameterArgs{
			Name:  pulumi.String("/production/comments/database/credentials"),
			Type:  pulumi.String(ssm.ParameterTypeSecureString),
			Value: pulumi.String(dbUrl + "?" + qs.Encode()),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
