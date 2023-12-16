<h1>ToDo API em Go </h1>

<p style='font-size:10pt'>Este é um simples projeto de API RESTful em Go para gerenciar uma lista de tarefas (to-do list). A aplicação fornece endpoints para criação e listagem de tarefas, permitindo aos usuários acompanhar suas atividades.</p>

<h2>Como Usar</h2>

<h3>Configuração do Banco de Dados</h3>

<p style='font-size:10pt'>A aplicação utiliza um banco de dados leve, como SQLite. Certifique-se de criar um banco de dados antes de executar o aplicativo.</p>

<h3> Executando o Aplicativo </h3>

<p style='font-size:10pt'>Execute o aplicativo utilizando o comando go run main.go no diretório do projeto.
Acesse a API através de um cliente HTTP ou ferramentas como cURL ou Postman.</p>

<h3> Endpoints da API </h3>

<div style="padding:calc(20px/1.3);border-radius:16px; border:2px solid #0CA;background-color:#09a"><ol style='list-style:none'><code>GET /tasks: Lista todas as tarefas.</code></ol><ol style='list-style:none'><code>GET /tasks/{id}: Retorna os detalhes de uma tarefa específica.</code>
</ol><ol style='list-style:none'><code>POST /tasks: Cria uma nova tarefa.</code></ol></div>
<!-- PUT /tasks/{id}: Atualiza os detalhes de uma tarefa existente.
DELETE /tasks/{id}: Exclui uma tarefa. -->

<h3> Estrutura do Projeto </h3>

<li><b>main.go:</b> Ponto de entrada da aplicação.</li>
<li><b>handlers.go:</b> Manipuladores para cada endpoint da API.</li>
<li><b>models.go:</b> Definição de estrutura de dados (Task).</li>
<li><b>database.go:</b> Configuração e interação com o banco de dados.</li>
<li><b>routes.go:</b> Configuração das rotas da API.</li>

<h3> Tecnologias Utilizadas </h3>

<li>Go (Golang)</li>
<li>Gorilla Mux para roteamento</li>
<li>SQLite para persistência de dados</li>
<li>JSON para comunicação entre cliente e - servidor</li>
