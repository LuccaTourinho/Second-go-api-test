package main

import (
	"encoding/json" // Importa o pacote para codificação e decodificação JSON.
	"fmt"          // Importa o pacote para formatar strings.
	"log"          // Importa o pacote para registrar logs de eventos.
	"net/http"     // Importa o pacote para criar servidores HTTP.

	"github.com/gorilla/mux" // Importa o pacote mux para roteamento HTTP.
)

// WriteJSON envia uma resposta JSON ao cliente com o status especificado.
// Recebe um http.ResponseWriter para escrever a resposta, um código de status
// e um valor que será codificado em JSON.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json") // Define o tipo de conteúdo como JSON.
	w.WriteHeader(status) // Define o código de status da resposta.
	return json.NewEncoder(w).Encode(v) // Codifica o valor em JSON e escreve na resposta.
}

// apiFunc é um tipo de função que recebe um ResponseWriter e um Request,
// e retorna um erro.
type apiFunc func(http.ResponseWriter, *http.Request) error

// APIError representa um erro que será enviado como resposta JSON.
type APIError struct {
	Error string // Mensagem de erro.
}

// makeHTTPHandleFunc cria um manipulador HTTP a partir de uma função apiFunc.
// Ele trata erros retornados pela função, enviando uma resposta JSON apropriada.
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{err.Error()}) // Se houver erro, responde com status 400.
		}
	}
}

// APIServer representa o servidor da API com um endereço de escuta.
type APIServer struct {
	listenAddress string // Endereço onde o servidor irá escutar requisições.
	store         Storage
}

// NewAPIServer cria e retorna uma nova instância de APIServer.
// Recebe o endereço de escuta como parâmetro.
func NewAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

// Run inicia o servidor HTTP, configurando as rotas e escutando requisições.
// Retorna um erro se ocorrer ao tentar iniciar o servidor.
func (s *APIServer) Run() error {
	router := mux.NewRouter() // Cria um novo roteador.

	// Define a rota "/accounts" e associa a função handleAccount a ela.
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	// Define a rota "/account/{id}" para obter dados de uma conta específica.
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Println("API server running on ", s.listenAddress) // Loga que o servidor está em execução.

	// Inicia o servidor HTTP escutando no endereço especificado.
	return http.ListenAndServe(s.listenAddress, router)
}

// handleAccount manipula requisições para a rota "/accounts".
// Dependendo do método HTTP (GET, POST, DELETE), chama a função correspondente.
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r) // Manipula requisição GET.
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r) // Manipula requisição POST.
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r) // Manipula requisição DELETE.
	}

	// Retorna um erro se o método HTTP não for suportado.
	return fmt.Errorf("method not supported %s", r.Method)
}

// handleGetAccount manipula requisições GET para "/accounts".
//Retorna os dados da conta solicitada, se implementado.
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"] // Extrai o ID da rota.

	fmt.Println(id) // Exibe o ID da conta no console.

	// Retorna uma conta fictícia com status OK. 
	// Deve ser implementado para buscar os dados reais.
	return WriteJSON(w, http.StatusOK, &Account{})
}

// handleCreateAccount manipula requisições POST para "/accounts".
// Cria uma nova conta, se implementado.
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil // Implementação futura.
}

// handleDeleteAccount manipula requisições DELETE para "/accounts".
// Exclui uma conta existente, se implementado.
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil // Implementação futura.
}

// handleTransfer manipula requisições para transferências, se implementado.
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil // Implementação futura.
}
