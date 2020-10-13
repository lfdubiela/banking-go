/**
 *  V001 - Script para criação das tabelas de account e transaction
 */

CREATE TABLE account (
    id INT AUTO_INCREMENT PRIMARY KEY,
    document_number VARCHAR(14) NOT NULL UNIQUE,
    credit_limit DECIMAL(8,2) NOT NULL,
    available_limit DECIMAL(8,2) NOT NULL
);

CREATE TABLE transaction (
    id INT AUTO_INCREMENT PRIMARY KEY,
    account_id INT NOT NULL,
    operation_id INT NOT NULL,
    amount DECIMAL(8,2) NOT NULL,
    event_date TIMESTAMP NOT NULL,

    FOREIGN KEY (account_id) REFERENCES account(id)
);