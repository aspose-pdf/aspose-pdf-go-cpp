This folder contains platform-specific dynamic libraries required for the package.  
Each library is provided in a compressed `.bz2` format along with its extracted version.  

- `.bz2` files: Compressed shared libraries downloaded from the repository.  
- Extracted libraries: Decompressed `.dll`, `.so`, or `.dylib` files used at runtime.  

If missing, libraries can be downloaded and extracted automatically using:  
`go generate`  

Do not manually modify files in this folder.  
