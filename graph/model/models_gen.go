// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateProductInput struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type Product struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Stock *int    `json:"stock"`
}
