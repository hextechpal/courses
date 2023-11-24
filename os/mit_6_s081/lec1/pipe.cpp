#include <unistd.h>
#include <cstring>
#include <cstdio>
#include <sys/wait.h>

int main(int argc, char* argv[]){
    int p[2];
    char *args[2];

    args[0]=strdup("wc");
    args[1]=nullptr;

    pipe(p);

    if(fork() == 0){
        close(0);
        dup(p[0]);
        close(p[0]); // This is not strictly required but its a good practice to close any descriptors not used
//        close(p[1]); // We have to close all the read
        execvp("/usr/bin/wc",args);
    }else{
        close(p[0]);
        write(p[1],"helloworld\n",12);
        close(p[1]);
        wait(nullptr);
    }
}