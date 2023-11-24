#include <unistd.h>
#include <cstdlib>
#include <cstdio>
#include <fcntl.h>
#include <cstring>

const int MIN = 1;
const int MAX = 2;

void exp1(){
    int child = fork();
    if(child == 0){
        write(1, "hello", 6);
        exit(0);
    }else{
        wait(nullptr);
        write(1, "world", 5);
    }
    write(1, "\n", 1);
}

void exp2(){
    char *argv[2];
    argv[0] = strdup("cat");
    argv[1] = nullptr;

    int child = fork();

    if(child == 0){
        close(0);
        open("input.txt", O_RDONLY);
        execvp("cat", argv);
    }else{
        write(1, "world", 5);    
    }
}

/*
 * This experiment illustrates that the child and parent process share not only the file descriptor but also
 * the file offsets. In the following code snippet the data is not override in the file but both the processes
 */
void exp3(){
    const char *filename = "example.txt";
    int fd;

    // Open the file with O_CREAT and O_WRONLY flags
    // The third parameter specifies the file permissions (in octal)
    fd = open(filename, O_CREAT | O_WRONLY);
    if (fd ==  -1){
        printf("error opening file");
        exit(EXIT_FAILURE);
    }

    int child = fork();
    if (child == 0){
        write(fd, "abc\n", 4);
        sleep(1);
        write(fd, "def\n", 4);
        sleep(1);
        write(fd, "ghi\n", 4);
    }else{
        write(fd, "jkl\n", 4);
        sleep(1);
        write(fd, "mno\n", 4);
        sleep(1);
        write(fd, "pqr\n", 4);
    }
}

int main(int argc, char *argv[]){
    if (argc != 2){
        printf("please pass the experiment number to run\n");
        return 1;
    }

    int num = atoi(argv[1]);
    switch (num){
    case 1:
        exp1();
        break;
    case 2:
        exp2();
        break;
    case 3:
        exp3();
        break;
    default:
        printf("please pass valid experiment number between %d and %d\n", MIN, MAX);
    }
    
}

