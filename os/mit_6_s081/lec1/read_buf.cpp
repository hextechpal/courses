//
// Created by Prashant Pal on 24/11/23.
//

#include <cstdio>
#include <iostream>

int getcmd(char *buf, int nbuf){
    fprintf(stdout, "$ ");
    memset(buf, 0, nbuf);
    std::cin.getline(buf, nbuf);
    printf("buf:%s\n", buf);
    if(buf[0] == '\0'){
        return -1;
    }
    return 0;
}

int main() {
    static char buf[10];
    // Read and run input commands.
    while (getcmd(buf, sizeof(buf)) >= 0) {
        printf("main:%s\n", buf);
    }
    exit(0);
}