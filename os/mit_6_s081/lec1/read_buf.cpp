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
        printf("returning -1");
        return -1;
    }

    std::cout << "Remaining characters in input buffer: ";
    int ch;
    while ((ch = std::cin.peek()) != EOF) {
        std::cout.put(ch);
    }
    std::cout << std::endl;
    return 0;
}

int main() {
//    static char buf[10];
//
//    // Read and run input commands.
//    while (getcmd(buf, sizeof(buf)) >= 0) {
//        printf("main:%s\n", buf);
//    }
//
//
//    exit(0);

    char buf[10];
    std::cin.getline(buf, 10);

    std::cout << "You entered: " << buf << std::endl;

    // Attempt to read an integer
    memset(buf, 0, 10);
    std::cout << "Remaing: " << std::endl;
    std::cin.getline(buf, 10);

    std::cout << "You entered: " << buf << std::endl;
    // Check if the input operation succeeded
    if (std::cin.fail()) {
        std::cin.clear();  // Clear the fail state
        std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n');  // Ignore remaining characters
        std::cout << "Invalid input for integer. Please enter a valid integer." << std::endl;
    } else {
        std::cout << "You entered the integer: " << buf << std::endl;
    }

    return 0;
}