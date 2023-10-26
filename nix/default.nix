{ self, pkgs, nixpkgs }:
let
  gitnessVersion =
    if (self ? shortRev)
    then self.shortRev
    else "dev";
  src = ../.;
  gitness-ui = pkgs.stdenv.mkDerivation {
    pname = "gitness-ui";
    src = "${src}/web";
    version = gitnessVersion;
    nativeBuildInputs = with pkgs; [ nodejs yarn fixup_yarn_lock ];
    offlineCache = pkgs.fetchYarnDeps {
      yarnLock = "${src}/nix/yarn.lock";
      sha256 = "sha256-CuyxhtEiUqKAEbutRKtACostA/QXd8LcTDNC+SkusJQ=";
    };
    configurePhase = ''
      runHook preConfigure

      export HOME=$PWD/tmp
      mkdir -p $HOME

      fixup_yarn_lock yarn.lock
      yarn config --offline set yarn-offline-mirror $offlineCache
      yarn install --offline --frozen-lockfile --ignore-platform --ignore-scripts --no-progress --non-interactive
      patchShebangs node_modules

      runHook postConfigure
      '';
    buildPhase = ''
      runHook preBuild
      yarn build
      runHook postBuild
    '';

    installPhase = ''
      runHook preInstall
      mv dist $out
      runHook postInstall
    '';
  };
  protoc-gen-go = pkgs.buildGoModule rec {
    pname = "protoc-gen-go";
    version = "1.28.1";

    src = pkgs.fetchFromGitHub {
      owner = "protocolbuffers";
      repo = "protobuf-go";
      rev = "v${version}";
      sha256 = "sha256-7Cg7fByLR9jX3OSCqJfLw5PAHDQi/gopkjtkbobnyWM=";
    };

    vendorSha256 = "sha256-yb8l4ooZwqfvenlxDRg95rqiL+hmsn0weS/dPv/oD2Y=";

    subPackages = [ "cmd/protoc-gen-go" ];
  };
  protoc-gen-go-grpc = pkgs.buildGoModule rec {
    pname = "protoc-gen-go-grpc";
    version = "1.2.0";
    modRoot = "cmd/protoc-gen-go-grpc";

    src = pkgs.fetchFromGitHub {
      owner = "grpc";
      repo = "grpc-go";
      rev = "cmd/protoc-gen-go-grpc/v${version}";
      sha256 = "sha256-pIHMykwwyu052rqwSV5Js0+JCKCNLrFVN6XQ3xjlEOI=";
    };

    vendorSha256 = "sha256-yxOfgTA5IIczehpWMM1kreMqJYKgRT5HEGbJ3SeQ/Lg=";
  };
  dbmate = pkgs.buildGoModule rec {
    pname = "dbmate";
    version = "1.15.0";

    src = pkgs.fetchFromGitHub {
      owner = "amacneil";
      repo = "dbmate";
      rev = "v${version}";
      sha256 = "sha256-eBes5BqoR7K6ntCKjWECwWuoTwAodNtLqcTei5WocLU=";
    };

    vendorSha256 = "sha256-U9VTS0rmLHxweFiIcFyoybHMBihy5ezloDC2iLc4IMc=";

    doCheck = false;
  };
in
pkgs.buildGoModule {
  pname = "gitness";
  inherit src;
  version = gitnessVersion;
  vendorSha256 = "sha256-6TL6jJtVkXucc4rZrRUtZ2pKqCqjb44QOk/BuIU5yk0=";

  nativeBuildInputs = [
    pkgs.makeWrapper
    pkgs.wire
    pkgs.gci
    pkgs.gotools
    pkgs.protobuf3_21
    protoc-gen-go
    protoc-gen-go-grpc
  ];

  buildInputs = [
    pkgs.wire
    dbmate
  ];

  preBuild = ''
    cp -R ${gitness-ui}/ web/dist/
    wire gen ./cmd/gitness
    protoc --proto_path=./gitrpc/proto --go_out=./gitrpc/rpc --go_opt=paths=source_relative --go-grpc_out=./gitrpc/rpc --go-grpc_opt=paths=source_relative ./gitrpc/proto/*.proto
  '';
  postInstall = ''
    wrapProgram $out/bin/gitness \
      --prefix PATH : ${pkgs.git}/bin
  '';
}