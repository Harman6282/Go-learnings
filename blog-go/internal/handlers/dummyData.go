package handlers

import "github.com/Harman6282/blog-go/internal/types"

var Blogs = []types.Blog{
	{Id: "1", Title: "Getting Started with Go", Content: "This blog explains how to start learning Golang from scratch."},
	{Id: "2", Title: "Why Go is Fast", Content: "Go is fast because of compiled binaries and efficient concurrency."},
	{Id: "3", Title: "Understanding Goroutines", Content: "Goroutines are lightweight threads managed by the Go runtime."},
	{Id: "4", Title: "Channels in Go", Content: "Channels are used to communicate between goroutines safely."},
	{Id: "5", Title: "Go vs Node.js", Content: "A comparison between Go and Node.js for backend development."},
	{Id: "6", Title: "Building REST APIs in Go", Content: "Learn how to build REST APIs using net/http."},
	{Id: "7", Title: "Project Structure in Go", Content: "Best practices for organizing Go backend projects."},
	{Id: "8", Title: "Error Handling in Go", Content: "Go handles errors explicitly using return values."},
	{Id: "9", Title: "Pointers Made Easy", Content: "A beginner-friendly explanation of pointers in Go."},
	{Id: "10", Title: "Structs and Methods", Content: "How structs and methods work together in Go."},
	{Id: "11", Title: "Interfaces Explained", Content: "Interfaces define behavior, not implementation."},
	{Id: "12", Title: "Working with JSON in Go", Content: "Encoding and decoding JSON using encoding/json."},
	{Id: "13", Title: "Go Concurrency Model", Content: "Understanding Go’s CSP-based concurrency model."},
	{Id: "14", Title: "Middleware in Go", Content: "How to create middleware using higher-order functions."},
	{Id: "15", Title: "Authentication Basics", Content: "Basic concepts of authentication in backend systems."},
	{Id: "16", Title: "Using Context in Go", Content: "Context helps manage request lifecycles and cancellations."},
	{Id: "17", Title: "Testing in Go", Content: "Writing unit tests using the testing package."},
	{Id: "18", Title: "Logging Best Practices", Content: "Structured logging techniques for Go applications."},
	{Id: "19", Title: "Deploying Go Apps", Content: "Steps to deploy Go applications to production."},
	{Id: "20", Title: "Common Go Mistakes", Content: "Common mistakes beginners make while learning Go."},
}
