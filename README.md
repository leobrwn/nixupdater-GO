# nixupdater-GO
A go rewrite of my nix update tool, includes a read me for instructions on how to invoke the tool via the CMD.


To add this tool to your system so you can use it by typing "nixupdate" do this:

1) mkdir -p ~/.local/bin
2) Move into the dir with the nix executable then run "install -m 755 nixupdate ~/.local/bin/nixupdate"
3) <-- Export the path -->
4) Export PATH="$HOME/.local/bin:$PATH"
5) Then reload the shell config "source ~/.bashrc"
6) Now open a new terminal and run "nixupdate"
