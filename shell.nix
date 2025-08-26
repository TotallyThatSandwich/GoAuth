{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {

	buildInputs = [
    	pkgs.go
    	pkgs.sqlc
		pkgs.cloudflared
  	];
	shellHook = ''
    	echo "Starting Cloudflared tunnels..."

    	cloudflared access tcp --hostname pg.genericcursed.com --url localhost:5432 &
    	PG_TUNNEL_PID=$!

    	cloudflared access tcp --hostname rds.genericcursed.com --url localhost:6379 &
    	REDIS_TUNNEL_PID=$!

    	export PG_TUNNEL_PID REDIS_TUNNEL_PID

    	trap "kill $PG_TUNNEL_PID $REDIS_TUNNEL_PID" EXIT
  '';

}
