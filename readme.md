​💳 EMV Transaction Engine - Tupi Fintech
​Este repositório contém o desafio técnico desenvolvido para a posição de Desenvolvedor Backend Go na Tupi Fintech.o projeto consiste em um motor de processamento de transações EMV (Chip) escrito em Go,simulando o fluxo entre um terminal de pagamento e o emissor.

​🎯 Objetivo

​Implementar um serviço robusto capaz de decodificar dados de cartões, validar regras de negócio financeiras (Luhn, Validade, CVM) e registrar logs auditáveis de cada operação.


​🏗️ Arquitetura e Design

​A arquitetura foi pensada para ser modular e extensível, seguindo boas práticas de Go:
​SoC (Separation of Concerns): Divisão clara entre lógica de validação, regras de negócio EMV e persistência.
​Idiosincrasia Go: Uso de tratamento de erros explícito e structs para representação de dados.

​Logs Auditáveis: Registro estruturado em JSON, essencial para conciliação bancária e debug em ambiente de produção.

🛠️ Tecnologias e Padrões
​Go 1.21+ (Foco em performance e concorrência).

​JSON Logging: Implementação de logs estruturados para facilitar integração com ELK Stack/Splunk.

​Unit Testing: Cobertura de testes nos componentes críticos de validação.

​EMV Standard: Simulação baseada nas tags 5A (PAN), 5F24 (Expiry) e 9F34 (CVM).

#como rodar o projeto
go run main.go

🧪 Validações Implementadas

​✅ Algoritmo de Luhn: Verificação de integridade do PAN (Número do Cartão).

​✅ Check de Expiração: Bloqueio de transações com cartões vencidos (comparação com data atual).

​✅ CVM Check: Validação de presença de método de verificação do portador.

​✅ Mock Gateway: Simulação de latência e resposta randômica de autorização online.
