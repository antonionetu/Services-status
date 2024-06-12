package instituto_terra

import (
	"service-status/model"
	"service-status/pkg"
)

const Url string = "https://instituto-terra.escolaez.com.br"

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
		Path:    "/cadastro/aluno",
		PathAKA: "Pagina de Cadastro do Aluno",
		Status:  pkg.GetStatus(Url, "/cadastro/aluno"),
	}, {
		Path:    "/esqueci-a-senha",
		PathAKA: "Pagina de Esqueci a Senha",
		Status:  pkg.GetStatus(Url, "/esqueci-a-senha"),
	},
}
