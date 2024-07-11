# Hash

Hash is a command-line application that calculates the hash value of a file or a string. It supports different hash algorithms such as CRC32, CRC64, MD4, MD5, BLAKE2B, BLAKE2S, BLAKE3, SHA-1, SHA-224, SHA-256, SHA-384, and SHA-512.

## Usage

```bash
$ hash -a sha256 -t "Hello World!"
# sha256: 7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069

$ hash -a blake3,sha256 -f LICENSE
# blake3: c501c3954d09d06edb3526f374d1cc9a31accec56334b2670439b1b2f1cdf46f
# sha256: 5d34d28675f1bd0cf44da629cc0dafcbf07b06dcc8e207004bbb40cf19796b2b
```
