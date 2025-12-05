# Installation Guide - go-projo

## Quick Install

### Method 1: Install via `go install` (Recommended)

Once the project is published to GitHub:

```bash
go install github.com/yogabagas/go-projo@latest
```

Then you can use it anywhere:
```bash
go-projo gen -name myapp -module github.com/user/myapp -type api
```

### Method 2: Build and Install Locally

```bash
# Navigate to the go-projo directory
cd go-projo

# Install to your GOPATH/bin
make install
# or
go install
```

This will install the `go-projo` binary to your `$GOPATH/bin` directory.

**Make sure `$GOPATH/bin` is in your PATH:**

```bash
# Add to your ~/.bashrc, ~/.zshrc, or ~/.profile
export PATH="$PATH:$(go env GOPATH)/bin"

# Reload your shell configuration
source ~/.bashrc  # or ~/.zshrc
```

### Method 3: Manual Build and Copy

```bash
# Build the binary
cd go-projo
make build

# Copy to a directory in your PATH
sudo cp go-projo /usr/local/bin/

# For macOS with Homebrew
cp go-projo /opt/homebrew/bin/

# Or add to your local bin (create if needed)
mkdir -p ~/bin
cp go-projo ~/bin/
export PATH="$PATH:$HOME/bin"  # Add to your shell config
```

## Verify Installation

Check if go-projo is installed correctly:

```bash
go-projo version
```

You should see:
```
go-projo version 1.0.0
```

## Update

### If installed via `go install`:
```bash
go install github.com/yogabagas/go-projo@latest
```

### If installed locally:
```bash
cd go-projo
git pull  # if you have it in a git repo
make install
```

## Uninstall

### If installed via `go install` or `make install`:
```bash
rm $(go env GOPATH)/bin/go-projo
# or
cd go-projo
make uninstall
```

### If manually installed:
```bash
# Remove from wherever you copied it
sudo rm /usr/local/bin/go-projo
# or
rm /opt/homebrew/bin/go-projo
# or
rm ~/bin/go-projo
```

## First Steps After Installation

1. **Verify installation:**
   ```bash
   go-projo version
   ```

2. **See available commands:**
   ```bash
   go-projo help
   ```

3. **Generate your first project:**
   ```bash
   go-projo gen -name demo -module github.com/myuser/demo -type api
   ```

4. **Navigate and run:**
   ```bash
   cd demo
   go mod tidy
   make run
   ```

## Troubleshooting

### Command not found

**Issue:** `go-projo: command not found`

**Solutions:**

1. Check if GOPATH/bin is in your PATH:
   ```bash
   echo $PATH | grep $(go env GOPATH)/bin
   ```

2. If not, add it to your shell configuration:
   ```bash
   echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc
   source ~/.bashrc
   ```

3. Verify GOPATH:
   ```bash
   go env GOPATH
   ```

4. Check if binary exists:
   ```bash
   ls -la $(go env GOPATH)/bin/go-projo
   ```

### Permission denied

**Issue:** `permission denied` when running go-projo

**Solution:**
```bash
chmod +x $(go env GOPATH)/bin/go-projo
```

### Build fails

**Issue:** Build errors during installation

**Solutions:**

1. Ensure you have Go installed:
   ```bash
   go version
   ```

2. Update Go to a compatible version (1.20+):
   - Visit https://go.dev/dl/

3. Clean and rebuild:
   ```bash
   cd go-projo
   make clean
   make build
   ```

## Platform-Specific Notes

### macOS

- Use `/opt/homebrew/bin/` for Apple Silicon Macs
- Use `/usr/local/bin/` for Intel Macs
- May need to allow the binary in System Preferences → Security & Privacy

### Linux

- Use `/usr/local/bin/` for system-wide installation
- Use `~/bin/` or `~/.local/bin/` for user installation
- Ensure the installation directory is in your PATH

### Windows

Using PowerShell:

```powershell
# Build
cd go-projo
go build -o go-projo.exe .

# Copy to a directory in PATH
Copy-Item go-projo.exe -Destination "$env:GOPATH\bin\"

# Or install directly
go install
```

Add GOPATH\bin to your PATH if not already:
1. Open System Properties → Advanced → Environment Variables
2. Edit Path variable
3. Add `%GOPATH%\bin`

## Development Installation

If you want to contribute or modify go-projo:

```bash
# Clone the repository
git clone https://github.com/yogabags/go-projo.git
cd go-projo

# Build
make build

# Test
./go-projo version

# Install for development
make install
```

## Shell Completion (Optional)

You can add shell completion by creating a wrapper script:

**Bash completion example** (~/.bash_completion or ~/.bashrc):
```bash
_go_projo_completion() {
    local cur prev
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    case ${COMP_CWORD} in
        1)
            COMPREPLY=($(compgen -W "gen generate version help" -- ${cur}))
            ;;
        *)
            if [[ ${prev} == "gen" ]] || [[ ${prev} == "generate" ]]; then
                COMPREPLY=($(compgen -W "-name -module -type -desc -author -go-version -output -help" -- ${cur}))
            fi
            ;;
    esac
}

complete -F _go_projo_completion go-projo
```

## Next Steps

After installation, see:
- [README.md](README.md) for usage examples
- Run `go-projo help` for command overview
- Run `go-projo gen -help` for generation options

Happy coding! 🚀
