package ez

import (
	"service-status/model"
	"service-status/pkg"
)

const Url string = "https://app.escolaez.com.br"

var Endpoints = []model.EndPoint{
	{
		Path:    "/auth/entrar",
		PathAKA: "Pagina de Login",
		Status:  pkg.GetStatus(Url, "/auth/entrar"),
	}, {
		Path:    "/autocadastro-professor",
		PathAKA: "Pagina de Cadastro do Professor",
		Status:  pkg.GetStatus(Url, "/autocadastro-professor"),
	}, {
		Path:    "/auth/registro/aluno",
		PathAKA: "Pagina de Cadastro do Aluno",
		Status:  pkg.GetStatus(Url, "/auth/registro/aluno"),
	}, {
		Path:    "/esqueci-a-senha",
		PathAKA: "Pagina de Esqueci a Senha",
		Status:  pkg.GetStatus(Url, "/esqueci-a-senha"),
	},
}
