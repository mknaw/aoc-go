{
  description = "Go project development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            go-tools
            golangci-lint
            delve
            gomodifytags
            impl
            gotests
            go-tools
          ];

          shellHook = ''
            # Set GOPATH to local .go directory
            export GOPATH="$PWD/.go"
            export PATH="$GOPATH/bin:$PATH"
            
            # Add any custom environment variables here
            export GO111MODULE=on
            
            echo "Go development environment loaded!"
          '';
        };
      });
}