## 🚀 Guia de Contribuição para o Projeto Ceres

Estamos animados por você querer contribuir para o Projeto Ceres! Este guia passo a passo foi elaborado para ajudá-lo a fazer uma pull request de forma eficiente. Agradecemos sua participação!

### 📚 Passos para Contribuir

1. **Faça um Fork do Repositório**
   - Clique no botão **"Fork"** no canto superior direito da página inicial do Projeto Ceres no GitHub para criar uma cópia do repositório na sua conta.
   - Você também pode fazer um fork [clicando aqui](https://github.com/KillovSky/Ceres/fork).

2. **Clone o Seu Repositório**
   - Abra o terminal.
   - Clone o repositório em sua máquina local usando o seguinte comando:
     ```bash
     git clone https://github.com/SEU_USUÁRIO/Ceres.git
     ```
   - Substitua `SEU_USUÁRIO` pelo seu nome de usuário do GitHub.

3. **Crie uma Nova Branch**
   - Navegue até o diretório do seu repositório clonado:
     ```bash
     cd Ceres
     ```
   - Crie uma nova branch para sua contribuição:
     ```bash
     git checkout -b nome-da-sua-branch
     ```
   - Escolha um nome descritivo, como `nova-feature` ou `correcao-bug`.

4. **Faça Suas Alterações**
   - Realize as alterações necessárias no código em Go.
   - Instale as dependências do projeto, se houver, utilizando o comando Go:
     ```bash
     go mod tidy
     ```
   - Certifique-se de que o código esteja limpo e siga as diretrizes do projeto.

5. **Commit das Alterações**
   - Adicione suas alterações ao commit com:
     ```bash
     git add .
     git commit -m "Descrição CURTA das suas alterações"
     ```

6. **Push das Alterações**
   - Envie suas alterações para o seu repositório no GitHub com:
     ```bash
     git push origin nome-da-sua-branch
     ```

7. **Crie uma Pull Request**
   - Volte para a página do Projeto Ceres no GitHub e clique em **"Pull Requests"**.
   - Clique no botão **"Compare & pull request"**.

8. **Descreva sua Pull Request**
   - Forneça uma descrição DETALHADA do que você fez e por que as alterações são necessárias.

9. **Revisão e Merge**
   - Após a análise por parte de @KillovSky, suas alterações serão mescladas no repositório principal, caso sejam aprovadas.

### ✨ Dicas para uma Boa Pull Request

- Mantenha a Pull Request focada em uma única tarefa ou correção.
- Forneça uma descrição clara e detalhada do que foi feito.
- Mantenha seu código Go limpo e siga as diretrizes do projeto, incluindo comentários claros e concisos.
- Certifique-se de que os testes foram realizados e que a Pull Request está funcional.
- Melhorias pequenas podem ser adiadas para serem enviadas junto com correções maiores.
- Em caso de dúvidas, visite nossas [redes sociais](https://linktr.ee/killovsky) para obter ajuda com sua Pull Request.
- Alterações menores podem ter sua Pull Request negada ou atrasada se não forem urgentes ou essenciais.
- Para edições, prefira incluir um número significativo de mudanças em vez de poucas. Isso aumenta a relevância da contribuição.
- Este guia futuramente será transferido para a Wiki própria do Projeto Ceres!

## ❤️ Obrigado por contribuir! ❤️