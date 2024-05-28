package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/auth"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

func init() {
	os.Setenv("WEAVIATE_INSTANCE_URL", "localhost:8086")
	os.Setenv("WEAVIATE_SCHEME", "http")
	os.Setenv("WEAVIATE_API_KEY", "soulteary")
}

func CreateClient(host string, scheme string, key string) (*weaviate.Client, error) {
	cfg := weaviate.Config{
		Host:       host,
		Scheme:     scheme,
		AuthConfig: auth.ApiKey{Value: key},
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateDB(client *weaviate.Client) error {
	classObj := &models.Class{
		Class:      "TraditionalFestival",
		Vectorizer: "text2vec-transformers",
	}

	// add the schema
	err := client.Schema().ClassCreator().WithClass(classObj).Do(context.Background())
	if err != nil {
		return err
	}

	buf, err := os.ReadFile("./traditional-festival.json")
	if err != nil {
		return err
	}

	var items []map[string]string
	err = json.Unmarshal(buf, &items)
	if err != nil {
		return err
	}

	objects := make([]*models.Object, len(items))
	for i := range items {
		objects[i] = &models.Object{
			Class: "TraditionalFestival",
			Properties: map[string]any{
				"SolarTerms":  items[i]["SolarTerms"],
				"Title":       items[i]["Title"],
				"Author":      items[i]["Author"],
				"Poem":        items[i]["Poem"],
				"Description": items[i]["Description"],
			},
		}
	}

	batchRes, err := client.Batch().ObjectsBatcher().WithObjects(objects...).Do(context.Background())
	if err != nil {
		return err
	}
	for _, res := range batchRes {
		if res.Result.Errors != nil {
			return fmt.Errorf("error: %v", res.Result.Errors)
		}
	}

	return nil
}

func Query(client *weaviate.Client) {
	fields := []graphql.Field{
		{Name: "solarTerms"},
		{Name: "title"},
		{Name: "author"},
		{Name: "poem"},
		{Name: "description"},
	}

	nearText := client.GraphQL().
		NearTextArgBuilder().
		WithConcepts([]string{"秋天看红叶"})

	result, err := client.GraphQL().Get().
		WithClassName("TraditionalFestival").
		WithFields(fields...).
		WithNearText(nearText).
		WithLimit(2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	buf, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(buf))
}

func main() {
	hostURL := os.Getenv("WEAVIATE_INSTANCE_URL")
	scheme := os.Getenv("WEAVIATE_SCHEME")
	apikey := os.Getenv("WEAVIATE_API_KEY")

	client, err := CreateClient(hostURL, scheme, apikey)
	if err != nil {
		panic(err)
	}
	// CreateDB(client)
	Query(client)
}
