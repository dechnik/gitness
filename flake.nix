{
  description = "Gitness is an Open Source developer platform with Source Control management, Continuous Integration and Continuous Delivery";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      supportedSystems = [ "x86_64-linux" "aarch64-linux" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      pkgsFor = nixpkgs.legacyPackages;
    in
    rec {
      nixosModules.default = import nix/module.nix self;

      packages = forAllSystems (system: {
        default = pkgsFor.${system}.callPackage nix/default.nix { inherit self nixpkgs; };
      });
      hydraJobs = packages;
    };
}