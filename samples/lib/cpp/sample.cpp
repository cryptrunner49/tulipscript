#include "libtulip.h"
#include <iostream>
#include <string>

int main(int argc, char** argv) {
    Tulip_Init(argc, argv);

    int result;
    if (argc > 1) {
        const char* path = argv[1];
        result = Tulip_RunFile((char*)path);
    } else {
        const char* source = "println(\"Hello from TulipScript!\");";
        const char* name = "<test>";
        result = Tulip_Interpret((char*)source, (char*)name);
    }

    std::cout << "Result: " << result << std::endl;
    Tulip_Free();
    return 0;
}

/*
make build
gcc -o samples/lib/test samples/lib/test.c -Ibin -Lbin -lseedvm
LD_LIBRARY_PATH=bin ./samples/lib/test
*/