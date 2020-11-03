package model

var client *db.Client
var ctx context.Context

func init()  {
	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL:" "
	}
	opt := option.WithCredentialsFile("")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil{
		log.Fatalln("error initializing app", err)
	}

	client, err= app.Database(ctx)
	if err != nil {
		log.Fatalln("error initializing database client", err)
	}
}