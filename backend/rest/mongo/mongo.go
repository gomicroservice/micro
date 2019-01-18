package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client, err = mongo.Connect(context.Background(), "ds247410.mlab.com:47410/gomicroservice", nil)
