<p align="center">
    <h1 align="center">Projeto Ceres</h1>
    <a href="https://github.com/KillovSky/Ceres/blob/main/LICENSE"><img alt="GitHub License" src="https://img.shields.io/github/license/KillovSky/Ceres?color=blue&label=License&style=flat-square"></a>
    <a href="https://github.com/KillovSky/Ceres"><img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/KillovSky/Ceres?label=Size%20%28With%20.git%20folder%29&style=flat-square"></a>
    <a href="https://api.github.com/repos/KillovSky/Ceres/languages"><img alt="GitHub Languages" src="https://img.shields.io/github/languages/count/KillovSky/Ceres?label=Code%20Languages&style=flat-square"></a>
    <a href="https://github.com/KillovSky/Ceres/blob/main/.github/CHANGELOG.md"><img alt="GitHub Version" src="https://img.shields.io/github/package-json/v/KillovSky/Ceres?label=Latest%20Version&style=flat-square"></a>
    <a href="https://github.com/KillovSky/Ceres/blob/main/.github/CHANGELOG.md"><img alt="Project Codename" src="https://img.shields.io/github/package-json/build_name/KillovSky/Ceres?label=Latest%20Codename"></a>
    <a href="https://github.com/KillovSky/Ceres/blob/main/.github/CHANGELOG.md"><img alt="Last Update" src="https://img.shields.io/github/package-json/build_date/KillovSky/Ceres?label=Latest%20Update"></a>
    <a href="https://github.com/KillovSky/Ceres/commits/main"><img alt="GitHub Commits" src="https://img.shields.io/github/commit-activity/y/KillovSky/Ceres?label=Commits&style=flat-square"></a>
    <a href="https://github.com/KillovSky/Ceres/stargazers/"><img title="GitHub Stars" src="https://img.shields.io/github/stars/KillovSky/Ceres?label=Stars&style=flat-square"></a>
    <a href="https://github.com/KillovSky/Ceres/network/members"><img title="GitHub Forks" src="https://img.shields.io/github/forks/KillovSky/Ceres?label=Forks&style=flat-square"></a>
    <a href="https://github.com/KillovSky/Ceres/watchers"><img title="GitHub Watchers" src="https://img.shields.io/github/watchers/KillovSky/Ceres?label=Watchers&style=flat-square"></a>
    <a href="http://isitmaintained.com/project/killovsky/Ceres"><img alt="Issue Resolution" src="http://isitmaintained.com/badge/resolution/killovsky/Ceres.svg"></a>
    <a href="http://isitmaintained.com/project/killovsky/Ceres"><img alt="Open Issues" src="http://isitmaintained.com/badge/open/killovsky/Ceres.svg"></a>
    <a href="https://hits.seeyoufarm.com"><img src="https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2FKillovSky%2FCeres&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=Views&edge_flat=false"/></a>
    <a href="https://github.com/KillovSky/Ceres/pulls"><img alt="Pull Requests" src="https://img.shields.io/github/issues-pr/KillovSky/Ceres?label=Pull%20Requests&style=flat-square"></a>
    <a href="https://github.com/KillovSky/Ceres/graphs/contributors"><img alt="Contributors" src="https://img.shields.io/github/contributors/KillovSky/Ceres?label=Contribuidores&style=flat-square"></a>
</p>

# O que é?

O Projeto Ceres é um plugin opcional desenvolvido em Golang para o [Projeto Íris](https://github.com/KillovSky/Iris). Este plugin possibilita a implementação de todas as funcionalidades Golang, incluindo a compilação de códigos GO. Com isso, a Ceres elimina a necessidade de modificar o código principal da Íris ou de aprender Node.js (JavaScript) para realizar alterações.

## Requisitos

Para garantir o correto funcionamento do Projeto Ceres, o Projeto Íris deve estar ativo. A versão atual do Projeto Ceres é beta e foi desenvolvida rapidamente para fins de aprendizado, podendo conter erros menores.

1. **Golang**:
   - É recomendada a versão mais recente do Go.
2. **Projeto Íris**:
   - Deve estar instalada e em execução.
3. **Dependências do Projeto Íris**:
   - Instale todas as dependências necessárias do Projeto Íris para assegurar o correto funcionamento da Ceres.

## Instalação

Para instalar as dependências do Projeto Ceres, você pode usar um dos métodos a seguir:

1. **Usando Go**:
   - Execute o seguinte comando:
     ```bash
     go mod tidy
     ```
   - Isso irá baixar e instalar todas as dependências necessárias listadas no `go.mod`.

2. **Usando NPM**:
   - Embora o Projeto Ceres seja desenvolvido em Golang e **NÃO UTILIZE NODE (JS)**, você pode instalar as dependências via NPM também.
   - Utilize o seguinte comando:
     ```bash
     npm run install
     ```
   - Este comando executará o `go mod tidy` por meio do NPM para instalar as dependências do Go.

## Execução

Após a instalação das dependências, você pode executar o Projeto Ceres de duas maneiras:

1. **Usando Go**:
   - Para executar o Projeto Ceres diretamente, utilize o seguinte comando:
     ```bash
     go run main.go
     ```

2. **Usando NPM**:
   - Se preferir, você também pode iniciar a Ceres via NPM com o seguinte comando:
     ```bash
     npm start
     ```
   - Isso executará o script Go diretamente através do NPM.

## Modificação

Para modificar o Projeto Ceres, a maneira mais simples é através do sistema de cases que você encontrará na pasta `src/Commands/Cases`. Você encontrará exemplos de comandos que podem ser usados como base para criar novas funcionalidades.

Todos os parâmetros do Projeto Íris estão acessíveis via `env["nomeDaVariavel"]`, permitindo que você utilize as funcionalidades do Projeto Íris em seu código GO, mas esteja atento a sintaxe dele, que é bem rigida quanto a certas tarefas.

## Detalhes Adicionais

**Informações da Versão**:
- **Codinome**: RABBIT
- **Versão**: v1.0.0
- **Tipo**: BETA
- **Erros**: Nenhum bug grave detectado
- **Data de Lançamento**: 04/08/2024
- **Observações**: Esta versão pode apresentar problemas menores não graves devido à ausência de alguns parâmetros opcionais ainda não integrados no Projeto Íris. Atualizações futuras do Projeto Íris resolverão essas questões, garantindo a integração completa e o funcionamento adequado dos parâmetros. Não será necessário reinstalar a Ceres para aplicar essas atualizações, pois os parâmetros já estarão incorporados nas futuras versões da Íris, e nenhuma intervenção adicional será necessária no Projeto Ceres, a menos que haja novas atualizações da mesma.

## Desenvolvimento Futuro

Mais novidades poderão chegar em breve! Fique atento às atualizações e acompanhe as redes sociais para mais informações!

Obrigado pelo seu interesse e apoio! Vamos continuar evoluindo juntos a um open-source melhor! ❤️
