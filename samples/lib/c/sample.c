#include "libtulip.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv) {
    // Initialize the Tulip scripting environment
    Tulip_Init(argc, argv);

    if (argc > 1) {
        // Run Tulip script from a file
        Tulip_RunFile(argv[1]);
    } else {
        int exitCode;
        char* result;

        // Interpret a Tulip script and capture the result
        result = Tulip_InterpretWithResult("1 + 2;", "<test>", &exitCode);

        if (exitCode == 0) {
            printf("Last value: %s\n", result);
        } else {
            printf("Execution failed with code %d\n", exitCode);
        }

        // Free the result string to prevent memory leaks
        free(result);
    }
    
    // Clean up Tulip scripting environment resources
    Tulip_Free();
    return 0;
}