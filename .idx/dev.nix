{ pkgs, ...}: {
    channel = "stable-23.11";
    packages = [
        pkgs.gnumake
        pkgs.go_1_22
        pkgs.docker       
    ];
    services.docker.enable = true;
    idx = {
            previews = {
                enable = false;
            };
            
            extensions = [
                "golang.go"
		        "octref.vetur"
                "eamodio.gitlens"
            ];
    };
}
