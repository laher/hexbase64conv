# hexbase64conv

convert between base64 and hexidecimal representations

Two commands:

 * `hextob64`
 * `b64tohex`


Note that base64 has 2 or 3 variants (that I know of), so:

 * hextob64 presents multiple outputs, unless you specify `-u` (URL) or `-s` (Standard) variants

```
❱ hextob64 0a2b82d7b9fb0bbfa3797067dc7db133f24d9bc94d1f6a7e254847ad6b3b424b6e1d6a34a6b30b51cb7b2f1bca
url: CiuC17n7C7-jeXBn3H2xM_JNm8lNH2p-JUhHrWs7QktuHWo0prMLUct7LxvK
std: CiuC17n7C7+jeXBn3H2xM/JNm8lNH2p+JUhHrWs7QktuHWo0prMLUct7LxvK

❱ hextob64 -u 0a2b82d7b9fb0bbfa3797067dc7db133f24d9bc94d1f6a7e254847ad6b3b424b6e1d6a34a6b30b51cb7b2f1bca
CiuC17n7C7-jeXBn3H2xM_JNm8lNH2p-JUhHrWs7QktuHWo0prMLUct7LxvK

❱ hextob64 -s 0a2b82d7b9fb0bbfa3797067dc7db133f24d9bc94d1f6a7e254847ad6b3b424b6e1d6a34a6b30b51cb7b2f1bca
CiuC17n7C7+jeXBn3H2xM/JNm8lNH2p+JUhHrWs7QktuHWo0prMLUct7LxvK
```

 * b64tohex tries both the URL variant and the standard variant, so it can handle either

```
❱ b64tohex CiuC17n7C7-jeXBn3H2xM_JNm8lNH2p-JUhHrWs7QktuHWo0prMLUct7LxvK
0a2b82d7b9fb0bbfa3797067dc7db133f24d9bc94d1f6a7e254847ad6b3b424b6e1d6a34a6b30b51cb7b2f1bca

❱ b64tohex CiuC17n7C7+jeXBn3H2xM/JNm8lNH2p+JUhHrWs7QktuHWo0prMLUct7LxvK
0a2b82d7b9fb0bbfa3797067dc7db133f24d9bc94d1f6a7e254847ad6b3b424b6e1d6a34a6b30b51cb7b2f1bca
```


I haven't tried to deal with the third variant, urlencoded strings for query parameters. Happy to add it
