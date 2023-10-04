package main

import "os"

type Repository struct {
	args []string
}

func NewRepository(args []string) *Repository {
	return &Repository{args: args}
}

func main() {
	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	repo := NewRepository(os.Args)
	switch command {
	case "help":
		// 显示帮助信息
		repo.Help()
	case "version":
		// 显示版本信息
		repo.Version()
	case "open":
		repo.Open()
	case "read":
		repo.Read()
	case "readall":
		repo.ReadAll()
	default:
		repo.Help()
	}
}
