package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Usuario representa um usuário
type Usuario struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

// GerenciadorUsuarios gerencia operações CRUD de usuários
type GerenciadorUsuarios struct {
	arquivo string
	usuarios []Usuario
}

// Inicializa o gerenciador de usuários
func NovoGerenciadorUsuarios(arquivo string) *GerenciadorUsuarios {
	g := &GerenciadorUsuarios{arquivo: arquivo}
	g.carregarUsuarios()
	return g
}

// Carrega os usuários do arquivo JSON
func (g *GerenciadorUsuarios) carregarUsuarios() {
	file, err := os.Open(g.arquivo)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&g.usuarios)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}
}

// Salva os usuários no arquivo JSON
func (g *GerenciadorUsuarios) salvarUsuarios() {
	file, err := os.Create(g.arquivo)
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(g.usuarios)
	if err != nil {
		fmt.Println("Erro ao codificar JSON:", err)
		return
	}
}

// Adiciona um novo usuário
func (g *GerenciadorUsuarios) AdicionarUsuario(nome string, idade int) {
	g.usuarios = append(g.usuarios, Usuario{Nome: nome, Idade: idade})
	g.salvarUsuarios()
	fmt.Println("USUÁRIO ADICIONADO COM SUCESSO!")
}

// Lista todos os usuários
func (g *GerenciadorUsuarios) ListarUsuarios() {
	fmt.Println("LISTA DE USUÁRIOS:")
	for _, usuario := range g.usuarios {
		fmt.Printf("* NOME: %s, IDADE: %d\n", usuario.Nome, usuario.Idade)
	}
}

// Atualiza um usuário existente
func (g *GerenciadorUsuarios) AtualizarUsuario(nomeAntigo, novoNome string, novaIdade int) {
	for i, usuario := range g.usuarios {
		if usuario.Nome == nomeAntigo {
			usuario.Nome = novoNome
			usuario.Idade = novaIdade
			g.usuarios[i] = usuario
			g.salvarUsuarios()
			fmt.Println("USUÁRIO ATUALIZADO COM SUCESSO!")
			return
		}
	}
	fmt.Println("USUÁRIO NÃO ENCONTRADO!")
}

// Exclui um usuário
func (g *GerenciadorUsuarios) ExcluirUsuario(nome string) {
	for i, usuario := range g.usuarios {
		if usuario.Nome == nome {
			g.usuarios = append(g.usuarios[:i], g.usuarios[i+1:]...)
			g.salvarUsuarios()
			fmt.Println("🗑 USUÁRIO EXCLUÍDO COM SUCESSO!")
			return
		}
	}
	fmt.Println("USUÁRIO NÃO ENCONTRADO!")
}

func exibirMenu() {
	fmt.Println("\nMENU:")
	fmt.Println("1. ADICIONAR USUÁRIO")
	fmt.Println("2. LISTAR USUÁRIOS")
	fmt.Println("3. ATUALIZAR USUÁRIO")
	fmt.Println("4. EXCLUIR USUÁRIO")
	fmt.Println("5. SAIR")
}

func main() {
	gerenciador := NovoGerenciadorUsuarios("usuarios.json")

	for {
		exibirMenu()
		var opcao int
		fmt.Print("ESCOLHA UMA OPÇÃO:\n>>> ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			var nome string
			var idade int
			fmt.Print("DIGITE O NOME: ")
			fmt.Scanln(&nome)
			fmt.Print("DIGITE A IDADE: ")
			fmt.Scanln(&idade)
			gerenciador.AdicionarUsuario(nome, idade)
		case 2:
			gerenciador.ListarUsuarios()
		case 3:
			var nomeAntigo, novoNome string
			var novaIdade int
			fmt.Print("DIGITE O NOME A SER ATUALIZADO: ")
			fmt.Scanln(&nomeAntigo)
			fmt.Print("DIGITE O NOVO NOME: ")
			fmt.Scanln(&novoNome)
			fmt.Print("DIGITE A NOVA IDADE: ")
			fmt.Scanln(&novaIdade)
			gerenciador.AtualizarUsuario(nomeAntigo, novoNome, novaIdade)
		case 4:
			var nome string
			fmt.Print("DIGITE O NOME DO USUÁRIO A SER EXCLUÍDO: ")
			fmt.Scanln(&nome)
			gerenciador.ExcluirUsuario(nome)
		case 5:
			fmt.Println("SAINDO...")
			return
		default:
			fmt.Println("OPÇÃO INVÁLIDA. TENTE NOVAMENTE!")
		}
	}
}
