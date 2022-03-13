package golicenseserver

import (
	"context"
	"fmt"
	"log"

	"github.com/thuanvh/golicenseserver/ent"
	"github.com/thuanvh/golicenseserver/ent/todo"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func Example_Todo() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("Failed to open connection %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("Error %v", err)
	}

	task1, err := client.Todo.Create().SetText("Add GraphQL Example").Save(ctx)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Println(task1)
	task2, err := client.Todo.Create().SetText("Add Tracing Example").Save(ctx)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Println(task2)
	if err := task2.Update().SetParent(task1).Exec(ctx); err != nil {
		log.Fatalf("failed connecting todo2 to its parent: %v", err)
	}

	items, err := client.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("Error in query all task %v", err)
	}
	for _, i := range items {
		fmt.Printf("%d: %q\n", i.ID, i.Text)
	}

	items2, err := client.Todo.Query().Where(todo.HasParent()).All(ctx)
	if err != nil {
		log.Fatalf("Error in checking has parent %v", err)
	}
	for _, i := range items2 {
		fmt.Printf("%d: %q\n", i.ID, i.Text)
	}

	items3, err := client.Todo.Query().Where(todo.Not(todo.HasParent()), todo.HasChildren()).All(ctx)
	if err != nil {
		log.Fatalf("Error in get parent %v", err)
	}
	for _, i := range items3 {
		fmt.Printf("%d: %q\n", i.ID, i.Text)
	}

	// Get a parent item through its children and expect the
	// query to return exactly one item.
	parent, err := client.Todo.
		Query().                 // Query all todos.
		Where(todo.HasParent()). // Filter only those with parents.
		QueryParent().           // Continue traversals to the parents.
		Only(ctx)                // Expect exactly one item.
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	fmt.Printf("%d: %q\n", parent.ID, parent.Text)

	// item, err := client.Todo.Query().Where(todo.HasParent()).QueryParent().Only(ctx)
	// if err != nil {
	// 	log.Fatalf("Error of parent %v", err)
	// }
	// fmt.Printf("%d: %q\n", item.ID, item.Text)
}
