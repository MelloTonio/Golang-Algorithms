#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Usuario {
    char nome[15];
    int idade;
    char cidade;
};


struct  Node
{
    // Informação que o Node carrega.
    struct Usuario Pessoa;
    // Ponteiro para o proximo Node.
    struct Node* next;
    // Ponteiro para o Node anterior.
    struct Node* prev;
};

struct Node* head;

struct Node* GetNewNode(struct Usuario Person){
    // Aloca um espaçõ de memoria dinamico para um novo Node.
    struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
    // Define todos os campos desse novo Node.
    newNode->Pessoa = Person;
    newNode->next = NULL;
    newNode->prev = NULL;
    return newNode;
}

void InsertAtHead(struct Usuario Person){
    struct Node* newNode = GetNewNode(Person);

    if (head == NULL){
        head = newNode;
        return;
    }
}

void Print(){
    struct Node* temp = head;
    // Enquanto houver Nodes mostre na tela o conteúdo deles.
    while (temp != NULL){
        printf("%s ",temp->Pessoa.nome);
        // Avança para o proximo Node.
        temp = temp->next;
    }
    printf("\n");
}

void ReversePrint(){
    struct Node* temp = head;
    if(temp == NULL) return; // empty list, exit
    while (temp->next != NULL){
        // Avança até o ultimo Node.
        temp = temp->next;
    }

    while(temp != NULL){
        printf("%s ", temp->Pessoa.nome);
        // Como estamos voltando do final ao começo, atribuimos prev ao temp.
        temp = temp->prev;
    }

    printf("\n");
}

void InsertAtTail(struct Usuario Person) {
	struct Node* temp = head;
	struct Node* newNode = GetNewNode(Person);

	if(head == NULL) {
		head = newNode;
		return;
	}
    // Temp = ultimo elemento da lista.
	while(temp->next != NULL) temp = temp->next; 

    // Swap
    // Coloca como proximo (do atual ultimo node) o novo node
	temp->next = newNode;
    // Coloca como prévio do novo node, o antigo ulitmo node (agora penultimo).
	newNode->prev = temp;
}

int main() {
    head = NULL;
    struct Usuario Pessoa1, Pessoa2;

    strcpy(Pessoa1.nome, "Sara" );
    strcpy(Pessoa2.nome, "Joana" );

    InsertAtHead(Pessoa1); Print(); ReversePrint();
    InsertAtTail(Pessoa2); Print(); ReversePrint();
}

