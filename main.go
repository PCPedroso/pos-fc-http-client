package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Criando um contexto
	ctx := context.Background()
	// Definindo um tempo para o Timeout
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// Faz uma requisição utilizando o contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
