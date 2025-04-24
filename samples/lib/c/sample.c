#include "libtulip.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv) {
    Tulip_Init(argc, argv);

    
    if (argc > 1) {
        Tulip_RunFile(argv[1]);
    } else {
        int exitCode;
        char* result;

        result = Tulip_InterpretWithResult("1 + 2;", "<test>", &exitCode);

        if (exitCode == 0) {
            printf("Last value: %s\n", result);
        } else {
            printf("Execution failed with code %d\n", exitCode);
        }

        free(result); // Free the returned string
    }
    
    Tulip_Free();
    return 0;
}