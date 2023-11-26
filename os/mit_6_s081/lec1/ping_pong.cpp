#include <unistd.h>
#include <cstring>
#include <cstdio>
#include <sys/wait.h>
#include <cstdlib>
#include <sys/time.h>

/*
 * Program that uses UNIX system calls to “ping-pong” a byte between two processes over a pair of pipes,
 * one for each direction. Measure the program’s performance, in exchanges per second
 */


int main(){
    int p1[2]; // Pipe 1 -> Process 1 sends ping from write end Process 2 receives on read end
    int p2[2]; // Pipe 2 -> Process 2 sends ping from write end Process 1 receives on read end

    if (pipe(p1) == -1 || pipe(p2) == -1) {
        perror("pipe");
        exit(EXIT_FAILURE);
    }

    pid_t child = fork();

    if (child == -1){
        perror("fork");
        exit(EXIT_FAILURE);
    }

    if (child == 0){
        close(p1[1]);
        close(p2[0]);

        for (int i = 0; i <1000; ++i) {
            char byte;
            read(p1[0], &byte, 1);
            write(p2[1], &byte, 1);
        }

        close(p1[0]);
        close(p2[1]);
        return 0;
    }

    close(p1[0]);
    close(p2[1]);

    struct timeval start = {}, end = {};

    gettimeofday(&start, nullptr);
    for (int i = 0; i <1000; ++i) {
        char byte = 'A' + (i % 26);
        write(p1[1], &byte, 1);
        read(p2[0], &byte, 1);
    }
    gettimeofday(&end, nullptr);

    double elapsed = double (end.tv_sec - start.tv_sec) + ((end.tv_usec - start.tv_usec) / 1000000.0);

    printf("time elapsed %f", elapsed);
    printf("Exchanges per sec : %f", 1000/elapsed);
    close(p1[1]);
    close(p2[0]);
    wait(nullptr);
    return 0;
}