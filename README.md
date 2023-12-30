# nostr-whois

Show profile metadata of npub

## Usage

```
$ nostr-whois -h
Usage of ./nostr-whois:
  -json
    	output JSON
  -relay value
    	relays to connect

$ nostr-whois npub1937vv2nf06360qn9y8el6d8sevnndy7tuh5nzre4gj05xc32tnwqauhaj6
Pubkey: 2c7cc62a697ea3a7826521f3fd34f0cb273693cbe5e9310f35449f43622a5cdc
Name: mattn
DisplayName: mattn
WebSite: https://compile-error.net
Picture: https://nostrcheck.me/media/public/eee2f2752096a9e1edc98999dcc03f017dd3ed8fd2267f46106a183fde35a37f.webp
NIP-05: _@compile-error.net
LUD-16: grimyend76@walletofsatoshi.com
About: Long-time #Golang user&contributor,  #GoogleDevExpert  Go, #Vim, #Windows hacker, #GitHubStars, #runner.
```

## Installation

```
go install github.com/mattn/nostr-whois@latest
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
