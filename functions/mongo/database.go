package mongof

// operator docs ($set, $mul) - https://www.mongodb.com/docs/manual/reference/operator/
import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupCreateCollectionsAndStuffForMongoDB() {

}
func CollectionExists(colName string, db *mongo.Database) bool {
	list, err := db.ListCollectionNames(context.TODO(), bson.M{})	
	if err != nil {
		return false
	}
	for _, v := range list {
		if colName == v {
			return false
		}
	}
	return false
}
func UpdateByID(id interface{}, update interface{}, options *options.UpdateOptions, db *mongo.Database, collectionName string) (*mongo.UpdateResult, error) {
	coll := db.Collection(collectionName)
	result, err := coll.UpdateByID(context.TODO(), id, update, options)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func Clone(options *options.CollectionOptions, db *mongo.Database, collectionName string) (*mongo.Collection, error) {
	coll := db.Collection(collectionName)
	result, err := coll.Clone(options)
	return result, err
}
func DeleteMany(documents []interface{}, options *options.DeleteOptions, db *mongo.Database, collectionName string) (*mongo.DeleteResult, error) {
	coll := db.Collection(collectionName)
	/* name := []interface{}{
		bson.D{{"title", "hello"}},
		bson.D{{"title", "helloss"}},
	}
	*/

	result, err := coll.DeleteMany(context.TODO(), documents, options)
	return result, err
}
func UpdateMany(documents []interface{}, options *options.UpdateOptions, db *mongo.Database, collectionName string) (*mongo.UpdateResult, error) {
	coll := db.Collection(collectionName)
	/* name := []interface{}{
		bson.D{{"title", "hello"}},
		bson.D{{"title", "helloss"}},
	}
	*/
	result, err := coll.UpdateMany(context.TODO(), documents, options)
	return result, err
}
func InserMany(documents []interface{}, options *options.InsertManyOptions, db *mongo.Database, collectionName string) (*mongo.InsertManyResult, error) {
	coll := db.Collection(collectionName)
	/* name := []interface{}{
		bson.D{{"title", "hello"}},
		bson.D{{"title", "helloss"}},
	}
	*/
	result, err := coll.InsertMany(context.TODO(), documents, options)
	return result, err
}
func CountDocuments(filter interface{}, options *options.CountOptions, db *mongo.Database, collectionName string) (int64, error) {
	coll := db.Collection(collectionName)
	num, err := coll.CountDocuments(context.TODO(), filter, options)
	return num, err
}
func GetColl(collName string, db *mongo.Database) *mongo.Collection {
	return db.Collection(collName)
}
func FindOneAndDelete(filter interface{}, options *options.FindOneAndDeleteOptions, db *mongo.Database, collectionName string) (bson.M, error) {
	coll := db.Collection(collectionName)
	var result bson.M
	err := coll.FindOneAndDelete(context.TODO(), filter, options).Decode(&result)
	//if err != nil {
	//	if err == mongo.ErrNoDocuments {
	//		return bson.M{}, err
	//	}
	//	return bson.M{}, err
	//}
	return result, err
}
func FindOneAndUpdate(doc interface{}, filter interface{}, options *options.FindOneAndUpdateOptions, db *mongo.Database, collectionName string) (bson.M, error) {
	coll := db.Collection(collectionName)
	var result bson.M
	err := coll.FindOneAndUpdate(context.TODO(), filter, doc, options).Decode(&result)
	//if err != nil {
	//	if err == mongo.ErrNoDocuments {
	//		return bson.M{}, err
	//	}
	//	return bson.M{}, err
	//}
	return result, err
}
func FindOneAndReplace(updateDoc interface{}, filter interface{}, options *options.FindOneAndReplaceOptions, db *mongo.Database, collectionName string) (bson.M, error) {
	coll := db.Collection(collectionName)
	var result bson.M
	err := coll.FindOneAndReplace(context.TODO(), filter, updateDoc, options).Decode(&result)
	//if err != nil {
	//	if err == mongo.ErrNoDocuments {
	//		return bson.M{}, err
	//	}
	//	panic(err)
	//}
	return result, err
}
func Find(options *options.FindOptions, filter interface{}, db *mongo.Database, collectionName string) ([]bson.M, error) {
	coll := db.Collection(collectionName)
	cursor, err := coll.Find(context.TODO(), filter, options)
	//if err != nil {
	//	panic(err)
	//}
	// end find

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
	return results, err
}
func FindOne(doc interface{}, options *options.FindOneOptions, db *mongo.Database, collectionName string) (bson.M, error) {
	coll := db.Collection(collectionName)
	var result bson.M
	err := coll.FindOne(context.TODO(), doc, options).Decode(&result)
	//if err != nil {
	//	if err == mongo.ErrNoDocuments {
	//		return bson.M{}, nil
	//	}
	//	return nil, err
	//}
	return result, err

}
func DeleteOne(filter interface{}, options *options.DeleteOptions, db *mongo.Database, collectionName string) (*mongo.DeleteResult, error) {
	coll := db.Collection(collectionName)
	result, err := coll.DeleteOne(context.TODO(), filter, options)
	//if err != nil {
	//	panic(err)
	//}
	return result, err
}
func ReplaceOne(replacement interface{}, filter interface{}, options *options.ReplaceOptions, db *mongo.Database, collectionName string) (*mongo.UpdateResult, error) {
	coll := db.Collection(collectionName)
	result, err := coll.ReplaceOne(context.TODO(), filter, replacement, options)
	//if err != nil {
	//	panic(err)
	//}
	return result, err
}
func UpdateOne(doc interface{}, filter interface{}, options *options.UpdateOptions, db *mongo.Database, collectionName string) (*mongo.UpdateResult, error) { // example update object -  interface{}{{"$set", interface{}{{"avg_rating", 4.4}}}}
	coll := db.Collection(collectionName)
	result, err := coll.UpdateOne(context.TODO(), filter, doc, options)
	//if err != nil {
	//	panic(err)
	//}
	return result, err
}
func InsertOne(doc interface{}, options *options.InsertOneOptions, db *mongo.Database, collectionName string) (*mongo.InsertOneResult, error) {
	coll := db.Collection(collectionName)
	result, err := coll.InsertOne(context.TODO(), doc, options)
	//if err != nil {
	//	panic(err)
	//}
	return result, err
}

func RunCommand(command interface{}, options *options.RunCmdOptions, db *mongo.Database) bson.M {
	//interface{}{{"create", "s4daccounts"}}
	var result bson.M
	if err := db.RunCommand(context.TODO(), command, options).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result
}
