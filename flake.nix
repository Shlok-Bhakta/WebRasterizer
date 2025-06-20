{
  description = "Golang Rasterizer";

  nixConfig = {
    extra-substituters = [
      "https://cuda-maintainers.cachix.org"
    ];
    extra-trusted-public-keys = [
      "cuda-maintainers.cachix.org-1:0dq3bujKpuEPMCX6U4WylrUDZ9JyUG0VpVZa7CNfq5E="
    ];
  };


  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils}:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config = { allowUnfree = true; };
        };
        buildScript = pkgs.writeShellScriptBin "build" ''
          GOOS=js GOARCH=wasm go build -o main.wasm
        '';
        deploy = pkgs.writeShellScriptBin "deploy" ''
          echo "copying public to docs"
          cp -r ./public/* ./docs/
        '';
      in
      {
        devShells.default = pkgs.mkShell {
          name = "Golang Rasterizer DevShell";
          nativeBuildInputs = with pkgs; [
            # Golang
            go
            # Live reload in go
            air
            buildScript
            deploy
          ];
          # GOOS = "js";
          # GOARCH = "wasm";
          shellHook = ''
            echo "You show those triangles who's boss!"
            zsh
            alias build='GOOS=js GOARCH=wasm go build -o main.wasm'
          '';
        };
      });
}
