<h1>Processo seletivo Intelbras – Sistema de Mensageria para Telemetria Veicular.</h1> 

<h2>Introdução:</h2> O objetivo deste projeto é realizar a estruturação de um sistema de mensageria para dados de telemetria veicular, através do desenvolvimento de duas aplicações de envio e recebimento de informações. 

 

<h2>Arquitetura do Sistema:</h2> O sistema deverá composto por duas aplicações independentes: 
 

1.1. Aplicação de Publicação de Eventos: Esta aplicação será responsável por gerar e publicar eventos de telemetria veicular no Kafka. Ela deverá projetada para criar, em um frequência pré-estabelecida, eventos de telemetria dos veículos, processá-las e publicá-las em tópicos específicos do Kafka. A aplicação pode ser desenvolvida utilizando uma linguagem de programação de sua escolha (por exemplo, Golang, C++, Node.js, Python). 
 

1.2. Aplicação de Consumo de Informações: Essa aplicação será responsável por consumir as informações de telemetria veicular publicadas no Kafka. Ela receberá as mensagens dos tópicos relevantes e realizará o processamento necessário para análise, armazenamento ou exibição dos dados. Esta aplicação também pode ser desenvolvida utilizando uma linguagem de programação de sua escolha. 
 

<h2>Configuração do Kafka:</h2> Para este projeto, será necessário configurar um cluster do Kafka. O cluster deve ter ao menos um nó de broker para armazenar as mensagens e um ou mais tópicos configurados para receber os eventos de telemetria veicular. 
 

<h2>Definição dos Eventos de Telemetria: </h2> O sistema deverá enviar mensagens no formato JSON, contendo as seguintes informações: Tipo da informação, horário de criação do evento, valor do sensor(este dado deve ser gerado aleatóriamente). 
 
Os seguintes eventos deverão ser contemplados pela aplicação: 
 
> -Velocidade do veículo 
> 
> -RPM do motor 
> 
> -Temperatura do motor 
> 
> -Nível de combustível  
> 
> -Quilometragem percorrida 
> 
> -Localização GPS 
> 
> -Status das luzes (faróis, lanternas, etc.) 
 
<h2> Fluxo de Funcionamento:</h2>

O fluxo básico de funcionamento do sistema será o seguinte: 

<br/>

<b>Aplicação de Publicação de Eventos: </b>

Gera informações de telemetria veicular simulando um dispositivo de sensoriamento. 

Processa e estrutura os dados conforme a definição dos eventos de telemetria. 

Publica as mensagens em tópicos do Kafka. 

<br/>

 
<b> Aplicação de Consumo de Informações: </b>

Conecta-se ao cluster do Kafka e subscreve-se aos tópicos relevantes. 

Recebe as mensagens de telemetria veicular do Kafka. 

Processa as mensagens de acordo com as necessidades do projeto (análise, armazenamento, exibição, etc.). 


<h2>Resultado esperado: </h2>Sistema de envio e recebimento de dados em tempo-real, através de plataforma de streaming de dados Kafka com suporte a monitoramento de tráfego de informações através de Interface Gráfica(GUI, CLI, UI). 

 
A aplicação deverá ser publicada em uma página do Github e enviada para o e-mail paulo.cora@intelbras.com.br até o dia 09/12/2025 às 17:00h. 