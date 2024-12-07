# INFO: this is ran *instead* of  dot profile.. We set path stuff manually in dot bashrc.
if [ -n "$BASH_VERSION" ]; then
    printf "Have a nice day hacking! (*<*)\n";
    [ -f "$HOME/.bashrc" ]; . "$HOME/.bashrc"
fi
