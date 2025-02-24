# GoLibBasics

This guide lists all 150 packages in the Go standard library (as of February 24, 2025), sorted by learning difficulty for beginners. Start with Tier 1 and work your way up as you grow your skills. Each tier builds on the last, from simple basics to advanced tools.

---

## Tier 1: Beginner Essentials
Start here—these are the easiest and most useful for newbies.

- **`fmt`** - Print and format text (e.g., `fmt.Println`).
- **`strings`** - Work with strings (e.g., `strings.Split`).
- **`time`** - Handle dates and times (e.g., `time.Now`).
- **`os`** - Interact with the OS (e.g., open files).
- **`errors`** - Make and manage errors (e.g., `errors.New`).
- **`math`** - Basic math (e.g., `math.Max`).
- **`strconv`** - Convert strings to numbers (e.g., `strconv.Atoi`).
- **`log`** - Log messages (e.g., `log.Println`).
- **`flag`** - Parse command-line args (e.g., `flag.String`).
- **`sort`** - Sort lists (e.g., `sort.Ints`).

---

## Tier 2: Next Steps
Add these once you’re comfy with Tier 1—great for small projects.

- **`io`** - Basic input/output (e.g., copy data).
- **`bytes`** - Handle byte slices (like strings).
- **`encoding/json`** - Work with JSON (e.g., `json.Marshal`).
- **`net/http`** - Build web servers (e.g., `http.Get`).
- **`sync`** - Manage goroutines (e.g., `sync.Mutex`).
- **`testing`** - Write tests (e.g., `TestX` functions).
- **`bufio`** - Efficient reading/writing.
- **`path`** - Handle paths (e.g., URLs).
- **`path/filepath`** - Work with file paths.
- **`context`** - Manage timeouts.

---

## Tier 3: Intermediate Tools
Try these when you’re ready to level up.

- **`math/big`** - Big numbers.
- **`math/rand`** - Random numbers.
- **`unicode`** - Unicode characters.
- **`unicode/utf8`** - UTF-8 encoding.
- **`encoding/base64`** - Base64 encoding.
- **`mime`** - MIME types.
- **`net/url`** - Parse URLs.
- **`os/exec`** - Run commands.
- **`text/template`** - Text templates.
- **`html/template`** - Safe HTML templates.
- **`regexp`** - Regular expressions.
- **`container/list`** - Linked list.
- **`container/heap`** - Heap structure.
- **`container/ring`** - Circular list.

---

## Tier 4: Specialized Utilities
Use these for specific tasks as you grow.

- **`archive/tar`** - Tar files.
- **`archive/zip`** - ZIP files.
- **`compress/gzip`** - GZIP compression.
- **`encoding/csv`** - CSV files.
- **`encoding/xml`** - XML data.
- **`net`** - Networking basics.
- **`net/mail`** - Email parsing.
- **`net/smtp`** - Send emails.
- **`os/signal`** - Handle signals.
- **`runtime`** - Runtime info.
- **`sync/atomic`** - Atomic ops.
- **`text/tabwriter`** - Tabbed text.
- **`time/tzdata`** - Time zones.

---

## Tier 5: Advanced Tools
For bigger projects or deeper systems.

