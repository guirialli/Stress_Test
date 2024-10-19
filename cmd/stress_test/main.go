package main

import (
	"fmt"
	"github.com/guirialli/stress_test/infra/services"
	"github.com/spf13/cobra"
)

func main() {
	var url string
	var requests int
	var concurrency int

	rootCmd := &cobra.Command{
		Use:   "stress test",
		Short: "Um simples testador de carga para serviços web",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Iniciando o stress test em: %s\nConcorrência: %d; Request: %d\n\n", url, concurrency, requests)
			result, _ := services.BenchUrl(url, requests, concurrency)
			fmt.Println(result.GenerateReport())
		},
	}

	rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado")
	rootCmd.Flags().IntVar(&requests, "requests", 0, "Número total de requests")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
