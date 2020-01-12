pluginDir="/usr/share/beluga/plugins"

# Make sure the plugin directory exists
if [ ! -d ${pluginDir} ]
then
    sudo install -dm 00755 "${pluginDir}"
fi
# Install the plugin library file
sudo install -Dm 00644 beluga-facts-plugin.so "${pluginDir}/beluga-facts-plugin.so"
