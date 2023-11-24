#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>


int main(int argc, char *argv[]){
    printf("starting the program.. pid:%d\n", (int) getpid());
    int child = fork();

    if(child < 0){
        fprintf(stderr, "child process was not created");
        exit(1);
    }else if(child == 0){
        printf("hello from child process\n");
        char *args[4];
        args[0] = strdup("wc");
        args[1] = strdup("-w");
        args[2] = strdup("fork.cpp");
        args[3] = NULL;
        execvp(args[0], args);
        printf("this doesnt print");
    }else{
        wait(NULL);
        printf("hello from parent process, child pid:%d\n", child);
    }
    return 0;    
}