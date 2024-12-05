# Kafka Docker Com SASL-SSL
> SASL (Simple Authentication and Security Layer) SSL é um mecanismo de autenticação do Kafka que combina a segurança do SASL Plain com a criptografia SSL/TLS para proteger tanto as credenciais de usuário quanto os dados trafegados. Esta abordagem robusta oferece uma das opções mais seguras para autenticação no Kafka, garantindo a integridade e a confidencialidade das comunicações.

---------------------------------

# Configurar variaveis de ambiente

Primeiramente se deve copiar e renomear os arquivos .env_exemplo, kafka.env_exemplo para .env e kafka.env respectivamente. Para isso basta navegar até a pasta raiz do projeto /data/work/kafka-docker e executar os comandos abaixo:

```bash
cp .env_exemplo .env
```

```bash
cp kafka/environments/kafka.env_exemplo kafka/environments/kafka.env
```


O arquivo .env conta com as seguintes variaveis que podem ser configuradas:

| Variavel | Descrição |
| ------ | ------ |
| ```PROJETO``` | Nome do projeto (por padrão é o nome da pasta/repo gitlab) |
| ```PROJETO_DIR``` | Diretorio do projeto (por padrão /data/work/{PROJETO})|
| ```HOST_NAME``` | Nome do host |
| ```KAFKA_PORTS``` |  Portas do broker kafka (por padrão "9092:9092")|

O arquivo **kafka.env** conta com as seguintes variaveis que podem ser configuradas:

##### KRAFT:
| Variavel | Descrição |
| ------ | ------ |
| ```KAFKA_CFG_NODE_ID``` |  Identifica o ID único do nó Kafka no cluster |
| ```KAFKA_CFG_PROCESS_ROLES``` | Define os papéis que este nó Kafka desempenha no cluster |
| ```KAFKA_CFG_CONTROLLER_QUORUM_VOTERS``` | Especifica o endereço e a porta para o controlador do Kafka |

##### LISTENERS:
| Variavel | Descrição |
| ------ | ------ |
| ```KAFKA_CFG_LISTENERS``` | Define os protocolos e portas em que o Kafka escuta conexões de entrada. |
| ```KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP``` | Mapeia os nomes dos listeners para os protocolos de segurança correspondentes.  |
| ```KAFKA_CFG_ADVERTISED_LISTENERS``` | Define os endereços que o Kafka informa aos clientes para conexão. |
| ```KAFKA_CFG_CONTROLLER_LISTENER_NAMES``` |  Lista os nomes dos listeners que são usados para comunicação do controlador. |
| ```KAFKA_CFG_INTER_BROKER_LISTENER_NAME``` |  Define o nome do listener usado para comunicação entre brokers |
| ```KAFKA_CLIENT_LISTENER_NAME``` |  Define o nome do listener usado pelos clientes para se conectarem. |
| ```KAFKA_CFG_SASL_MECHANISM_CONTROLLER_PROTOCOL``` |  Especifica o mecanismo SASL usado para autenticação do protocolo do controlador. |
| ```KAFKA_CFG_SASL_MECHANISM_INTER_BROKER_PROTOCOL``` | Especifica o mecanismo SASL usado para autenticação entre os brokers. |
| ```KAFKA_CONTROLLER_USER``` | Nome de usuário utilizado para autenticação do controlador.  |
| ```KAFKA_CONTROLLER_PASSWORD``` | Senha associada ao usuário do controlador. |
| ```KAFKA_INTER_BROKER_USER``` | Nome de usuário utilizado para autenticação entre brokers. |
| ```KAFKA_INTER_BROKER_PASSWORD``` | Senha associada ao usuário para autenticação entre brokers. |
| ```KAFKA_CLIENT_USERS``` | Nome de usuário usado pelos clientes para autenticação. |
| ```KAFKA_CLIENT_PASSWORDS``` | Senha associada ao usuário dos clientes para autenticação.  |

##### SSL:
| Variavel | Descrição |
| ------ | ------ |
| ```KAFKA_TLS_TYPE``` |  Tipo de keystore/truststore usado para configuração de SSL/TLS. |
| ```KAFKA_CERTIFICATE_PASSWORD``` | Senha associada ao arquivo de certificado utilizado para SSL/TLS. |

-----------------------------------------

# Gerando keystore, truststore e ca.pem

obs: Quando pedir Nome e Sobrenome ou Common Name a resposta deve ser igual ao HOST_NAME condifurado no .env

#### 1° passo
**Geração de Keystore e Par de Chaves**: 
- Cria um arquivo de keystore (`server.keystore.jks`) e gera um par de chaves RSA com o alias `localhost`.

##### Gera o kafka.keystore.jks:
obs: Resposta para o Qual é o seu nome e o seu sobrenome? deve ser igual ao HOST_NAME do .env
`o <alias> deve ser o valor definido no HOST_NAME`
```bash
keytool -keystore /data/work/kafka-docker/kafka/secrets/kafka.keystore.jks -alias localhost -validity 365 -keyalg RSA -genkey
```

#### 2° passo
**Criação e Importação do Certificado da Autoridade Certificadora (CA)**:
- Usa OpenSSL para gerar um certificado autoassinado da CA e o importa no truststore do servidor (`kafka.truststore.jks`).

