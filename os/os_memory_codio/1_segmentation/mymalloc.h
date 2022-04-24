#include<stdio.h>
#include<stddef.h>

/**
 * Array representing memory
 */
char memory[20000];

/**
 * Metadata block structure It has 3 properties
 *
 * size : size of the representative data block
 * free : indicated if the chunk is free
 * free : pointer to next chunk
 *       --------------
 *      | Meta Data   |
 *      --------------
 *      |   Data     |
 *      |            |
 *      -------------
 */
struct block{
    size_t size;
    int free;
    struct block *next;
};

/**
 * freelist pointing to the start of the memory
 */
struct block *freeList=(void*)memory;

void initialize();
void split(struct block *fitting_slot,size_t size);
void *MyMalloc(size_t noOfBytes);
void merge();
void MyFree(void* ptr);