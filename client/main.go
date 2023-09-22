package main

import (
	"context"
	"log"

	feast "github.com/feast-dev/feast/sdk/go"
)

func main() {
	cli, err := feast.NewGrpcClient("127.0.0.1", 6379)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	req := feast.OnlineFeaturesRequest{
		Features: []string{
			"driver_hourly_stats:acc_rate",
			"transformed_conv_rate:conv_rate_plus_val1",
			"transformed_conv_rate:conv_rate_plus_val2",
		},
		Entities: []feast.Row{
			{"driver_id": feast.Int64Val(1001), "val_to_add": feast.Int64Val(1000), "val_to_add_2": feast.Int64Val(2000)},
			{"driver_id": feast.Int64Val(1002), "val_to_add": feast.Int64Val(1001), "val_to_add_2": feast.Int64Val(2002)},
		},
		Project: "feast",
	}

	res, err := cli.GetOnlineFeatures(ctx, &req)
	if err != nil {
		panic(err)
	}

	// returns a list of rows
	out := res.Rows()
	log.Println(out)

}
