# Task manager

## Project description
The project was created to reinforce my understanding of basic Go syntax, application architecture practices, and to explore a new topic: Structures.

## How to install and run the project
1. Clone the repository
```go
git clone https://github.com/Viperesss/task-manager.git
cd task-manager
```

2. Install dependencies
```go
go mod tidy
```

3. Run the application
```go
go run main.go
```

## Adding a note via file
You need to create or transfer a file with a note (must be .txt) to the repository.
The note in it should contain the following entry format:
```txt
Task name;tag1,tag2;priority
```
Example:
```txt
Buy milk;home,food;2
Finish project;coding,study;5
```