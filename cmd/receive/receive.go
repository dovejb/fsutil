package main

import (
	"context"
	"flag"
	"os"

	"github.com/dovejb/fsutil"
	"github.com/dovejb/fsutil/util"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		panic("dest path not set")
	}

	ctx := context.Background()
	s := util.NewProtoStream(ctx, os.Stdin, os.Stdout)

	if err := fsutil.Receive(ctx, s, flag.Args()[0], fsutil.ReceiveOpt{}); err != nil {
		panic(err)
	}
}
