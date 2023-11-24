#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>



int main(int argc, char *argv[]){
    printf("starting the program.. pid:%d\n", (int) getpid());
    int child = fork();

    if(child < 0){
        fprintf(stderr, "child process was not created");
        exit(1);
    }else if(child == 0){
        printf("hello from child process\n");
        sleep(5);
        printf("bye from child process\n");
    }else{
        wait(NULL);
        printf("hello from parent process, child pid:%d\n", child);
    }
    return 0;    
}