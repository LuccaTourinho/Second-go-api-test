package main

import (
	"math/rand"
)

// Account representa uma conta bancária com ID, nome do titular, número da conta e saldo.
type Account struct {
	ID        int     `json:"id"`         // Identificador único da conta.
	FirstName string  `json:"first_name"` // Primeiro nome do titular da conta.
	LastName  string  `json:"last_name"`  // Sobrenome do titular da conta.
	Number    int64   `json:"number"`     // Número da conta bancária.
	Balance   float64 `json:"balance"`    // Saldo atual da conta bancária.
}

// newAccount cria uma nova conta bancária.
// Recebe o primeiro e o último nome do titular e retorna um ponteiro para uma instância de Account.
// A conta é inicializada com um ID e número de conta aleatórios, mas com saldo zero.
func newAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),     // Gera um ID aleatório entre 0 e 9999.
		FirstName: firstName,            // Define o primeiro nome do titular da conta.
		LastName:  lastName,             // Define o sobrenome do titular da conta.
		Number:    int64(rand.Intn(100000)), // Gera um número de conta aleatório entre 0 e 99999.
	}
}
