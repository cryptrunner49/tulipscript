use libc::{c_char, c_int};
use std::env;
use std::ffi::CString;
use std::ptr;

// FFI declarations for Tulip functions
#[link(name = "tulip")]
unsafe extern "C" {
    fn Tulip_Init(argc: c_int, argv: *const *mut c_char);
    fn Tulip_Interpret(csrc: *const c_char, cname: *const c_char) -> c_int;
    fn Tulip_RunFile(cpath: *const c_char) -> c_int;
    fn Tulip_Free();
}

fn main() {
    // Convert command-line arguments to C-style argc/argv
    let args: Vec<String> = env::args().collect();
    let argc = args.len() as c_int;
    let mut argv: Vec<*mut c_char> = args
        .iter()
        .map(|arg| {
            CString::new(arg.as_str())
                .expect("Failed to convert argument to CString")
                .into_raw()
        })
        .collect();
    argv.push(ptr::null_mut()); // Null-terminate argv

    // Initialize TulipScript VM
    unsafe {
        Tulip_Init(argc, argv.as_ptr());
    }

    // Run Tulip script
    let result: c_int;
    if args.len() > 1 {
        // Run script from file
        let path = match CString::new(args[1].as_str()) {
            Ok(path) => path,
            Err(e) => {
                eprintln!("Error: Invalid file path: {}", e);
                return;
            }
        };
        result = unsafe { Tulip_RunFile(path.as_ptr()) };
    } else {
        // Run inline script
        let source = CString::new("println(\"Hello from TulipScript!\");").expect("Failed to create source CString");
        let name = CString::new("<test>").expect("Failed to create name CString");
        result = unsafe { Tulip_Interpret(source.as_ptr(), name.as_ptr()) };
    }

    // Print result
    println!("Result: {}", result);
    if result != 0 {
        eprintln!("TulipScript execution failed with code {}", result);
    }

    // Free TulipScript VM
    unsafe {
        Tulip_Free();
    }

    // Clean up argv
    for arg in argv {
        if !arg.is_null() {
            unsafe {
                let _ = CString::from_raw(arg);
            }
        }
    }
}