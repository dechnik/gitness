self: { config, lib, pkgs, ... }:
with lib;
let
  cfg = config.services.gitness;
in {
  options.services.gitness = {
    enable = mkEnableOption "gitness";
    package = mkOption {
      type = types.package;
      default = self.packages.${pkgs.system}.default;
      description = "The package implementing gitness";
    };
    environmentFile = mkOption {
      type = types.nullOr types.path;
      default = null;
      description = mdDoc ''
        Environment file with variables described in
        https://docs.gitness.com/reference/configuration
      '';
    };
    httpPort = lib.mkOption {
      type = lib.types.int;
      default = 3000;
      description = "The port on which the system listens for incoming calls.";
    };
    openFirewall = lib.mkOption {
      type = lib.types.bool;
      default = true;
      description = "Whether to open port in the firewall for the server.";
    };
    dataDir = lib.mkOption {
      type = lib.types.path;
      description = "Path to data dir";
      default = "/var/lib/gitness";
    };
    environment = mkOption {
      type = with types; attrsOf str;
      default = { };
      example = {
        GITNESS_DATABASE_DATASOURCE = "database.sqlite3";
        GITNESS_DATABASE_DRIVER = "sqlite3";
        GITNESS_DEBUG = "false";
        GITNESS_ENCRYPTER_MIXED_CONTENT = "false";
        GITNESS_ENCRYPTER_SECRET = "";
        GITNESS_GRACEFUL_SHUTDOWN_TIME = "300s";
        GITNESS_PRINCIPAL_ADMIN_PASSWORD = "password";
        GITNESS_TOKEN_COOKIE_NAME = "token";
        GITNESS_TOKEN_EXPIRE = "720h";
        GITNESS_TRACE = "false";
        GITNESS_URL_API = "http://localhost:3000/api";
        GITNESS_URL_BASE = "http://localhost:3000";
        GITNESS_URL_GIT = "http://localhost:3000/git";
        GITNESS_URL_UI = "http://localhost:3000";
        GITNESS_USER_SIGNUP_ENABLED = "true";
      };
      description = lib.mdDoc ''
        Environment variables passed to the service.
        https://docs.gitness.com/reference/configuration
      '';
    };
  };
  config = mkIf cfg.enable {
    networking.firewall =
      lib.mkIf cfg.openFirewall { allowedTCPPorts = [ cfg.httpPort ]; };
    systemd.services.gitness = {
      after = [ "network-online.target" ];
      wantedBy = [ "multi-user.target" ];
      serviceConfig = {
        ExecStart = "${cfg.package}/bin/gitness server";
        WorkingDirectory = "${cfg.dataDir}";
        ReadWriteDirectories = "${cfg.dataDir}";
        Restart = "on-failure";
        User = "gitness";
        Group = "gitness";
        EnvironmentFile = lib.mkIf (cfg.environmentFile != null) cfg.environmentFile;
      };
      environment = cfg.environment // {
        GITNESS_HTTP_PORT = builtins.toString cfg.httpPort;
      };
    };
    users.users.gitness = {
      group = "gitness";
      home = "${cfg.dataDir}";
      createHome = true;
      isSystemUser = true;
    };
    users.groups.gitness = { };
  };
}