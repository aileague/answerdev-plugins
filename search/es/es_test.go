package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"testing"
)

const (
	testIndex = "answer_index"
)

var (
	testEndpoints = []string{""}
	testUsername  = ""
	testPassword  = ""
)

func TestSearchEngine_Index(t *testing.T) {
	operator, err := NewOperator(testEndpoints, testUsername, testPassword)
	if err != nil {
		t.Fatal(err)
	}
	err = operator.CreateIndex(context.Background(), testIndex, indexJson)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchEngine_SaveDoc(t *testing.T) {
	operator, err := NewOperator(testEndpoints, testUsername, testPassword)
	if err != nil {
		t.Fatal(err)
	}
	err = operator.SaveDoc(context.Background(), testIndex, "1", &AnswerPostDoc{
		Id:          "10010000000001587",
		ObjectID:    "10010000000001587",
		Title:       "How to build new answer with plugin?",
		Type:        "question",
		Content:     "I need to build new answer with plugin, but I don't know how to do it.",
		UserID:      "10040000000000198",
		QuestionID:  "10010000000001587",
		Answers:     5,
		Status:      1,
		Views:       156,
		Created:     1687909352,
		Active:      1,
		Score:       2,
		HasAccepted: false,
		Tags:        []string{"go", "js"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchEngine_QueryDoc(t *testing.T) {
	operator, err := NewOperator(testEndpoints, testUsername, testPassword)
	if err != nil {
		t.Fatal(err)
	}
	doc, err := operator.QueryDoc(context.Background(), testIndex, elastic.NewMatchAllQuery(), nil, nil, 0, 5)
	if err != nil {
		t.Fatal(err)
	}
	for i, hit := range doc.Hits.Hits {
		data, _ := hit.Source.MarshalJSON()
		t.Logf("%d: %+v", i, string(data))
	}
}
