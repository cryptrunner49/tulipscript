#include "libtulip.h"
#include <iostream>

int main(int argc, char** argv) {
    // Initialize the Tulip scripting environment
    Tulip_Init(argc, argv);

    if (argc > 1) {
        // Run Tulip script from a file
       Tulip_RunFile(argv[1]);
    } else {
        int exitCode;
        const char* source = "1 + 2;";
        const char* name = "<test>";

        // Interpret a Tulip script and capture the result
        char* result = Tulip_InterpretWithResult(const_cast<char*>(source), const_cast<char*>(name), &exitCode);
        if (exitCode == 0) {
            std::cout << "Last value: " << result << std::endl;
        } else {
            std::cout << "Execution failed with code " << exitCode << std::endl;
        }

        // Free the result string to prevent memory leaks
        free(result);
    }

    // Clean up Tulip scripting environment resources
    Tulip_Free();
    return 0;
}