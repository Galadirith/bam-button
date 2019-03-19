# BAM Button :radio_button:

Once upon a time there was a brave young person whos only wish was for a big
red button that connected to a USB port on a computer and when you hit that
button it would produce legs :jeans:

Unfortunately single use USB buttons are *FRAKING EXPENSIVE* :joy: but their
fairy hacker had a plan. Instead of a USB button, why not use something that is
already in their possession, a phone :phone:

## Installation :rocket:

### Pre-built binary :gift:

Go to the [releases](https://github.com/Galadirith/bam-button/releases) page
and download the latest binary archive for your platform :blush: Once you have
downloaded the archive extract the pre-built binary `bam`.

### Build your own binary :wrench:

1. **Clone `bam-button`**

   ```bash
   git clone
   cd bam-button
   ```

2. **Get Go dependencies**

   ```bash
   GOPATH=$(pwd)/go go get github.com/go-vgo/robotgo
   GOPATH=$(pwd)/go go get github.com/gorilla/websocket
   ```

   You can omit the `GOPATH` variable if you would like to use the global Go
   workspace for building `bam-button`

3. **Patch `robotgo` for Max (Optional)**

   If `robotgo` fails to install into your workspace then the following patch
   maybe fix the problem.

   ```bash
   cd go/src/github.com/go-vgo/robotgo/
   git checkout ef847f3
   git apply --unidiff-zero ../../../../../robotgo.patch
   ```

4. **Build `bam`**

   ```bash
   GOPATH=$(pwd)/go CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" ./bam.go
   ```

## Usage :tada:

Navigate to https://bam-button.glitch.me on your phone and wait until the
website has loaded and you can see the *BAM Button*. Now, using the binary
`bam` that you got from either installation method, simply run the binary `bam`.

```bash
./bam
```

Once it is running whenever you press the bam button on
https://bam-button.glitch.me you will get a pair of legs :jeans: written to
whatever document you have open :blush:

## License :sunglasses:

`bam-button` is released under the [MIT license](LICENSE.md).
