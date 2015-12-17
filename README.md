# mahmut
A tiny encryption library for golang.

#Installation
`go get -u github.com/melihmucuk/mahmut`

#Usage
import package

`import "github.com/melihmucuk/mahmut"`

####func Bagla(key string, text string) (string, error)

<pre><code>encryptedText, _ := mahmut.Bagla("your secret key!", "mahmut sıkı bağla da düşmeyelim yeğenim")
fmt.Println(encryptedText)
// output: S9u5pPR7h/XkgzznPwZvsluTDYC/JvrKYUvE++j7q2z+OwhKkKO+2BeasN4fiZymjBUCo+1K47Dm88oEHw==
</code></pre>

####func Coz(key string, cryptoText string) (string, error)

<pre><code>decryptedText, _ := mahmut.Coz("your secret key!", encryptedText)
fmt.Println(decryptedText)
// output: mahmut sıkı bağla da düşmeyelim yeğenim
</code></pre>

#Notes
`key` length must be 16, 24 or 32 chars.
