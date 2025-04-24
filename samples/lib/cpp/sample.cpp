#include "libtulip.h"
#include <iostream>

int main(int argc, char** argv) {
    Tulip_Init(argc, argv);

    if (argc > 1) {
       Tulip_RunFile(argv[1]);
    } else {
        int exitCode;
        const char* source = "1 + 2;";
        const char* name = "<test>";
        char* result = Tulip_InterpretWithResult(const_cast<char*>(source), const_cast<char*>(name), &exitCode);
        if (exitCode == 0) {
            std::cout << "Last value: " << result << std::endl;
        } else {
            std::cout << "Execution failed with code " << exitCode << std::endl;
        }
        free(result); // Free the returned string
    }

    Tulip_Free();
    return 0;
}