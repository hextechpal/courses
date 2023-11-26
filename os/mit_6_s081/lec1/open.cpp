//
// Created by Prashant Pal on 25/11/23.
//

#include <cstdlib>
#include <fcntl.h>
#include <unistd.h>

int main(){
    int fd = open("output.txt", O_WRONLY | O_CREAT);
    if(fd == -1){
        exit(EXIT_FAILURE);
    }
    write(fd, "lolo\n", 5);
    return 0;
}