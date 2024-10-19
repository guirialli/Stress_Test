package dtos

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type ResulStatusCode struct {
	mu          sync.Mutex
	status      Report
	totalTime   time.Duration
	request     int
	concurrency int
	errors      int
}

func NewResulStatusCode(request, concurrency int) *ResulStatusCode {
	return &ResulStatusCode{
		status:      make(Report),
		request:     request,
		concurrency: concurrency,
	}
}

func (r *ResulStatusCode) Increment(code int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.status[code]++
}

func (r *ResulStatusCode) Get(code int) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.status[code]
}

func (r *ResulStatusCode) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.status = make(map[int]int)
}

func (r *ResulStatusCode) Status() Report {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.status
}

func (r *ResulStatusCode) TotalTime() time.Duration {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.totalTime
}
func (r *ResulStatusCode) SetTotalTime(totalTime time.Duration) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.totalTime = totalTime
}

func (r *ResulStatusCode) Request() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.request
}

func (r *ResulStatusCode) Concurrency() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.concurrency
}

func (r *ResulStatusCode) IncrementErrors() {
	r.mu.Lock()
	r.errors += 1
	r.mu.Unlock()
}

func (r *ResulStatusCode) GenerateReport() string {
	r.mu.Lock()
	defer r.mu.Unlock()

	var reportBuilder strings.Builder

	reportBuilder.WriteString(fmt.Sprintf("Relatório de Execução\n"))
	reportBuilder.WriteString(fmt.Sprintf("======================\n"))
	reportBuilder.WriteString(fmt.Sprintf("Tempo Total: %s\n", r.totalTime))
	reportBuilder.WriteString(fmt.Sprintf("Total de Requisições: %d\n", r.request))
	reportBuilder.WriteString(fmt.Sprintf("Concorrência: %d\n", r.concurrency))
	reportBuilder.WriteString(fmt.Sprintf("Respostas HTTP 200: %d\n", r.status[200]))
	reportBuilder.WriteString(fmt.Sprintf("Erros HTTP: %d\n", r.errors))

	reportBuilder.WriteString("\nDistribuição de Outros Status HTTP:\n")
	for code, count := range r.status {
		if code != 200 {
			reportBuilder.WriteString(fmt.Sprintf("Status %d: %d vez(es)\n", code, count))
		}
	}

	return reportBuilder.String()
}
