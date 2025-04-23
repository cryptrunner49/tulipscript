#include "libtulip.h"
#include <stdio.h>

int main(int argc, char** argv) {
    Tulip_Init(argc, argv);

    int result;
    if (argc > 1) {
        result = Tulip_RunFile(argv[1]);
    } else {
        result = Tulip_Interpret("println(\"Hello from TulipScript!\");", "<test>");
    }

    printf("Result: %d\n", result);
    Tulip_Free();
    return 0;
}

/*
make build
gcc -o samples/lib/test samples/lib/test.c -Ibin -Lbin -lseedvm
LD_LIBRARY_PATH=bin ./samples/lib/test
*/