fn main() {
    // Specify the library search path (equivalent to -Lbin)
    println!("cargo:rustc-link-search=native=../../../bin");
    // Link against libtulip.so (equivalent to -ltulip)
    println!("cargo:rustc-link-lib=tulip");
    // Embed the library path for runtime (equivalent to -Wl,-rpath,bin)
    println!("cargo:rustc-link-arg=-Wl,-rpath,../../../bin");
}