- **`builtin`** - Built-in stuff (auto-included).
- **`cmp`** - Compare values.
- **`crypto`** - Crypto basics.
- **`crypto/aes`** - AES encryption.
- **`crypto/cipher`** - Cipher tools.
- **`crypto/des`** - DES encryption.
- **`crypto/ecdsa`** - ECDSA signatures.
- **`crypto/ed25519`** - Ed25519 signatures.
- **`crypto/elliptic`** - Elliptic curves.
- **`crypto/hmac`** - HMAC auth.
- **`crypto/md5`** - MD5 hash.
- **`crypto/rand`** - Secure random.
- **`crypto/rc4`** - RC4 cipher.
- **`crypto/rsa`** - RSA encryption.
- **`crypto/sha1`** - SHA-1 hash.
- **`crypto/sha256`** - SHA-256 hash.
- **`crypto/sha512`** - SHA-512 hash.
- **`crypto/subtle`** - Subtle crypto.
- **`crypto/tls`** - TLS security.
- **`crypto/x509`** - X.509 certs.
- **`crypto/x509/pkix`** - PKIX utils.
- **`database/sql`** - SQL databases.
- **`database/sql/driver`** - SQL drivers.
- **`embed`** - Embed files.
- **`expvar`** - Export vars.
- **`hash`** - Hash interfaces.
- **`hash/adler32`** - Adler-32.
- **`hash/crc32`** - CRC-32.
- **`hash/crc64`** - CRC-64.
- **`hash/fnv`** - FNV hash.
- **`hash/maphash`** - General hash.
- **`image`** - Image base.
- **`image/color`** - Colors.
- **`image/draw`** - Draw images.
- **`image/gif`** - GIF support.
- **`image/jpeg`** - JPEG support.
- **`image/png`** - PNG support.
- **`index/suffixarray`** - Substring search.
- **`io/fs`** - File system abstraction.
- **`log/slog`** - Structured logs.
- **`math/bits`** - Bit ops.
- **`math/cmplx`** - Complex numbers.
- **`net/http/cgi`** - CGI support.
- **`net/http/cookiejar`** - Cookies.
- **`net/http/fcgi`** - FastCGI.
- **`net/http/httptest`** - HTTP tests.
- **`net/http/httptrace`** - HTTP tracing.
- **`net/http/pprof`** - HTTP profiling.
- **`net/netip`** - IP handling.
- **`net/rpc`** - RPC support.
- **`net/rpc/jsonrpc`** - JSON-RPC.
- **`net/textproto`** - Text protocols.
- **`os/user`** - User info.
- **`plugin`** - Load plugins.
- **`reflect`** - Reflection.
- **`runtime/cgo`** - Cgo support.
- **`runtime/debug`** - Debug tools.
- **`runtime/metrics`** - Metrics.
- **`runtime/pprof`** - Profiling.
- **`runtime/trace`** - Tracing.
- **`syscall`** - System calls.
- **`text/scanner`** - Scan text.
- **`unicode/utf16`** - UTF-16.

---

## Tier 6: Expert Tools
For pros or Go internals.

- **`compress/bzip2`** - Bzip2 compression.
- **`compress/flate`** - DEFLATE compression.
- **`compress/lzw`** - LZW compression.
- **`compress/zlib`** - Zlib compression.
- **`crypto/dsa`** - DSA signatures.
- **`debug/buildinfo`** - Build info.
- **`debug/dwarf`** - DWARF debug.
- **`debug/elf`** - ELF parsing.
- **`debug/gosym`** - Symbol tables.
- **`debug/macho`** - Mach-O parsing.
- **`debug/pe`** - PE parsing.
- **`encoding/ascii85`** - ASCII85 encoding.
- **`encoding/asn1`** - ASN.1 encoding.
- **`encoding/base32`** - Base32 encoding.
- **`encoding/binary`** - Binary encoding.
- **`encoding/gob`** - Go serialization.
- **`encoding/hex`** - Hex encoding.
- **`encoding/pem`** - PEM encoding.
- **`go/ast`** - Go AST parsing.
- **`go/build`** - Build packages.
- **`go/build/constraint`** - Build constraints.
- **`go/constant`** - Constants.
- **`go/doc`** - Extract docs.
- **`go/format`** - Format code.
- **`go/importer`** - Import info.
- **`go/parser`** - Parse Go.
- **`go/printer`** - Print AST.
- **`go/scanner`** - Scan tokens.
- **`go/token`** - Token handling.
- **`go/types`** - Type checking.
- **`html`** - HTML escaping.
- **`io/ioutil`** - I/O utils (old).
- **`mime/multipart`** - Multipart MIME.
- **`mime/quotedprintable`** - Quoted-Printable.
- **`regexp/syntax`** - Regexp syntax.
- **`testing/fstest`** - FS tests.
- **`testing/iotest`** - I/O tests.
- **`testing/quick`** - Quick tests.
- **`text/template/parse`** - Template parsing.
- **`unsafe`** - Unsafe ops.

---

## How to Use This
1. **Begin with Tier 1**: Try `fmt`, `strings`, and `time` in a "Hello, World" app.
2. **Move Up**: Add Tier 2 for web or JSON projects.
3. **Explore**: Use higher tiers as your projects grow (e.g., `crypto` for security).

Happy coding! Want examples? Ask away.