##### Gera ca-key e ca-cert:
obs: Resposta para o Common Name deve ser igual ao HOST_NAME do .env
```bash
openssl req -new -x509 -keyout /data/work/kafka-docker/kafka/secrets/ca-key -out /data/work/kafka-docker/kafka/secrets/ca-cert -days 365
```

##### Gera o kafka.truststore.jks: 
` o <-alias> será o nome utilizado na hora de gerar a chave .pem utilizada pelos clientes`
```bash
keytool -keystore /data/work/kafka-docker/kafka/secrets/kafka.truststore.jks -alias CARoot -import -file /data/work/kafka-docker/kafka/secrets/ca-cert
```

#### 3° passo
**Geração e Assinatura do Certificado do Servidor Kafka**:
- Gera um pedido de certificado (CSR) para o servidor.
- Usa OpenSSL para assinar o CSR com o certificado da CA gerado anteriormente.
- Importa o certificado assinado de volta no keystore do servidor (`server.keystore.jks`).

##### Gera o CSR - Certificate Signing Request:
`o <alias> deve ser o valor definido no HOST_NAME`
```bash
keytool -keystore /data/work/kafka-docker/kafka/secrets/kafka.keystore.jks -alias localhost -certreq -file /data/work/kafka-docker/kafka/secrets/cert-file
```

##### Assina o CSR utilizando o CA:
`<-passin pass> Senha necessária para acessar a chave privada do CA`
```bash
openssl x509 -req -CA /data/work/kafka-docker/kafka/secrets/ca-cert -CAkey /data/work/kafka-docker/kafka/secrets/ca-key -in /data/work/kafka-docker/kafka/secrets/cert-file -out /data/work/kafka-docker/kafka/secrets/cert-signed -days 365 -CAcreateserial -passin pass:abc123
```

##### Importa a CA para a Keystore:
`<-passin pass> Senha necessária para acessar a chave privada do CA`
```bash
keytool -keystore /data/work/kafka-docker/kafka/secrets/kafka.keystore.jks -alias CARoot -import -file /data/work/kafka-docker/kafka/secrets/ca-cert
```

##### Importa o Certificado assinado para a Keystore:
`<-passin pass> Senha necessária para acessar a chave privada do CA`
```bash
keytool -keystore /data/work/kafka-docker/kafka/secrets/kafka.keystore.jks -alias localhost -import -file /data/work/kafka-docker/kafka/secrets/cert-signed
```

#### 4° passo
**Geração do certificado .pem para utilizar nos clientes**:
- Exporta o certificado da autoridade certificadora (CA) raiz do keystore do Kafka para um arquivo PEM .

`O nome do certificado é o alias utilizado no ca`
```bash
keytool -exportcert -alias CARoot -keystore /data/work/kafka-docker/kafka/secrets/kafka.keystore.jks -rfc -file /data/work/kafka-docker/kafka/secrets/CARoot.pem
```

O arquivo .pem será utilizado nos clientes para realizar a criptografia dos dados, o kafka.keystore.jks e o kafka.truststore.jks ficaração na pasta /data/work/kafka-docker/kafka/secrets para serem utilizados pelo kafka.

-------------------------------


# Iniciar o container

Após tudo configurado, basta rodar o comando como root na pasta raiz do projeto:
( `esse comando deve ser executado como root`)
```bash
docker-compose up -d
```

#### Comandos adicionais

Após o build do projeto o mesmo podera ser acessado e configurado com os comandos abaixo:


Acessando o container kafka:
( `esse comando deve ser executado como root`)
```bash
docker exec -it <nome-container> bash
```

Criar topico
```bash
kafka-topics.sh --bootstrap-server localhost:9092 --command-config /opt/bitnami/kafka/config/consumer.properties --topic meu-topico-9999 --create --replication-factor 1 --partitions 1
```

Listar topicos
```bash
kafka-topics.sh --bootstrap-server localhost:9092 --command-config /opt/bitnami/kafka/config/consumer.properties --list
```
Iniciar um producer
```bash
kafka-console-producer.sh --bootstrap-server localhost:9092 --topic test --producer.config /opt/bitnami/kafka/config/producer.properties
```
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic telemetry
Iniciar um consumer
```bash
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --consumer.config /opt/bitnami/kafka/config/consumer.properties
```

links uteis: 
Onde encontrei a solução pro erro de SSL handshake failed

[kafka failed authentication due to: SSL handshake failed - Stack Overflow](https://stackoverflow.com/questions/54903381/kafka-failed-authentication-due-to-ssl-handshake-failed)

Onde consegui informações para criar o container com sasl_ssl

[containers/bitnami/kafka/README.md at main · bitnami/containers · GitHub](https://github.com/bitnami/containers/blob/main/bitnami/kafka/README.md)

Config dos topicos

[librdkafka/CONFIGURATION.md](https://github.com/confluentinc/librdkafka/blob/master/CONFIGURATION.md#topic-configuration-properties)

Conduktor

[conduktor install](https://docs.conduktor.io/desktop/conduktor-first-steps/install/)