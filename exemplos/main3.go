package exemplos

import (
	"context"
	"fmt"
	"time"
)

func main_2() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "key", "value")

	defer func() {
		if err := recover(); err != nil {	
			fmt.Println("Recover:", err)}
	}()

	fmt.Println("Iniciando processamento")

	processar(ctx)
}

func processar(ctx context.Context) {
	tenantId := ctx.Value("tenantId")
	fmt.Println("Processando Tenant ID:", tenantId)
	time.Sleep(2 * time.Second)
	fmt.Println("Processamento concluído")
}
