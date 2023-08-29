type CustModel struct{
	ctx context.Context
	mclient *mongo.Client
	mongoCollection *mongo.Collection
}

ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	mongoConnection := options.Client().ApplyURI(constants.ConnectionString)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err!=nil {
		log.Fatal(err.Error())
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err.Error())
	}
	
	fmt.Println("Database Connected")
	// collection := mongoClient.Database("proto").Collection("customer")
	
	defer   mongoClient.Disconnect(ctx)
	fmt.Println("Done")