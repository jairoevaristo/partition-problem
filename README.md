# ALGORITMOS E COMPLEXIDADE COMPUTACIONAL

<hr />

## Problema de Partição (Partition Problem)
- Descrição: Dado um conjunto de números inteiros, determine se ele pode ser
particionado em dois subconjuntos cujas somas sejam iguais.

- Objetivo: Verificar se o conjunto pode ser dividido em dois subconjuntos com a
mesma soma total

<hr />

## Aplicabilidade do problema
 - *Divisão de Recursos Financeiros ou Materiais:* O problema de partição pode ser utilizado para verificar se é possível realizar a divisão sem sobras ou desbalanceamento.
 - *Balanceamento de Carga em Sistemas Computacionais:* Resolver o problema de partição ajuda a garantir que a carga de trabalho seja distribuída igualmente, reduzindo o tempo total de processamento.
 - *Planejamento de Produção:* O problema de partição auxilia no balanceamento da utilização das máquinas, maximizando a eficiência.
 - *Organização de Equipes ou Grupos:* O problema pode ser aplicado para garantir uma divisão justa entre as equipes.
 - *Criptografia e Teoria dos Números:* O problema de partição surge em análises e verificações de segurança e integridade de algoritmos.

<hr />

## Para instalar as dependências
 ```bash
go mod tidy
 ```

## Para executar os testes
```bash
go test ./test -v
```

<hr />

> [!IMPORTANT]  
> Ao executar o comando acima irá rodar o arquivo responsável por testar as funções com a implementação dos algoritmos, após a execução será criado um arquivo `table_output.pdf` que contém uma tabela com os resultados de todos os tempos de execução para cada algoritmo.

## Arquivo com a implementação dos algoritmos 
 - `internal/service/algorithm.go`
